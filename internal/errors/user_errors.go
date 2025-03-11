package errors

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExist = errors.New("user already exists")
var ErrPasswordIsWrong = errors.New("password is wrong")
