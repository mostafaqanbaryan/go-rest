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
		Username: "test",
		Password: "tset",
	}

	db := driver.NewMockDatabaseDriver()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	// Initialize
	err := userService.Register(user.Username, user.Password)
	if err != nil {
		t.Fatalf("register wants no error, got: <%v>", err)
	}

	t.Run("Username is taken", func(t *testing.T) {
		err := userService.Register(user.Username, user.Password)
		if err != userErrors.ErrUsernameTaken {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrUsernameTaken, err)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		_, err := userService.Login("notfound", user.Password)
		if !errors.Is(err, userErrors.ErrUserNotFound) {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrUserNotFound, err)
		}
	})

	t.Run("User password is wrong", func(t *testing.T) {
		_, err := userService.Login(user.Username, "wrongpassword")
		if !errors.Is(err, userErrors.ErrPasswordIsWrong) {
			t.Fatalf("want <%v>, got: <%v>", userErrors.ErrPasswordIsWrong, err)
		}
	})

	t.Run("User found", func(t *testing.T) {
		found, err := userService.Login(user.Username, user.Password)
		if err != nil {
			t.Fatalf("want nil, got: <%v>", err)
		}

		if found.Username != user.Username {
			t.Fatalf("want username %s, got: <%v>", found.Username, user.Username)
		}

		if found.ID == 0 {
			t.Fatalf("want an id, got: <%v>", user.ID)
		}
	})
}
