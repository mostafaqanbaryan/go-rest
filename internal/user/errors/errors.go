package usererrors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailTaken         = errors.New("email is taken")
	ErrPasswordIsWrong    = errors.New("password is wrong")
	ErrPasswordValidation = errors.New("password is weak")
)
