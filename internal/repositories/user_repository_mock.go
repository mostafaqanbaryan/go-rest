package repositories

import (
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/errors"
)

type UserRepositoryMock struct {
}

func NewUserRepositoryMock() UserRepositoryMock {
	return UserRepositoryMock{}
}

func (r UserRepositoryMock) FindByUsername(username string) (*entities.User, error) {
	return nil, errors.UserNotFound{}
}
