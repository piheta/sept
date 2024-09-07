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

type server_model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type message_model struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

var user_group = []models.User_model{}

var servers = []server_model{
	{1, "Sept"},
	{2, "Microsoft"},
	{3, "News"},
	{4, "Google"},
	{5, "Contoso"},
	{6, "Midjourney"},
	{7, "Apple"},
}

func (a *App) GetServers() []server_model {
	return servers
}

func (a *App) GetUsers() []models.User_model {
	return db.GetAllUsers()
}

func (a *App) GetRooms() []models.User_model {
	return user_group
}

func (a *App) SendMessage(message string, chat_id int) string {
	fmt.Println(message)
	db.AddMessage(chat_id, 1, message)
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetUserUserMessages(chat_id int) string {
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetUser(user_id int) models.User_model {
	return db.GetUser(user_id)
}
