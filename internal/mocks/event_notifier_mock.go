// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-issue-api/internal/events (interfaces: EventNotifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventNotifier is a mock of EventNotifier interface.
type MockEventNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockEventNotifierMockRecorder
}

// MockEventNotifierMockRecorder is the mock recorder for MockEventNotifier.
type MockEventNotifierMockRecorder struct {
	mock *MockEventNotifier
}

// NewMockEventNotifier creates a new mock instance.
func NewMockEventNotifier(ctrl *gomock.Controller) *MockEventNotifier {
	mock := &MockEventNotifier{ctrl: ctrl}
	mock.recorder = &MockEventNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventNotifier) EXPECT() *MockEventNotifierMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventNotifier) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEventNotifierMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventNotifier)(nil).Close))
}

// Notify mocks base method.
func (m *MockEventNotifier) Notify(arg0 uint64, arg1 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", arg0, arg1)
}

// Notify indicates an expected call of Notify.
func (mr *MockEventNotifierMockRecorder) Notify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockEventNotifier)(nil).Notify), arg0, arg1)
}
