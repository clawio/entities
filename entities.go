package entities

// User represents an user of the system.
// They are created by the authentication service.
type User interface {
	GetUsername() string
	GetEmail() string
	GetDisplayName() string
}
