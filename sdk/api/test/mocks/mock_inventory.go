// Code generated by MockGen. DO NOT EDIT.
// Source: ./inventory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	api "github.com/dhimasan0206/adapter/sdk/api"
	gomock "github.com/golang/mock/gomock"
)

// MockInventoryRepository is a mock of InventoryRepository interface.
type MockInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepositoryMockRecorder
}

// MockInventoryRepositoryMockRecorder is the mock recorder for MockInventoryRepository.
type MockInventoryRepositoryMockRecorder struct {
	mock *MockInventoryRepository
}

// NewMockInventoryRepository creates a new mock instance.
func NewMockInventoryRepository(ctrl *gomock.Controller) *MockInventoryRepository {
	mock := &MockInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryRepository) EXPECT() *MockInventoryRepositoryMockRecorder {
	return m.recorder
}

// UpdateStockLevel mocks base method.
func (m *MockInventoryRepository) UpdateStockLevel(ctx context.Context, req api.UpdateStockLevelRequest) (*api.UpdateStockLevelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockLevel", ctx, req)
	ret0, _ := ret[0].(*api.UpdateStockLevelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStockLevel indicates an expected call of UpdateStockLevel.
func (mr *MockInventoryRepositoryMockRecorder) UpdateStockLevel(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockLevel", reflect.TypeOf((*MockInventoryRepository)(nil).UpdateStockLevel), ctx, req)
}

// MockInventoryService is a mock of InventoryService interface.
type MockInventoryService struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServiceMockRecorder
}

// MockInventoryServiceMockRecorder is the mock recorder for MockInventoryService.
type MockInventoryServiceMockRecorder struct {
	mock *MockInventoryService
}

// NewMockInventoryService creates a new mock instance.
func NewMockInventoryService(ctrl *gomock.Controller) *MockInventoryService {
	mock := &MockInventoryService{ctrl: ctrl}
	mock.recorder = &MockInventoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryService) EXPECT() *MockInventoryServiceMockRecorder {
	return m.recorder
}

// UpdateStockLevel mocks base method.
func (m *MockInventoryService) UpdateStockLevel(ctx context.Context, req api.UpdateStockLevelRequest) (*api.UpdateStockLevelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockLevel", ctx, req)
	ret0, _ := ret[0].(*api.UpdateStockLevelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStockLevel indicates an expected call of UpdateStockLevel.
func (mr *MockInventoryServiceMockRecorder) UpdateStockLevel(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockLevel", reflect.TypeOf((*MockInventoryService)(nil).UpdateStockLevel), ctx, req)
}
