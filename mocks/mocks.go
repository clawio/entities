package mocks

import (
	"github.com/clawio/entities"
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

// MockObjectInfo mocks an MockObjectInfo.
type MockObjectInfo struct {
	mock.Mock
	Otype    entities.ObjectType `json:"type"`
	PathSpec string              `json:"pathspec"`
	Size     uint64              `json:"size"`
	Checksum string              `json:"checksum"`
	MimeType string              `json:"mime"`
}

//GetType mocks the GetType call.
func (o *MockObjectInfo) GetType() entities.ObjectType { return o.Otype }

// GetPathSpec mocks the GetPathSpec call.
func (o *MockObjectInfo) GetPathSpec() string { return o.PathSpec }

// GetSize mocks the GetSize call.
func (o *MockObjectInfo) GetSize() uint64 { return o.Size }

// GetChecksum mocks the GetChecksum call.
func (o *MockObjectInfo) GetChecksum() string { return o.Checksum }

// GetMimeType mocks the GetMimeType call.
func (o *MockObjectInfo) GetMimeType() string { return o.MimeType }
