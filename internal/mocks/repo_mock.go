// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-issue-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	issues "github.com/ozoncp/ocp-issue-api/internal/issues"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddIssue mocks base method.
func (m *MockRepo) AddIssue(arg0 issues.Issue) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIssue", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIssue indicates an expected call of AddIssue.
func (mr *MockRepoMockRecorder) AddIssue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIssue", reflect.TypeOf((*MockRepo)(nil).AddIssue), arg0)
}

// AddIssues mocks base method.
func (m *MockRepo) AddIssues(arg0 []issues.Issue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIssues", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddIssues indicates an expected call of AddIssues.
func (mr *MockRepoMockRecorder) AddIssues(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIssues", reflect.TypeOf((*MockRepo)(nil).AddIssues), arg0)
}

// DescribeIssue mocks base method.
func (m *MockRepo) DescribeIssue(arg0 uint64) (*issues.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeIssue", arg0)
	ret0, _ := ret[0].(*issues.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIssue indicates an expected call of DescribeIssue.
func (mr *MockRepoMockRecorder) DescribeIssue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIssue", reflect.TypeOf((*MockRepo)(nil).DescribeIssue), arg0)
}

// ListIssues mocks base method.
func (m *MockRepo) ListIssues(arg0, arg1 uint64) ([]issues.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIssues", arg0, arg1)
	ret0, _ := ret[0].([]issues.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIssues indicates an expected call of ListIssues.
func (mr *MockRepoMockRecorder) ListIssues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssues", reflect.TypeOf((*MockRepo)(nil).ListIssues), arg0, arg1)
}

// RemoveIssue mocks base method.
func (m *MockRepo) RemoveIssue(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIssue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIssue indicates an expected call of RemoveIssue.
func (mr *MockRepoMockRecorder) RemoveIssue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIssue", reflect.TypeOf((*MockRepo)(nil).RemoveIssue), arg0)
}
