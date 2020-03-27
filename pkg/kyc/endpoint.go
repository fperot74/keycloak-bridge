package kyc

import (
	"context"

	cs "github.com/cloudtrust/common-service"
	commonerrors "github.com/cloudtrust/common-service/errors"
	apikyc "github.com/cloudtrust/keycloak-bridge/api/kyc"
	msg "github.com/cloudtrust/keycloak-bridge/internal/constants"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints for self service
type Endpoints struct {
	GetActions        endpoint.Endpoint
	GetUser           endpoint.Endpoint
	GetUserByUsername endpoint.Endpoint
	ValidateUser      endpoint.Endpoint
}

// MakeGetActionsEndpoint creates an endpoint for GetActions
func MakeGetActionsEndpoint(component Component) cs.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return component.GetActions(ctx)
	}
}

// MakeGetUserByUsernameEndpoint endpoint creation
func MakeGetUserByUsernameEndpoint(component Component) cs.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var m = req.(map[string]string)
		var user = m["username"]

		return component.GetUserByUsername(ctx, user)
	}
}

// MakeGetUserEndpoint endpoint creation
func MakeGetUserEndpoint(component Component) cs.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var m = req.(map[string]string)
		var user = m["userId"]
		return component.GetUser(ctx, user)
	}
}

// MakeValidateUserEndpoint endpoint creation
func MakeValidateUserEndpoint(component Component) cs.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var m = req.(map[string]string)
		var userID = m["userId"]
		var user, err = apikyc.UserFromJSON(m["body"])
		if err != nil {
			return nil, commonerrors.CreateBadRequestError(commonerrors.MsgErrInvalidParam + "." + msg.BodyContent)
		}
		return nil, component.ValidateUser(ctx, userID, user)
	}
}
