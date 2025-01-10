// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/account (interfaces: Component,UserProfileCache)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=./mock/component.go -package=mock -mock_names=Component=Component,UserProfileCache=UserProfileCache github.com/cloudtrust/keycloak-bridge/pkg/account Component,UserProfileCache
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	apiaccount "github.com/cloudtrust/keycloak-bridge/api/account"
	apicommon "github.com/cloudtrust/keycloak-bridge/api/common"
	keycloak "github.com/cloudtrust/keycloak-client/v2"
	gomock "go.uber.org/mock/gomock"
)

// Component is a mock of Component interface.
type Component struct {
	ctrl     *gomock.Controller
	recorder *ComponentMockRecorder
	isgomock struct{}
}

// ComponentMockRecorder is the mock recorder for Component.
type ComponentMockRecorder struct {
	mock *Component
}

// NewComponent creates a new mock instance.
func NewComponent(ctrl *gomock.Controller) *Component {
	mock := &Component{ctrl: ctrl}
	mock.recorder = &ComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Component) EXPECT() *ComponentMockRecorder {
	return m.recorder
}

// CancelEmailChange mocks base method.
func (m *Component) CancelEmailChange(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelEmailChange", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelEmailChange indicates an expected call of CancelEmailChange.
func (mr *ComponentMockRecorder) CancelEmailChange(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelEmailChange", reflect.TypeOf((*Component)(nil).CancelEmailChange), ctx)
}

// CancelPhoneNumberChange mocks base method.
func (m *Component) CancelPhoneNumberChange(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelPhoneNumberChange", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelPhoneNumberChange indicates an expected call of CancelPhoneNumberChange.
func (mr *ComponentMockRecorder) CancelPhoneNumberChange(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPhoneNumberChange", reflect.TypeOf((*Component)(nil).CancelPhoneNumberChange), ctx)
}

// DeleteAccount mocks base method.
func (m *Component) DeleteAccount(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *ComponentMockRecorder) DeleteAccount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*Component)(nil).DeleteAccount), arg0)
}

// DeleteCredential mocks base method.
func (m *Component) DeleteCredential(ctx context.Context, credentialID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCredential", ctx, credentialID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredential indicates an expected call of DeleteCredential.
func (mr *ComponentMockRecorder) DeleteCredential(ctx, credentialID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredential", reflect.TypeOf((*Component)(nil).DeleteCredential), ctx, credentialID)
}

// GetAccount mocks base method.
func (m *Component) GetAccount(ctx context.Context) (apiaccount.AccountRepresentation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx)
	ret0, _ := ret[0].(apiaccount.AccountRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *ComponentMockRecorder) GetAccount(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*Component)(nil).GetAccount), ctx)
}

// GetConfiguration mocks base method.
func (m *Component) GetConfiguration(arg0 context.Context, arg1 string) (apiaccount.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration", arg0, arg1)
	ret0, _ := ret[0].(apiaccount.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *ComponentMockRecorder) GetConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*Component)(nil).GetConfiguration), arg0, arg1)
}

// GetCredentialRegistrators mocks base method.
func (m *Component) GetCredentialRegistrators(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialRegistrators", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialRegistrators indicates an expected call of GetCredentialRegistrators.
func (mr *ComponentMockRecorder) GetCredentialRegistrators(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialRegistrators", reflect.TypeOf((*Component)(nil).GetCredentialRegistrators), ctx)
}

// GetCredentials mocks base method.
func (m *Component) GetCredentials(ctx context.Context) ([]apiaccount.CredentialRepresentation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", ctx)
	ret0, _ := ret[0].([]apiaccount.CredentialRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *ComponentMockRecorder) GetCredentials(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*Component)(nil).GetCredentials), ctx)
}

// GetUserProfile mocks base method.
func (m *Component) GetUserProfile(arg0 context.Context) (apicommon.ProfileRepresentation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", arg0)
	ret0, _ := ret[0].(apicommon.ProfileRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *ComponentMockRecorder) GetUserProfile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*Component)(nil).GetUserProfile), arg0)
}

// MoveCredential mocks base method.
func (m *Component) MoveCredential(ctx context.Context, credentialID, previousCredentialID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveCredential", ctx, credentialID, previousCredentialID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveCredential indicates an expected call of MoveCredential.
func (mr *ComponentMockRecorder) MoveCredential(ctx, credentialID, previousCredentialID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveCredential", reflect.TypeOf((*Component)(nil).MoveCredential), ctx, credentialID, previousCredentialID)
}

// SendVerifyEmail mocks base method.
func (m *Component) SendVerifyEmail(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerifyEmail", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerifyEmail indicates an expected call of SendVerifyEmail.
func (mr *ComponentMockRecorder) SendVerifyEmail(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerifyEmail", reflect.TypeOf((*Component)(nil).SendVerifyEmail), ctx)
}

// SendVerifyPhoneNumber mocks base method.
func (m *Component) SendVerifyPhoneNumber(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerifyPhoneNumber", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerifyPhoneNumber indicates an expected call of SendVerifyPhoneNumber.
func (mr *ComponentMockRecorder) SendVerifyPhoneNumber(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerifyPhoneNumber", reflect.TypeOf((*Component)(nil).SendVerifyPhoneNumber), ctx)
}

// UpdateAccount mocks base method.
func (m *Component) UpdateAccount(arg0 context.Context, arg1 apiaccount.UpdatableAccountRepresentation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *ComponentMockRecorder) UpdateAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*Component)(nil).UpdateAccount), arg0, arg1)
}

// UpdateLabelCredential mocks base method.
func (m *Component) UpdateLabelCredential(ctx context.Context, credentialID, label string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLabelCredential", ctx, credentialID, label)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabelCredential indicates an expected call of UpdateLabelCredential.
func (mr *ComponentMockRecorder) UpdateLabelCredential(ctx, credentialID, label any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabelCredential", reflect.TypeOf((*Component)(nil).UpdateLabelCredential), ctx, credentialID, label)
}

// UpdatePassword mocks base method.
func (m *Component) UpdatePassword(ctx context.Context, currentPassword, newPassword, confirmPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, currentPassword, newPassword, confirmPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *ComponentMockRecorder) UpdatePassword(ctx, currentPassword, newPassword, confirmPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*Component)(nil).UpdatePassword), ctx, currentPassword, newPassword, confirmPassword)
}

// UserProfileCache is a mock of UserProfileCache interface.
type UserProfileCache struct {
	ctrl     *gomock.Controller
	recorder *UserProfileCacheMockRecorder
	isgomock struct{}
}

// UserProfileCacheMockRecorder is the mock recorder for UserProfileCache.
type UserProfileCacheMockRecorder struct {
	mock *UserProfileCache
}

// NewUserProfileCache creates a new mock instance.
func NewUserProfileCache(ctrl *gomock.Controller) *UserProfileCache {
	mock := &UserProfileCache{ctrl: ctrl}
	mock.recorder = &UserProfileCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserProfileCache) EXPECT() *UserProfileCacheMockRecorder {
	return m.recorder
}

// GetRealmUserProfile mocks base method.
func (m *UserProfileCache) GetRealmUserProfile(ctx context.Context, realmName string) (keycloak.UserProfileRepresentation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRealmUserProfile", ctx, realmName)
	ret0, _ := ret[0].(keycloak.UserProfileRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRealmUserProfile indicates an expected call of GetRealmUserProfile.
func (mr *UserProfileCacheMockRecorder) GetRealmUserProfile(ctx, realmName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRealmUserProfile", reflect.TypeOf((*UserProfileCache)(nil).GetRealmUserProfile), ctx, realmName)
}
