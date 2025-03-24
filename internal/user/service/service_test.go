package service_test

import (
	"errors"
	"strings"
	"testing"

	"math/rand"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	userErrors "mostafaqanbaryan.com/go-rest/internal/user/errors"
	"mostafaqanbaryan.com/go-rest/internal/user/service"

	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
)

type MockUserRepository struct {
	list map[int64]*entities.User
}

func newMockUserRepository() MockUserRepository {
	return MockUserRepository{
		list: make(map[int64]*entities.User, 0),
	}
}

func (r MockUserRepository) FindByEmail(email string) (entities.User, error) {
	for _, user := range r.list {
		if email == user.Email {
			return *user, nil
		}
	}

	return entities.User{}, driverErrors.ErrRecordNotFound
}
func (r MockUserRepository) FindUser(userID int64) (entities.User, error) {
	res, ok := r.list[userID]
	if !ok {
		return entities.User{}, driverErrors.ErrRecordNotFound
	}
	return *res, nil
}

func (r MockUserRepository) Create(hashID, email, password string) error {
	userID := rand.Int63()
	r.list[userID] = &entities.User{
		ID:       userID,
		HashID:   hashID,
		Email:    email,
		Password: password,
	}
	return nil
}

func (r MockUserRepository) Update(userID int64, fullname string) error {
	r.list[userID].Fullname = fullname
	return nil
}

func TestUserService(t *testing.T) {
	t.Parallel()

	user := entities.User{
		Email:    "test",
		Password: "tset",
	}

	userRepository := newMockUserRepository()
	userService := service.NewUserService(userRepository)

	// Initialize
	err := userService.Register(user.Email, user.Password)
	if err != nil {
		t.Fatalf("register wants no error, got: <%v>", err)
	}

	t.Run("Email is taken", func(t *testing.T) {
		err := userService.Register(user.Email, user.Password)
		if err != userErrors.ErrEmailTaken {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrEmailTaken, err)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		_, err := userService.Login("notfound", user.Password)
		if !errors.Is(err, userErrors.ErrUserNotFound) {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrUserNotFound, err)
		}
	})

	t.Run("User password is wrong", func(t *testing.T) {
		_, err := userService.Login(user.Email, "wrongpassword")
		if !errors.Is(err, userErrors.ErrPasswordIsWrong) {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrPasswordIsWrong, err)
		}
	})

	t.Run("User found", func(t *testing.T) {
		found, err := userService.Login(user.Email, user.Password)
		if err != nil {
			t.Fatalf("want nil, got: <%v>", err)
		}

		if found.Email != user.Email {
			t.Fatalf("want email %s, got: <%v>", found.Email, user.Email)
		}

		if found.ID == 0 {
			t.Fatalf("want an id, got: <%v>", user.ID)
		}
	})
}
