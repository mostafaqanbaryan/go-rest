package userservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"mostafaqanbaryan.com/go-rest/internal/argon2"
	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/user/errors"
)

type userRepository interface {
	Create(string, string, string) error
	FindUser(int64) (entities.User, error)
	FindByEmail(string) (entities.User, error)
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

	duplicate, err := s.repo.FindByEmail(email)
	if duplicate.ID > 0 || err == nil {
		return usererrors.ErrEmailTaken
	}

	if err != driverErrors.ErrRecordNotFound {
		return err
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
		return entities.User{}, usererrors.ErrUserNotFound
	}

	match, err := argon2.CompareHash(password, user.Password)
	if err != nil || !match {
		return entities.User{}, usererrors.ErrPasswordIsWrong
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
