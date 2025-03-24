package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	authhandler "mostafaqanbaryan.com/go-rest/internal/auth/http"
	authrepo "mostafaqanbaryan.com/go-rest/internal/auth/repository"
	authservice "mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/driver"
	userhandler "mostafaqanbaryan.com/go-rest/internal/user/http"
	userrepo "mostafaqanbaryan.com/go-rest/internal/user/repository"
	userservice "mostafaqanbaryan.com/go-rest/internal/user/service"
	"mostafaqanbaryan.com/go-rest/pkg/validation"
)

func main() {
	e := echo.New()

	cache := driver.NewRedisDriver()
	defer cache.Close()

	db := driver.NewMySQLDriver("")
	defer db.Close()

	validator := validation.NewValidator()

	authRepository := authrepo.NewAuthRepository(cache)
	authService := authservice.NewAuthService(authRepository)

	userRepository := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepository)

	authHandler := authhandler.NewAuthHandler(validator, authService, userService)
	userHandler := userhandler.NewUserHandler(authService, userService)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/logout", authHandler.Logout)

	e.GET("/me", userHandler.Me)
	// e.PATCH("/me", authHandler.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
