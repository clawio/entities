package entities

type User interface {
	Username() string
	Email() string
	DisplayName() string
}

func NewSimpleUser(name, email, displayName string) User {
	return &simpleUser{
		Name:  name,
		Mail:  email,
		DName: displayName,
	}
}

type simpleUser struct {
	Name  string `json:"username"`
	Mail  string `json:"email"`
	DName string `json:"display_name"`
}

func (u *simpleUser) Username() string {
	return u.Name
}
func (u *simpleUser) Email() string {
	return u.Mail
}
func (u *simpleUser) DisplayName() string {
	return u.DName
}
