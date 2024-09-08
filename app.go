package main

import (
	"context"
	"fmt"

	"github.com/piheta/sept/infra/db"
	"github.com/piheta/sept/models"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db.InitDb()
}

var user_group = []models.User{}

func (a *App) GetUsers() []models.User {
	return db.GetAllUsers()
}

func (a *App) GetRooms() []models.User {
	return user_group
}

func (a *App) SendMessage(message string, chat_id int) []models.Message {
	fmt.Println(message)
	db.AddMessage(chat_id, 1, message)
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetChatMessages(chat_id int) []models.Message {
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetUser(user_id int) models.User {
	return db.GetUser(user_id)
}
