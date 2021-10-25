// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "github.com/bariasabda/monorepo/packages/fetch/domain/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetResource mocks base method
func (m *MockRepositoryInterface) GetResource() (*[]entity.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource")
	ret0, _ := ret[0].(*[]entity.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource
func (mr *MockRepositoryInterfaceMockRecorder) GetResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockRepositoryInterface)(nil).GetResource))
}

// CurrencyConverter mocks base method
func (m *MockRepositoryInterface) CurrencyConverter(from, to string) (*entity.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrencyConverter", from, to)
	ret0, _ := ret[0].(*entity.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrencyConverter indicates an expected call of CurrencyConverter
func (mr *MockRepositoryInterfaceMockRecorder) CurrencyConverter(from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrencyConverter", reflect.TypeOf((*MockRepositoryInterface)(nil).CurrencyConverter), from, to)
}