package main

import (
	"context"
	"fmt"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

type App struct {
	ctx           context.Context
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

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	//get latest version
}

func (a *App) StartLoggedIn(user_id string) error {
	if err := db.DbExists(user_id); err != nil {
		return fmt.Errorf("db does not exist %w", err)
	}

	db.InitDb(user_id)
	return nil
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

func (a *App) GetUser(user_id string) (models.User, error) {
	return a.user_repo.GetUser(user_id)
}

func (a *App) Login(email, password string) (*models.User, error) {
	// if database does not exist (first time login)
	user, err := a.auth_service.Login(email, password)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to login %w", err)
	}

	if err := a.handleFirstTimeLogin(*user); err != nil {
		return &models.User{}, fmt.Errorf("failed to init db %w", err)
	}

	return user, nil
}

func (a *App) Register(username, email, password string) (*map[string]interface{}, error) {
	return a.auth_service.Register(username, email, password)
}

// helper
// function
func (a *App) handleFirstTimeLogin(user models.User) error {
	// Initialize the user's database based on their ID
	if err := db.InitDb(user.ID); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Reinit app fields
	a.user_repo = repos.NewUserRepo(db.DB)
	a.chat_repo = repos.NewChatRepo(db.DB)
	a.userchat_repo = repos.NewUserchatRepo(db.DB)
	a.message_repo = repos.NewMessageRepo(db.DB)

	// Add the user to the database
	if err := a.user_repo.AddUser(user); err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	// Add a new chat for the user
	if err := a.chat_repo.AddChat(user.Username); err != nil {
		return fmt.Errorf("failed to add chat: %w", err)
	}

	// Retrieve the created chat by username
	chat, err := a.chat_repo.GetChatByName(user.Username)
	if err != nil {
		return fmt.Errorf("failed to get chat: %w", err)
	}
	fmt.Println(chat)
	// Add the user to the chat
	if err := a.userchat_repo.AddUserToChat(user.ID, chat.ID); err != nil {
		return fmt.Errorf("failed to add user to chat: %w", err)
	}

	return nil
}
