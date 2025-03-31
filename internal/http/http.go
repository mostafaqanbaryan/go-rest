package http

import (
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Register(echo.Context) error
	Logout(echo.Context) error
	Login(echo.Context) error
}

type UserHandler interface {
	Me(echo.Context) error
	Update(echo.Context) error
}

func NewServer(authHandler AuthHandler, userHandler UserHandler) {
	e := echo.New()

	authGroup := e.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/logout", authHandler.Logout)

	e.GET("/me", userHandler.Me)
	e.POST("/me", userHandler.Update)

	e.Logger.Fatal(e.Start(":3000"))
}
