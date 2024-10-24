package controllers

import (
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
)

type UserController struct {
	user_repo *repos.UserRepo
}

func NewUserController(userRepo *repos.UserRepo) *UserController {
	return &UserController{
		user_repo: userRepo,
	}
}

func (uc *UserController) GetUser(user_id string) (models.User, error) {
	return uc.user_repo.GetUser(user_id)
}

func (uc *UserController) GetUsers() ([]models.User, error) {
	users, err := uc.user_repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return users, nil
}
