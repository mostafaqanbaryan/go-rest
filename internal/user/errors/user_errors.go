package userErrors

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUsernameTaken   = errors.New("username taken")
	ErrPasswordIsWrong = errors.New("password is wrong")
)
