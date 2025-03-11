package repositories

import (
	"mostafaqanbaryan.com/go-rest/internal/entities"
	cerrors "mostafaqanbaryan.com/go-rest/internal/errors"
)

type UserRepositoryMock struct {
}

func NewUserRepositoryMock() UserRepositoryMock {
	return UserRepositoryMock{}
}

func (r UserRepositoryMock) FindByUsername(username string) (entities.User, error) {
	if username == "test" {
		return entities.User{
			Username: "test",
			Password: "test",
		}, nil
	}
	return entities.User{}, cerrors.ErrUserNotFound
}
