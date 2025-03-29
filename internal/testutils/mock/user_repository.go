package mock

import (
	"math/rand"

	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
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

	return entities.User{}, driverErrors.ErrRecordNotFound
}
func (r MockUserRepository) FindUser(userID int64) (entities.User, error) {
	res, ok := r.List[userID]
	if !ok {
		return entities.User{}, driverErrors.ErrRecordNotFound
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
