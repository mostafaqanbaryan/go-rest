package userhttp_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/argon2"
	authservice "mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/testutils/mock"
	"mostafaqanbaryan.com/go-rest/internal/testutils/request"
	userhandler "mostafaqanbaryan.com/go-rest/internal/user/http"
	userservice "mostafaqanbaryan.com/go-rest/internal/user/service"
	"mostafaqanbaryan.com/go-rest/pkg/validation"
)

var (
	sessionID = "test123"
	user      = entities.User{
		ID:       1,
		Email:    "test@test.com",
		Password: "test",
	}
)
var hash, _ = argon2.CreateHash(user.Password)

var validator = validation.NewValidator()

var authRepository = mock.MockAuthRepository{
	List: map[string]int64{
		sessionID: user.ID,
	},
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
var userHandler = userhandler.NewUserHandler(authService, userService)

func TestUserHandler_Me(t *testing.T) {

	t.Run("Fail /me with no cookie", func(t *testing.T) {
		c, _ := request.NewGETRequest(e, "/", "")
		err := userHandler.Me(c)
		if err == nil {
			t.Fatalf("want error, got none")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "named cookie not present"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Fail /me with wrong cookie", func(t *testing.T) {
		c, _ := request.NewGETRequest(e, "/", "wrong_session_test")
		err := userHandler.Me(c)
		if err == nil {
			t.Fatalf("want error, got none")
		}

		var httpError *echo.HTTPError
		if !errors.As(err, &httpError) {
			t.Fatalf("want error to be of type echo.HTTPError, got %T", err)
		}

		if httpError.Code != http.StatusUnauthorized {
			t.Fatalf("want status code %d, got %d", http.StatusUnauthorized, httpError.Code)
		}

		want := "record not found"
		if !strings.Contains(httpError.Message.(error).Error(), want) {
			t.Fatalf("want %s, got %s,", want, httpError.Message)
		}
	})

	t.Run("Success /me", func(t *testing.T) {
		c, rec := request.NewGETRequest(e, "/me", sessionID)
		if err := userHandler.Me(c); err != nil {
			t.Fatalf("want no error, got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("want status code %d, got %d", http.StatusOK, rec.Code)
		}

		var tmp entities.User
		json.Unmarshal(rec.Body.Bytes(), &tmp)
		if user.Email != tmp.Email {
			t.Fatalf("want email %s, got %s", user.Email, tmp.Email)
		}
	})
}

func TestUserHandler_Update(t *testing.T) {
	t.Run("Update fullname with short name", func(t *testing.T) {
		fullname := strings.Repeat("a", 2)

		c, _ := request.NewPATCHRequest(e, "/me", "{\"fullname\":\""+fullname+"\"}", sessionID)
		err := userHandler.Update(c)
		if err == nil {
			t.Fatalf("want error, got none")
		}

		want := "'min' tag"
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("want %s, got %s,", want, err)
		}

		_, rec := request.NewGETRequest(e, "/", sessionID)

		var tmp entities.User
		json.Unmarshal(rec.Body.Bytes(), &tmp)
		if tmp.Fullname != user.Fullname {
			t.Fatalf("want fullname %s, got: <%v>", user.Fullname, tmp.Fullname)
		}
	})

	t.Run("Update fullname with long name", func(t *testing.T) {
		fullname := strings.Repeat("a", 256)

		c, _ := request.NewPOSTRequest(e, "/", "{\"fullname\":\""+fullname+"\"}", sessionID)
		err := userHandler.Update(c)
		if err == nil {
			t.Fatalf("want error, got none")
		}

		want := "'max' tag"
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("want %s, got %v,", want, err)
		}

		_, rec := request.NewGETRequest(e, "/", sessionID)

		var tmp entities.User
		json.Unmarshal(rec.Body.Bytes(), &tmp)
		if tmp.Fullname != user.Fullname {
			t.Fatalf("want fullname %s, got: <%v>", user.Fullname, tmp.Fullname)
		}
	})

	t.Run("Update fullname with valid characters", func(t *testing.T) {
		fullname := "test test-ts jr."

		c, _ := request.NewPOSTRequest(e, "/", "{\"fullname\":\""+fullname+"\"}", sessionID)
		err := userHandler.Update(c)
		if err != nil {
			t.Fatalf("want no error, got %v", err)
		}

		c, rec := request.NewGETRequest(e, "/me", sessionID)
		if err := userHandler.Me(c); err != nil {
			t.Fatalf("want no error, got %v", err)
		}

		var tmp entities.User
		json.Unmarshal(rec.Body.Bytes(), &tmp)
		if tmp.Fullname != fullname {
			t.Fatalf("want fullname %s, got: <%v>", fullname, tmp.Fullname)
		}
	})
}
