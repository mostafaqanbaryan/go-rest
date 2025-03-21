package services

import (
	"errors"
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/repositories"
)

func TestUserService(t *testing.T) {
	user := entities.User{
		Username: "test",
		Password: "tset",
	}

	db := database.NewMockDatabaseDriver()

	userRepository := repositories.NewUserRepository(db)
	userService := NewUserService(userRepository)

	// Initialize
	err := userService.Register(user.Username, user.Password)
	if err != nil {
		t.Fatalf("register wants no error, got: <%v>", err)
	}

	t.Run("Username is taken", func(t *testing.T) {
		err := userService.Register(user.Username, user.Password)
		if err != ErrUsernameTaken {
			t.Fatalf("want <%v>, got: <%v>", ErrUsernameTaken, err)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		_, err := userService.Login("notfound", user.Password)
		if !errors.Is(err, ErrUserNotFound) {
			t.Fatalf("want <%v>, got: <%v>", ErrUserNotFound, err)
		}
	})

	t.Run("User password is wrong", func(t *testing.T) {
		_, err := userService.Login(user.Username, "wrongpassword")
		if !errors.Is(err, ErrPasswordIsWrong) {
			t.Fatalf("want <%v>, got: <%v>", ErrPasswordIsWrong, err)
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
