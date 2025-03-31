package userhttp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type AuthService interface {
	GetSession(string) (int64, error)
}

type UserService interface {
	Find(int64) (entities.User, error)
	Update(int64, entities.User) error
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

type UpdateRequest struct {
	Fullname string
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

	user, err := h.userService.Find(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	v := UpdateRequest{}
	if err = c.Bind(&v); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.authService.GetSession(token.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return h.userService.Update(userID, entities.User{
		Fullname: v.Fullname,
	})
}
