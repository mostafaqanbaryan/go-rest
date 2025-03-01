package services

import (
	"errors"
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/repositories"
)

func TestUserService(t *testing.T) {
	t.Run("Login not found", func(t *testing.T) {
		userRepository := repositories.NewUserRepositoryMock()
		userService := NewUserService(userRepository)
		_, err := userService.Login("test", "test")
		if !errors.Is(err, UserNotFound{}) {
			t.Fatalf("want <%v>, got: <%v>", UserNotFound{}, err)
		}
	})
}
