package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type AuthService interface {
	GetSession(string) (int64, error)
}

type UserService interface {
	FindByID(int64) (entities.User, error)
}

type UserHandler struct {
	userService UserService
	authService AuthService
}

func NewUserHandler(authService AuthService, userService UserService) UserHandler {
	return UserHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *UserHandler) Me(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.authService.GetSession(token.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := h.userService.FindByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, user)
}
