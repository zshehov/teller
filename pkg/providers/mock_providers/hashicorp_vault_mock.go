// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/providers/hashicorp_vault.go

// Package mock_providers is a generated GoMock package.
package mock_providers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/vault/api"
)

// MockHashicorpClient is a mock of HashicorpClient interface.
type MockHashicorpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHashicorpClientMockRecorder
}

// MockHashicorpClientMockRecorder is the mock recorder for MockHashicorpClient.
type MockHashicorpClientMockRecorder struct {
	mock *MockHashicorpClient
}

// NewMockHashicorpClient creates a new mock instance.
func NewMockHashicorpClient(ctrl *gomock.Controller) *MockHashicorpClient {
	mock := &MockHashicorpClient{ctrl: ctrl}
	mock.recorder = &MockHashicorpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashicorpClient) EXPECT() *MockHashicorpClientMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockHashicorpClient) Read(path string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", path)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockHashicorpClientMockRecorder) Read(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockHashicorpClient)(nil).Read), path)
}