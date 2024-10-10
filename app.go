package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/handlers"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

type App struct {
	ctx context.Context

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
}

func (a *App) Exit() {
	runtime.Quit(a.ctx)
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

func (a *App) GetChats() ([]models.Chat, error) {
	chats, err := a.chat_repo.GetChats()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chats: %w", err)
	}
	return chats, nil
}

func (a *App) SendMessage(content string, chat_id string) ([]models.Message, error) {

	msg := models.Message{
		ChatID:  chat_id,
		UserID:  services.AuthedUser.ID,
		Content: content,
	}
	signedMsg, err := services.SignMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message %w, ", err)
	}

	err = a.message_repo.AddMessage(signedMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return a.message_repo.GetMessagesByChatID(chat_id)
}

func (a *App) GetChatMessages(chat_id string) ([]models.Message, error) {
	return a.message_repo.GetMessagesByChatID(chat_id)
}

func (a *App) GetIps() []string {
	return handlers.Ips
}

func (a *App) Search(searchString string) ([]string, error) {
	return db.Search(searchString)
}

//
// SIGNALLING
//

func (a *App) JoinSignallingServer(jwt string) error {

	return nil
}
