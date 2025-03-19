package repositories

import (
	"errors"

	"mostafaqanbaryan.com/go-rest/internal/entities"
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
	return entities.User{}, errors.New("user not found")
}
