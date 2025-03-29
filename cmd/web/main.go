package main

import (
	_ "github.com/joho/godotenv/autoload"
	authhttp "mostafaqanbaryan.com/go-rest/internal/auth/http"
	authrepo "mostafaqanbaryan.com/go-rest/internal/auth/repository"
	authservice "mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/driver"
	"mostafaqanbaryan.com/go-rest/internal/http"
	userhttp "mostafaqanbaryan.com/go-rest/internal/user/http"
	userrepo "mostafaqanbaryan.com/go-rest/internal/user/repository"
	userservice "mostafaqanbaryan.com/go-rest/internal/user/service"
	"mostafaqanbaryan.com/go-rest/pkg/validation"
)

func main() {
	cache := driver.NewRedisDriver()
	defer cache.Close()

	db := driver.NewMySQLDriver("")
	defer db.Close()

	validator := validation.NewValidator()

	authRepository := authrepo.NewAuthRepository(cache)
	authService := authservice.NewAuthService(authRepository)

	userRepository := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(validator, userRepository)

	authHandler := authhttp.NewAuthHandler(authService, userService)
	userHandler := userhttp.NewUserHandler(authService, userService)

	http.NewServer(&authHandler, &userHandler)
}
