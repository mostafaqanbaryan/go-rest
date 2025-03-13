package services

import (
	"mostafaqanbaryan.com/go-rest/internal/entities"
	cerrors "mostafaqanbaryan.com/go-rest/internal/errors"
)

type UserRepository interface {
	FindByUsername(username string) (entities.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return UserService{
		repo: userRepository,
	}
}

func (s UserService) Register(username, password string) (entities.User, error) {
	return s.repo.FindByUsername(username)
}

func (s UserService) Login(username, password string) (entities.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return entities.User{}, cerrors.ErrNotFound
	}

	if user.Password != password {
		return entities.User{}, cerrors.ErrPasswordIsWrong
	}

	return user, nil
}

func (s UserService) FindBySessionID(user, password string) (entities.User, error) {
	return s.repo.FindByUsername(user)
}
