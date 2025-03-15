package main

import (
	"github.com/labstack/echo/v4"
	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/handlers"
	"mostafaqanbaryan.com/go-rest/internal/repositories"
	"mostafaqanbaryan.com/go-rest/internal/services"
)

func main() {
	e := echo.New()

	cache := database.NewRedisDriver()
	defer cache.Close()

	db := database.NewMySQLDriver("")
	defer db.Close()

	database.MigrateUp(db)
	conn := entities.New(db)

	authRepository := repositories.NewAuthRepositoryCache(cache)
	authService := services.NewAuthService(authRepository)

	userRepository := repositories.NewUserRepositoryDB(conn)
	userService := services.NewUserService(userRepository)
	authHandler := handlers.NewAuthHandler(authService, userService)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	// authGroup.GET("/logout", authHandler.Logout)

	e.Logger.Fatal(e.Start(":3000"))
}
