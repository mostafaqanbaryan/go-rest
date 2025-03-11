package services

import (
	"errors"
	"testing"

	cerrors "mostafaqanbaryan.com/go-rest/internal/errors"
	"mostafaqanbaryan.com/go-rest/internal/repositories"
)

func TestUserService(t *testing.T) {
	username := "test"
	password := "test"

	userRepository := repositories.NewUserRepositoryMock()
	userService := NewUserService(userRepository)

	t.Run("User not found", func(t *testing.T) {
		_, err := userService.Login("notfound", password)
		if !errors.Is(err, cerrors.ErrUserNotFound) {
			t.Fatalf("want <%v>, got: <%v>", cerrors.ErrUserNotFound, err)
		}
	})

	t.Run("User password is wrong", func(t *testing.T) {
		_, err := userService.Login(username, "wrongpassword")
		if !errors.Is(err, cerrors.ErrPasswordIsWrong) {
			t.Fatalf("want <%v>, got: <%v>", cerrors.ErrPasswordIsWrong, err)
		}
	})

	t.Run("User found", func(t *testing.T) {
		user, err := userService.Login(username, password)
		if err != nil {
			t.Fatalf("want nil, got: <%v>", err)
		}
		if user.Username != "test" {
			t.Fatalf("want username %s, got: <%v>", username, user.Username)
		}
	})
}
