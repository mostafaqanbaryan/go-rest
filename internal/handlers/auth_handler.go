package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type AuthService interface {
	CreateSession(user entities.User) (string, error)
}

type UserService interface {
	Login(username, password string) (entities.User, error)
}

type AuthHandler struct {
	userService UserService
	authService AuthService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(authService AuthService, userService UserService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	v := LoginRequest{}
	if err := c.Bind(&v); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := h.userService.Login(v.Username, v.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	sessionId, err := h.authService.CreateSession(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	cookie := createCookie(sessionId, time.Now().Add(time.Hour*24*365))
	c.SetCookie(&cookie)
	return c.JSON(http.StatusNoContent, "")
}

func (h *AuthHandler) Logout(c echo.Context) {
	cookie := createCookie("", time.Time{})
	c.SetCookie(&cookie)
}

func createCookie(value string, expire time.Time) http.Cookie {
	return http.Cookie{
		Name:     "token",
		Value:    value,
		Path:     "/",
		Expires:  expire,
		HttpOnly: true,
		Secure:   true,
	}
}
