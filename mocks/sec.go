// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Sutheres/report-chaser/internal/sec (interfaces: SEC)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockSEC is a mock of SEC interface
type MockSEC struct {
	ctrl     *gomock.Controller
	recorder *MockSECMockRecorder
}

// MockSECMockRecorder is the mock recorder for MockSEC
type MockSECMockRecorder struct {
	mock *MockSEC
}

// NewMockSEC creates a new mock instance
func NewMockSEC(ctrl *gomock.Controller) *MockSEC {
	mock := &MockSEC{ctrl: ctrl}
	mock.recorder = &MockSECMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSEC) EXPECT() *MockSECMockRecorder {
	return m.recorder
}
