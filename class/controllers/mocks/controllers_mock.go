// Code generated by MockGen. DO NOT EDIT.
// Source: controllers.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	reflect "reflect"

	dto "github.com/JimySheepman/class/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockadminServices is a mock of adminServices interface.
type MockadminServices struct {
	ctrl     *gomock.Controller
	recorder *MockadminServicesMockRecorder
}

// MockadminServicesMockRecorder is the mock recorder for MockadminServices.
type MockadminServicesMockRecorder struct {
	mock *MockadminServices
}

// NewMockadminServices creates a new mock instance.
func NewMockadminServices(ctrl *gomock.Controller) *MockadminServices {
	mock := &MockadminServices{ctrl: ctrl}
	mock.recorder = &MockadminServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockadminServices) EXPECT() *MockadminServicesMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockadminServices) Read(req *dto.RestRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockadminServicesMockRecorder) Read(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockadminServices)(nil).Read), req)
}

// Write mocks base method.
func (m *MockadminServices) Write(req *dto.RestRequest) (*dto.RestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", req)
	ret0, _ := ret[0].(*dto.RestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockadminServicesMockRecorder) Write(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockadminServices)(nil).Write), req)
}