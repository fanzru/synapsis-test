// Code generated by MockGen. DO NOT EDIT.
// Source: app/account/repo/mysql.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	models "synapsis-test/app/account/domain/models"
	param "synapsis-test/app/account/domain/param"
	response "synapsis-test/app/account/domain/response"

	gomock "github.com/golang/mock/gomock"
)

// MockImpl is a mock of Impl interface.
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImplMockRecorder
}

// MockImplMockRecorder is the mock recorder for MockImpl.
type MockImplMockRecorder struct {
	mock *MockImpl
}

// NewMockImpl creates a new mock instance.
func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &MockImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImpl) EXPECT() *MockImplMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockImpl) CreateUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockImplMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockImpl)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockImpl) DeleteUser(ctx context.Context, findParam param.FindParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, findParam)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockImplMockRecorder) DeleteUser(ctx, findParam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockImpl)(nil).DeleteUser), ctx, findParam)
}

// FindFirstUser mocks base method.
func (m *MockImpl) FindFirstUser(ctx context.Context, user *models.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFirstUser", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFirstUser indicates an expected call of FindFirstUser.
func (mr *MockImplMockRecorder) FindFirstUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFirstUser", reflect.TypeOf((*MockImpl)(nil).FindFirstUser), ctx, user)
}

// FindUser mocks base method.
func (m *MockImpl) FindUser(ctx context.Context, user *models.User, param param.FindParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, user, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindUser indicates an expected call of FindUser.
func (mr *MockImplMockRecorder) FindUser(ctx, user, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockImpl)(nil).FindUser), ctx, user, param)
}

// FindUserWithPagination mocks base method.
func (m *MockImpl) FindUserWithPagination(ctx context.Context, param param.FindParam) (response.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserWithPagination", ctx, param)
	ret0, _ := ret[0].(response.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserWithPagination indicates an expected call of FindUserWithPagination.
func (mr *MockImplMockRecorder) FindUserWithPagination(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserWithPagination", reflect.TypeOf((*MockImpl)(nil).FindUserWithPagination), ctx, param)
}

// FindUserWithRole mocks base method.
func (m *MockImpl) FindUserWithRole(ctx context.Context, user *models.User, param param.FindParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserWithRole", ctx, user, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindUserWithRole indicates an expected call of FindUserWithRole.
func (mr *MockImplMockRecorder) FindUserWithRole(ctx, user, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserWithRole", reflect.TypeOf((*MockImpl)(nil).FindUserWithRole), ctx, user, param)
}

// RunInTransaction mocks base method.
func (m *MockImpl) RunInTransaction(ctx context.Context, f func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockImplMockRecorder) RunInTransaction(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockImpl)(nil).RunInTransaction), ctx, f)
}

// UpdateUser mocks base method.
func (m *MockImpl) UpdateUser(ctx context.Context, updateParam param.UpdateParam, findParam param.FindParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, updateParam, findParam)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockImplMockRecorder) UpdateUser(ctx, updateParam, findParam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockImpl)(nil).UpdateUser), ctx, updateParam, findParam)
}
