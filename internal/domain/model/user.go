package model

import "errors"

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Username string
	Password string
	Name     string
	Surname1 string
	Surname2 string
}
