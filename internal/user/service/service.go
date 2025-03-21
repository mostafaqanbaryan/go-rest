package service

import (
	"mostafaqanbaryan.com/go-rest/internal/argon2"
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
	if user.ID > 0 || err == nil {
		return userErrors.ErrUsernameTaken
	}

	if err != driverErrors.ErrRecordNotFound {
		return err
	}

	encrypted, err := argon2.CreateHash(password)
	if err != nil {
		return err
	}
	return s.repo.Create(username, encrypted)

}

func (s userService) Login(username, password string) (entities.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return entities.User{}, userErrors.ErrUserNotFound
	}

	match, err := argon2.CompareHash(password, user.Password)
	if err != nil || !match {
		return entities.User{}, userErrors.ErrPasswordIsWrong
	}

	return user, nil
}

func (s userService) Find(userID int64) (entities.User, error) {
	return s.repo.FindUser(userID)
}
