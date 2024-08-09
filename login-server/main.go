package main

import (
	"github.com/piheta/sept/login-server/config"
	"github.com/piheta/sept/login-server/controllers"
	"github.com/piheta/sept/login-server/repos"
	"github.com/piheta/sept/login-server/services"
)

// @Title sept-login-server
// @Version 0.1
func main() {
	config.Connect()

	// Initialize repos
	userRepo := repos.NewUserRepo(config.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo, authService)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	app := config.NewRouter(
		authController,
		userController,
	)

	app.Listen(":8080")
}
