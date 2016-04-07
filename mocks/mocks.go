package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockUser mocks a User
type MockUser struct {
	mock.Mock
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

// GetUsername mocks the GetUsername call.
func (m *MockUser) GetUsername() string {
	args := m.Called()
	return args.String(0)
}

// GetEmail mocks the GetEmail call.
func (m *MockUser) GetEmail() string {
	args := m.Called()
	return args.String(0)
}

// GetDisplayName mocks the GetDisplayName call.
func (m *MockUser) GetDisplayName() string {
	args := m.Called()
	return args.String(0)
}
