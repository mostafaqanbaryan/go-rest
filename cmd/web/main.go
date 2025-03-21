package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	authhandler "mostafaqanbaryan.com/go-rest/internal/auth/http"
	authrepo "mostafaqanbaryan.com/go-rest/internal/auth/repository"
	authservice "mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/driver"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	userhandler "mostafaqanbaryan.com/go-rest/internal/user/http"
	userrepo "mostafaqanbaryan.com/go-rest/internal/user/repository"
	userservice "mostafaqanbaryan.com/go-rest/internal/user/service"
)

func main() {
	e := echo.New()

	cache := driver.NewRedisDriver()
	defer cache.Close()

	db := driver.NewMySQLDriver("")
	defer db.Close()

	database.MigrateUp(db)
	conn := entities.New(db)

	authRepository := authrepo.NewAuthRepository(cache)
	authService := authservice.NewAuthService(authRepository)

	userRepository := userrepo.NewUserRepository(conn)
	userService := userservice.NewUserService(userRepository)

	authHandler := authhandler.NewAuthHandler(authService, userService)
	userHandler := userhandler.NewUserHandler(authService, userService)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/logout", authHandler.Logout)

	e.GET("/me", userHandler.Me)
	// e.PATCH("/me", authHandler.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
