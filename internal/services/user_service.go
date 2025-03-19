package services

import (
	"errors"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exists")
	ErrPasswordIsWrong  = errors.New("password is wrong")
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
		return entities.User{}, ErrUserNotFound
	}

	if user.Password != password {
		return entities.User{}, ErrPasswordIsWrong
	}

	return user, nil
}

func (s UserService) FindBySessionID(user, password string) (entities.User, error) {
	return s.repo.FindByUsername(user)
}
