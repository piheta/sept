package main

import (
	"fmt"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

type App struct {
	user_repo     *repos.UserRepo
	chat_repo     *repos.ChatRepo
	userchat_repo *repos.UserchatRepo
	message_repo  *repos.MessageRepo

	auth_service *services.AuthService
}

func NewApp(userRepo *repos.UserRepo, chatRepo *repos.ChatRepo, userchatRepo *repos.UserchatRepo, messageRepo *repos.MessageRepo, authService *services.AuthService) *App {
	return &App{
		user_repo:     userRepo,
		chat_repo:     chatRepo,
		userchat_repo: userchatRepo,
		message_repo:  messageRepo,

		auth_service: authService,
	}
}

///
/// AUTH ENDPOINTS
///

func (a *App) GetAuthedUser() models.User {
	err := a.auth_service.LogInWithExistingJwt()
	if err != nil {
		fmt.Println(err)
	}

	return services.AuthedUser
}

func (a *App) Login(email, password string) (*models.User, error) {
	user, err := a.auth_service.Login(email, password)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to login %w", err)
	}

	a.user_repo = repos.NewUserRepo(db.DB)
	a.chat_repo = repos.NewChatRepo(db.DB)
	a.userchat_repo = repos.NewUserchatRepo(db.DB)
	a.message_repo = repos.NewMessageRepo(db.DB)

	return user, nil
}

func (a *App) Register(username, email, password string) (*map[string]interface{}, error) {
	return a.auth_service.Register(username, email, password)
}

func (a *App) LogOut() error {
	err := a.auth_service.LogOut()
	if err != nil {
		return err
	}

	a.user_repo = repos.NewUserRepo(nil)
	a.chat_repo = repos.NewChatRepo(nil)
	a.userchat_repo = repos.NewUserchatRepo(nil)
	a.message_repo = repos.NewMessageRepo(nil)

	services.AuthedUser = models.User{}

	return nil
}

///
/// APP ENDPOINTS
///

func (a *App) GetUser(user_id string) (models.User, error) {
	return a.user_repo.GetUser(user_id)
}

func (a *App) GetUsers() ([]models.User, error) {
	users, err := a.user_repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return users, nil
}

func (a *App) SendMessage(message string, chat_id string, user_id string) ([]models.Message, error) {
	err := a.message_repo.AddMessage(chat_id, user_id, message)
	if err != nil {
		return nil, fmt.Errorf("failed to add message: %w", err)
	}
	return a.message_repo.GetMessagesByChatID(chat_id)
}

func (a *App) GetChatMessages(chat_id string) ([]models.Message, error) {
	return a.message_repo.GetMessagesByChatID(chat_id)
}
