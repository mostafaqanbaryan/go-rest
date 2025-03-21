package service

import (
	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	userErrors "mostafaqanbaryan.com/go-rest/internal/user/errors"
)

type userRepository interface {
	Create(string, string) error
	FindUser(int64) (entities.User, error)
	FindByUsername(string) (entities.User, error)
}
type userService struct {
	repo userRepository
}

func NewUserService(userRepository userRepository) userService {
	return userService{
		repo: userRepository,
	}
}

func (s userService) Register(username, password string) error {
	user, err := s.repo.FindByUsername(username)
	if user.ID > 0 {
		return userErrors.ErrUsernameTaken
	}

	if err == driverErrors.ErrRecordNotFound {
		return s.repo.Create(username, password)
	}

	return err
}

func (s userService) Login(username, password string) (entities.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return entities.User{}, userErrors.ErrUserNotFound
	}

	if user.Password != password {
		return entities.User{}, userErrors.ErrPasswordIsWrong
	}

	return user, nil
}

func (s userService) Find(userID int64) (entities.User, error) {
	return s.repo.FindUser(userID)
}
