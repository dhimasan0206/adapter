// Code generated by MockGen. DO NOT EDIT.
// Source: ./store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	api "github.com/dhimasan0206/adapter/sdk/api"
	gomock "github.com/golang/mock/gomock"
)

// MockStoreRepository is a mock of StoreRepository interface.
type MockStoreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStoreRepositoryMockRecorder
}

// MockStoreRepositoryMockRecorder is the mock recorder for MockStoreRepository.
type MockStoreRepositoryMockRecorder struct {
	mock *MockStoreRepository
}

// NewMockStoreRepository creates a new mock instance.
func NewMockStoreRepository(ctrl *gomock.Controller) *MockStoreRepository {
	mock := &MockStoreRepository{ctrl: ctrl}
	mock.recorder = &MockStoreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreRepository) EXPECT() *MockStoreRepositoryMockRecorder {
	return m.recorder
}

// GetByCode mocks base method.
func (m *MockStoreRepository) GetByCode(ctx context.Context, code string) (*api.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(*api.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode.
func (mr *MockStoreRepositoryMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockStoreRepository)(nil).GetByCode), ctx, code)
}

// GetByID mocks base method.
func (m *MockStoreRepository) GetByID(ctx context.Context, id string) (*api.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*api.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStoreRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStoreRepository)(nil).GetByID), ctx, id)
}

// MockStoreService is a mock of StoreService interface.
type MockStoreService struct {
	ctrl     *gomock.Controller
	recorder *MockStoreServiceMockRecorder
}

// MockStoreServiceMockRecorder is the mock recorder for MockStoreService.
type MockStoreServiceMockRecorder struct {
	mock *MockStoreService
}

// NewMockStoreService creates a new mock instance.
func NewMockStoreService(ctrl *gomock.Controller) *MockStoreService {
	mock := &MockStoreService{ctrl: ctrl}
	mock.recorder = &MockStoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreService) EXPECT() *MockStoreServiceMockRecorder {
	return m.recorder
}

// GetByCode mocks base method.
func (m *MockStoreService) GetByCode(ctx context.Context, code string) (*api.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(*api.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode.
func (mr *MockStoreServiceMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockStoreService)(nil).GetByCode), ctx, code)
}

// GetByID mocks base method.
func (m *MockStoreService) GetByID(ctx context.Context, id string) (*api.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*api.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStoreServiceMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStoreService)(nil).GetByID), ctx, id)
}
