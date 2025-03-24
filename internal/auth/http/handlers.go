package http

import (
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type authService interface {
	CreateSession(user entities.User) (string, error)
}

type userService interface {
	Login(email, password string) (entities.User, error)
}

type AuthHandler struct {
	userService userService
	authService authService
	validator   *validator.Validate
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

func NewAuthHandler(validator *validator.Validate, authService authService, userService userService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
		validator:   validator,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	v := LoginRequest{}
	if err := c.Bind(&v); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if err := h.validator.Struct(v); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := h.userService.Login(v.Email, v.Password)
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

func (h *AuthHandler) Logout(c echo.Context) error {
	cookie := createCookie("", time.Time{})
	c.SetCookie(&cookie)
	return nil
}

func createCookie(value string, expire time.Time) http.Cookie {
	return http.Cookie{
		Name:     "token",
		Value:    value,
		Domain:   os.Getenv("APP_DOMAIN"),
		Path:     "/",
		Expires:  expire,
		HttpOnly: true,
		Secure:   true,
	}
}
