package main

import (
	"context"
	"fmt"

	"github.com/piheta/sept/infra/auth"
	"github.com/piheta/sept/infra/db"
	"github.com/piheta/sept/models"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	//get latest version
}

func (a *App) GetUsers() ([]models.User, error) {
	users, err := db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return users, nil
}

func (a *App) SendMessage(message string, chat_id int) ([]models.Message, error) {
	err := db.AddMessage(chat_id, 1, message)
	if err != nil {
		return nil, fmt.Errorf("failed to add message: %w", err)
	}
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetChatMessages(chat_id int) ([]models.Message, error) {
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetUser(user_id int) (models.User, error) {
	return db.GetUser(user_id)
}

func (a *App) Login(email, password string) (*models.User, error) {
	return auth.Login(email, password)
}

func (a *App) Register(username, email, password string) (*map[string]interface{}, error) {
	return auth.Register(username, email, password)
}
