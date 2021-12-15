package user

import "strings"

type User struct {
	Email      string
	FamilyName string
	GivenName  string
}

func (u *User) FullName() string {
	return strings.Trim(strings.Join([]string{u.FamilyName, u.GivenName}, " "), " ")
}
