package http

import (
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Logout(echo.Context) error
	Login(echo.Context) error
}

type UserHandler interface {
	Me(echo.Context) error
}

func NewServer(authHandler AuthHandler, userHandler UserHandler) {
	e := echo.New()

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/logout", authHandler.Logout)

	e.GET("/me", userHandler.Me)
	// e.PATCH("/me", authHandler.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
