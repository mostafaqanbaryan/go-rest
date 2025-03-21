package services

import (
	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type UserRepository interface {
	Create(string, string) error
	FindByUsername(string) (entities.User, error)
}
type UserService struct {
	repo UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return UserService{
		repo: userRepository,
	}
}

func (s UserService) Register(username, password string) error {
	user, err := s.repo.FindByUsername(username)
	if user.ID > 0 {
		return ErrUsernameTaken
	}

	if err == database.ErrRecordNotFound {
		return s.repo.Create(username, password)
	}

	return err
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
