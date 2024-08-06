package main

import (
	"github.com/piheta/sept/config"
	"github.com/piheta/sept/controllers"
	"github.com/piheta/sept/repos"
	"github.com/piheta/sept/services"
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
