// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/eitanya/go/src/github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes/kube_namespace.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	namespace "github.com/solo-io/solo-kit/api/external/kubernetes/namespace"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// MockCloneableKubeNamespace is a mock of CloneableKubeNamespace interface
type MockCloneableKubeNamespace struct {
	ctrl     *gomock.Controller
	recorder *MockCloneableKubeNamespaceMockRecorder
}

// MockCloneableKubeNamespaceMockRecorder is the mock recorder for MockCloneableKubeNamespace
type MockCloneableKubeNamespaceMockRecorder struct {
	mock *MockCloneableKubeNamespace
}

// NewMockCloneableKubeNamespace creates a new mock instance
func NewMockCloneableKubeNamespace(ctrl *gomock.Controller) *MockCloneableKubeNamespace {
	mock := &MockCloneableKubeNamespace{ctrl: ctrl}
	mock.recorder = &MockCloneableKubeNamespaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloneableKubeNamespace) EXPECT() *MockCloneableKubeNamespaceMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method
func (m *MockCloneableKubeNamespace) GetMetadata() core.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata")
	ret0, _ := ret[0].(core.Metadata)
	return ret0
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockCloneableKubeNamespaceMockRecorder) GetMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockCloneableKubeNamespace)(nil).GetMetadata))
}

// SetMetadata mocks base method
func (m *MockCloneableKubeNamespace) SetMetadata(meta core.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMetadata", meta)
}

// SetMetadata indicates an expected call of SetMetadata
func (mr *MockCloneableKubeNamespaceMockRecorder) SetMetadata(meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadata", reflect.TypeOf((*MockCloneableKubeNamespace)(nil).SetMetadata), meta)
}

// Equal mocks base method
func (m *MockCloneableKubeNamespace) Equal(that interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", that)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockCloneableKubeNamespaceMockRecorder) Equal(that interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockCloneableKubeNamespace)(nil).Equal), that)
}

// Clone mocks base method
func (m *MockCloneableKubeNamespace) Clone() *namespace.KubeNamespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(*namespace.KubeNamespace)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockCloneableKubeNamespaceMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockCloneableKubeNamespace)(nil).Clone))
}