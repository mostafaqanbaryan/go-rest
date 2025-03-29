package authhttp_test

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/argon2"
	authhandler "mostafaqanbaryan.com/go-rest/internal/auth/http"
	authservice "mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/testutils/mock"
	"mostafaqanbaryan.com/go-rest/internal/testutils/request"
	userservice "mostafaqanbaryan.com/go-rest/internal/user/service"
	"mostafaqanbaryan.com/go-rest/pkg/validation"
)

var (
	user = entities.User{
		ID:       1,
		Email:    "test@test.com",
		Password: "test",
	}
)
var hash, _ = argon2.CreateHash(user.Password)

var validator = validation.NewValidator()

var authRepository = mock.MockAuthRepository{
	List: map[string]int64{},
}
var authService = authservice.NewAuthService(authRepository)

var userRepository = mock.MockUserRepository{
	List: map[int64]*entities.User{
		user.ID: {
			Email:    user.Email,
			Password: hash,
		},
	},
}

var e = echo.New()
var userService = userservice.NewUserService(validator, userRepository)
var authHandler = authhandler.NewAuthHandler(authService, userService)

func TestAuthHandler_Register(t *testing.T) {
	t.Run("Email is taken", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, user.Password)
		c, _ := request.NewPOSTRequest(e, "/", body, "")
		err := authHandler.Register(c)
		if err == nil {
			t.Fatalf("want error, got nothing")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "email is taken"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Password is weak", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, user.Password)
		c, _ := request.NewPOSTRequest(e, "/", body, "")
		err := authHandler.Register(c)
		if err == nil {
			t.Fatalf("want error, got nothing")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "email is taken"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Successful registration", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, "new@gorest.me", "wrong_password")
		c, _ := request.NewPOSTRequest(e, "/", body, "")
		err := authHandler.Register(c)
		if err != nil {
			t.Fatalf("want error, got nothing")
		}
	})

}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("Login with wrong password", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, "wrong_password")
		c, _ := request.NewPOSTRequest(e, "/", body, "")
		err := authHandler.Login(c)
		if err == nil {
			t.Fatalf("want error, got nothing")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "password is wrong"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Login with wrong email", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, "wrong_email", user.Password)
		c, _ := request.NewPOSTRequest(e, "/", body, "")
		err := authHandler.Login(c)
		if err == nil {
			t.Fatalf("want error, got nothing")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "user not found"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Successful login", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, user.Password)
		c, rec := request.NewPOSTRequest(e, "/", body, "")
		if err := authHandler.Login(c); err != nil {
			t.Fatalf("want no error, got %v", err)
		}

		if rec.Result().Cookies()[0].Name != "token" {
			t.Fatalf("want token, got %s", rec.Result().Cookies()[0])
		}
	})

}
