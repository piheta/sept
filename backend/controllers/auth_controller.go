package controllers

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/services"
)

type AuthController struct {
	ctx          context.Context
	auth_service *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		auth_service: authService,
	}
}

func (a *AuthController) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *AuthController) Exit() {
	runtime.Quit(a.ctx)
}

///
/// AUTH ENDPOINTS
///

func (a *AuthController) GetAuthedUser() models.User {
	return services.AuthedUser
}

func (a *AuthController) Login(email, password string) (*models.User, error) {
	user, err := a.auth_service.Login(email, password)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to login: %w", err)
	}

	return user, nil
}

func (a *AuthController) Register(username, email, password string) (*map[string]interface{}, error) {
	return a.auth_service.Register(username, email, password)
}

func (a *AuthController) LogOut() error {
	err := a.auth_service.LogOut()
	if err != nil {
		return err
	}

	services.AuthedUser = models.User{}

	return nil
}

///
/// APP ENDPOINTS
///

func (a *AuthController) Search(searchString string) ([]string, error) {
	return db.Search(searchString)
}
