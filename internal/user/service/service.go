package userservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"mostafaqanbaryan.com/go-rest/internal/argon2"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailTaken         = errors.New("email is taken")
	ErrPasswordIsWrong    = errors.New("password is wrong")
	ErrPasswordValidation = errors.New("password is weak")
)

type userRepository interface {
	Create(string, string, string) error
	FindUser(int64) (entities.User, error)
	FindByEmail(string) (entities.User, error)
	IsDuplicateEmail(string) (bool, error)
	Update(int64, string) error
}
type userService struct {
	validator *validator.Validate
	repo      userRepository
}

func NewUserService(validator *validator.Validate, userRepository userRepository) userService {
	return userService{
		validator: validator,
		repo:      userRepository,
	}
}

func (s userService) Register(email, password string) error {
	user := entities.User{
		Email:    email,
		Password: password,
	}
	if err := s.validator.Struct(user); err != nil {
		return err
	}

	duplicate, err := s.repo.IsDuplicateEmail(email)
	if err != nil {
		return err
	}

	if duplicate {
		return ErrEmailTaken
	}

	encrypted, err := argon2.CreateHash(password)
	if err != nil {
		return err
	}

	hashID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return s.repo.Create(hashID.String(), email, encrypted)

}

func (s userService) Login(email, password string) (entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return entities.User{}, ErrUserNotFound
	}

	match, err := argon2.CompareHash(password, user.Password)
	if err != nil || !match {
		return entities.User{}, ErrPasswordIsWrong
	}

	return user, nil
}

func (s userService) Update(userID int64, params entities.User) error {
	_, err := s.repo.FindUser(userID)
	if err != nil {
		return err
	}

	if err := s.validator.Struct(params); err != nil {
		return err
	}

	return s.repo.Update(userID, params.Fullname)
}

func (s userService) Find(userID int64) (entities.User, error) {
	return s.repo.FindUser(userID)
}
