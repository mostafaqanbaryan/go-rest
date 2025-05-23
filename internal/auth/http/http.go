package authhttp

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type authService interface {
	CreateSession(user entities.User) (string, error)
}

type userService interface {
	Login(email, password string) (entities.User, error)
	Register(email, password string) error
}

type AuthHandler struct {
	userService userService
	authService authService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(authService authService, userService userService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	v := LoginRequest{}
	if err := c.Bind(&v); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	err := h.userService.Register(v.Email, v.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return nil
}

func (h *AuthHandler) Login(c echo.Context) error {
	v := LoginRequest{}
	if err := c.Bind(&v); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
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
