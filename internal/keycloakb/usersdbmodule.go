package keycloakb

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cloudtrust/common-service/v2/database/sqltypes"
	"github.com/cloudtrust/common-service/v2/log"
	"github.com/cloudtrust/common-service/v2/security"
	"github.com/cloudtrust/keycloak-bridge/internal/dto"
)

const (
	updateUserDetailsStmt = `INSERT INTO user_details (realm_id, user_id, details, key_id)
	  VALUES (?, ?, ?, ?) 
	  ON DUPLICATE KEY UPDATE details=?, key_id=?;`
	selectUserDetailsStmt = `
	  SELECT details, key_id
	  FROM user_details
	  WHERE realm_id=?
		AND user_id=?;`
	deleteUserDetailsStmt = `DELETE FROM user_details WHERE realm_id=? AND user_id=?;`
	createCheckStmt       = `INSERT INTO checks (realm_id, user_id, operator, datetime, status, type, nature, proof_type, proof_data, comment, key_id, txn_id)
	  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
	  ON DUPLICATE KEY UPDATE realm_id=?, user_id=?, operator=?, datetime=?, status=?, type=?, nature=?, proof_type=?, proof_data=?, comment=?, key_id=?, txn_id=?;`
	createPendingCheckStmt = `INSERT INTO checks (realm_id, user_id, operator, datetime, status, type, nature, proof_type, proof_data, comment, key_id, txn_id)
	  VALUES (?, ?, ?, ?, 'PENDING', ?, ?, ?, ?, ?, ?, ?)`
	selectCheckStmt = `
	  SELECT check_id, realm_id, user_id, operator, unix_timestamp(datetime), status, type, nature, proof_type, proof_data, comment, key_id, txn_id
	  FROM checks
	  WHERE realm_id=?
		AND user_id=?
	  ORDER BY datetime DESC;`
	selectPendingCheckStmt = `
	  SELECT check_id, realm_id, user_id, operator, unix_timestamp(datetime), status, type, nature, proof_type, proof_data, comment, key_id, txn_id
	  FROM checks
	  WHERE realm_id=?
		AND user_id=?
		AND status='PENDING'
	  ORDER BY datetime DESC;`
)

// UsersDetailsDBModule interface
type UsersDetailsDBModule interface {
	StoreOrUpdateUserDetails(ctx context.Context, realm string, user dto.DBUser) error
	GetUserDetails(ctx context.Context, realm string, userID string) (dto.DBUser, error)
	DeleteUserDetails(ctx context.Context, realm string, userID string) error
	CreateCheck(ctx context.Context, realm string, userID string, check dto.DBCheck) error
	CreatePendingCheck(ctx context.Context, realm string, userID string, check dto.DBCheck) error
	GetChecks(ctx context.Context, realm string, userID string) ([]dto.DBCheck, error)
	GetPendingChecks(ctx context.Context, realm string, userID string) ([]dto.DBCheck, error)
}

type usersDBModule struct {
	db     sqltypes.CloudtrustDB
	cipher security.EncrypterDecrypter
	logger log.Logger
}

func nullStringToPtr(value sql.NullString) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}

func nullStringToDatePtr(value sql.NullString) *time.Time {
	if value.Valid {
		var dateInt, _ = strconv.ParseInt(strings.Split(value.String, ".")[0], 10, 64)
		var date = time.Unix(int64(dateInt), 0)
		return &date
	}
	return nil
}

// NewUsersDetailsDBModule returns a UsersDB module.
func NewUsersDetailsDBModule(db sqltypes.CloudtrustDB, cipher security.EncrypterDecrypter, logger log.Logger) UsersDetailsDBModule {
	return &usersDBModule{
		db:     db,
		cipher: cipher,
		logger: logger,
	}
}

func (c *usersDBModule) StoreOrUpdateUserDetails(ctx context.Context, realm string, user dto.DBUser) error {
	// transform user object into JSON string
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	// encrypt the JSON containing the details on the user
	encryptedData, err := c.cipher.Encrypt(userJSON, []byte(*user.UserID))
	if err != nil {
		c.logger.Warn(ctx, "msg", "Can't encrypt the user details", "err", err.Error(), "realmID", realm, "userID", &user.UserID)
		return err
	}

	keyID := c.cipher.GetCurrentKeyID()
	// update value in DB
	_, err = c.db.Exec(updateUserDetailsStmt, realm, user.UserID, encryptedData, keyID, encryptedData, keyID)
	return err
}

func (c *usersDBModule) GetUserDetails(ctx context.Context, realm string, userID string) (dto.DBUser, error) {
	var encryptedDetails []byte
	var keyID string
	var details = dto.DBUser{}
	row := c.db.QueryRow(selectUserDetailsStmt, realm, userID)

	switch err := row.Scan(&encryptedDetails, &keyID); err {
	case sql.ErrNoRows:
		return dto.DBUser{
			UserID: &userID,
		}, nil
	default:
		if err != nil {
			return dto.DBUser{}, err
		}
		//decrypt the user details & unmarshal
		detailsJSON, err := c.cipher.Decrypt(encryptedDetails, keyID, []byte(userID))
		if err != nil {
			c.logger.Warn(ctx, "msg", "Can't decrypt the user details", "err", err.Error(), "realmID", realm, "userID", userID)
			return dto.DBUser{}, err
		}
		err = json.Unmarshal(detailsJSON, &details)
		details.UserID = &userID
		return details, err
	}
}

func (c *usersDBModule) DeleteUserDetails(ctx context.Context, realm string, userID string) error {
	_, err := c.db.Exec(deleteUserDetailsStmt, realm, userID)
	return err
}

func (c *usersDBModule) CreateCheck(ctx context.Context, realm string, userID string, check dto.DBCheck) error {
	var proofData *[]byte
	var err error

	if check.ProofData != nil {
		// encrypt the proof data & protect integrity of userID associated to the proof data
		encryptedData, err := c.cipher.Encrypt(*check.ProofData, []byte(userID))
		if err != nil {
			c.logger.Warn(ctx, "msg", "Can't encrypt the proof data", "err", err.Error(), "realmID", realm, "userID", userID)
			return err
		}
		proofData = &encryptedData
	}

	// insert check in DB
	keyID := c.cipher.GetCurrentKeyID()
	_, err = c.db.Exec(createCheckStmt, realm, userID, check.Operator,
		check.DateTime, check.Status, check.Type, check.Nature,
		check.ProofType, proofData, check.Comment, keyID, check.TxnID,
		realm, userID, check.Operator, check.DateTime, check.Status, check.Type,
		check.Nature, check.ProofType, proofData, check.Comment, keyID, check.TxnID)

	return err
}

func (c *usersDBModule) CreatePendingCheck(ctx context.Context, realm string, userID string, check dto.DBCheck) error {
	var proofData *[]byte
	var err error

	if check.ProofData != nil {
		// encrypt the proof data & protect integrity of userID associated to the proof data
		encryptedData, err := c.cipher.Encrypt(*check.ProofData, []byte(userID))
		if err != nil {
			c.logger.Warn(ctx, "msg", "Can't encrypt the proof data", "err", err.Error(), "realmID", realm, "userID", userID)
			return err
		}
		proofData = &encryptedData
	}

	// insert check in DB
	keyID := c.cipher.GetCurrentKeyID()
	_, err = c.db.Exec(createPendingCheckStmt, realm, userID, check.Operator,
		check.DateTime, check.Type, check.Nature,
		check.ProofType, proofData, check.Comment, keyID, check.TxnID)

	return err
}

func (c *usersDBModule) GetChecks(ctx context.Context, realm string, userID string) ([]dto.DBCheck, error) {
	var rows, err = c.db.Query(selectCheckStmt, realm, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	return c.scanChecks(ctx, rows)
}

func (c *usersDBModule) GetPendingChecks(ctx context.Context, realm string, userID string) ([]dto.DBCheck, error) {
	var rows, err = c.db.Query(selectPendingCheckStmt, realm, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	return c.scanChecks(ctx, rows)
}

func (c *usersDBModule) scanChecks(ctx context.Context, rows sqltypes.SQLRows) ([]dto.DBCheck, error) {
	var result []dto.DBCheck
	var checkID int64
	var operator, datetime, status, checkType, nature, proofType, comment, txnID sql.NullString
	var encryptedProofData []byte
	var realm, userID, keyID string
	var err error

	for rows.Next() {
		err = rows.Scan(&checkID, &realm, &userID, &operator, &datetime, &status, &checkType, &nature, &proofType, &encryptedProofData, &comment, &keyID, &txnID)
		if err != nil {
			return nil, err
		}

		var proofData []byte

		if len(encryptedProofData) != 0 {
			//decrypt the proof data of the user
			proofData, err = c.cipher.Decrypt(encryptedProofData, keyID, []byte(userID))
			if err != nil {
				c.logger.Warn(ctx, "msg", "Can't decrypt the proof data", "err", err.Error(), "realmID", realm, "userID", userID)
				return nil, err
			}
		}

		result = append(result, dto.DBCheck{
			Operator:  nullStringToPtr(operator),
			DateTime:  nullStringToDatePtr(datetime),
			Status:    nullStringToPtr(status),
			Type:      nullStringToPtr(checkType),
			Nature:    nullStringToPtr(nature),
			ProofData: &proofData,
			ProofType: nullStringToPtr(proofType),
			Comment:   nullStringToPtr(comment),
			TxnID:     nullStringToPtr(txnID),
		})
	}

	return result, err
}
