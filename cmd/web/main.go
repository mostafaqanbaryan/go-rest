package main

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"mostafaqanbaryan.com/go-rest/internal/handlers"
	"mostafaqanbaryan.com/go-rest/internal/repositories"
	"mostafaqanbaryan.com/go-rest/internal/services"
)

func main() {
	e := echo.New()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	authRepository := repositories.NewAuthRepositoryCache(rdb)
	authService := services.NewAuthService(authRepository)

	userRepository := repositories.NewUserRepositoryDB()
	userService := services.NewUserService(userRepository)
	authHandler := handlers.NewAuthHandler(authService, userService)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	// authGroup.GET("/logout", authHandler.Logout)

	e.Logger.Fatal(e.Start(":3000"))
}
