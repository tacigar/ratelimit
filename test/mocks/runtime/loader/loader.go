// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lyft/goruntime/loader (interfaces: IFace)

// Package mock_loader is a generated GoMock package.
package mock_loader

import (
	gomock "github.com/golang/mock/gomock"
	snapshot "github.com/lyft/goruntime/snapshot"
	reflect "reflect"
)

// MockIFace is a mock of IFace interface
type MockIFace struct {
	ctrl     *gomock.Controller
	recorder *MockIFaceMockRecorder
}

// MockIFaceMockRecorder is the mock recorder for MockIFace
type MockIFaceMockRecorder struct {
	mock *MockIFace
}

// NewMockIFace creates a new mock instance
func NewMockIFace(ctrl *gomock.Controller) *MockIFace {
	mock := &MockIFace{ctrl: ctrl}
	mock.recorder = &MockIFaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFace) EXPECT() *MockIFaceMockRecorder {
	return m.recorder
}

// AddUpdateCallback mocks base method
func (m *MockIFace) AddUpdateCallback(arg0 chan<- int) {
	m.ctrl.Call(m, "AddUpdateCallback", arg0)
}

// AddUpdateCallback indicates an expected call of AddUpdateCallback
func (mr *MockIFaceMockRecorder) AddUpdateCallback(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpdateCallback", reflect.TypeOf((*MockIFace)(nil).AddUpdateCallback), arg0)
}

// Snapshot mocks base method
func (m *MockIFace) Snapshot() snapshot.IFace {
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(snapshot.IFace)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockIFaceMockRecorder) Snapshot() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockIFace)(nil).Snapshot))
}
