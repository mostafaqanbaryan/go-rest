package user_test

import (
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/driver"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/user/repository"
	"mostafaqanbaryan.com/go-rest/internal/user/service"
)

func TestUserHandler(t *testing.T) {
	t.Parallel()

	user := entities.User{
		Username: "test",
		Password: "tset",
	}
	db := driver.NewMockDatabaseDriver()
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	// Initialize
	err := userService.Register(user.Username, user.Password)
	if err != nil {
		t.Fatalf("register wants no error, got: <%v>", err)
	}

	// Create session
	user, err = userService.Login(user.Username, user.Password)
	if err != nil {
		t.Fatalf("login wants no error, got: <%v>", err)
	}

	t.Run("Get Me", func(t *testing.T) {
		found, err := userService.Find(user.ID)
		if err != nil {
			t.Fatalf("want user, got: <%v>", err)
		}

		if user.ID != found.ID {
			t.Fatalf("want userID 1, got: <%v>", found.ID)
		}
	})
}
