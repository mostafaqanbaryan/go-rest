package mock

import (
	"errors"
	"math/rand"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type MockUserRepository struct {
	List map[int64]*entities.User
}

func (r MockUserRepository) FindByEmail(email string) (entities.User, error) {
	for _, user := range r.List {
		if email == user.Email {
			return *user, nil
		}
	}

	return entities.User{}, errors.New("user not found")
}

func (r MockUserRepository) IsDuplicateEmail(email string) (bool, error) {
	for _, user := range r.List {
		if email == user.Email {
			return true, nil
		}
	}

	return false, nil
}

func (r MockUserRepository) FindUser(userID int64) (entities.User, error) {
	res, ok := r.List[userID]
	if !ok {
		return entities.User{}, errors.New("userid not found")
	}
	return *res, nil
}

func (r MockUserRepository) Create(hashID, email, password string) error {
	userID := rand.Int63()
	r.List[userID] = &entities.User{
		ID:       userID,
		HashID:   hashID,
		Email:    email,
		Password: password,
	}
	return nil
}

func (r MockUserRepository) Update(userID int64, fullname string) error {
	r.List[userID].Fullname = fullname
	return nil
}
