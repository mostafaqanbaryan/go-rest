package service_test

import (
	"errors"
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/driver"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	userErrors "mostafaqanbaryan.com/go-rest/internal/user/errors"
	"mostafaqanbaryan.com/go-rest/internal/user/repository"
	"mostafaqanbaryan.com/go-rest/internal/user/service"
)

func TestUserService(t *testing.T) {
	t.Parallel()

	user := entities.User{
		Email:    "test",
		Password: "tset",
	}

	db := driver.NewMockDatabaseDriver()

	userRepository := repository.NewUserRepository(db)
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
