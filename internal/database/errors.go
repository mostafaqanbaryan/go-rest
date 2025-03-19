package database

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrGetCommand     = errors.New("GET command error")
	ErrSetCommand     = errors.New("SET command error")
)
