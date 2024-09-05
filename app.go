package main

import (
	"context"
	"fmt"

	"github.com/piheta/sept/infra/db"
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

type user_model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

type message_model struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

var user_group = []user_model{
	{1, "User8", ""},
	{2, "User9", ""},
}

var servers = []server_model{
	{1, "Sept"},
	{2, "Microsoft"},
	{3, "News"},
	{4, "Google"},
	{5, "Contoso"},
	{6, "Midjourney"},
	{7, "Apple"},
}

var users = []user_model{
	{1, "User1", ""},
	{2, "User2", ""},
	{3, "User3", ""},
	{4, "User4", ""},
	{5, "User5", ""},
	{6, "User6", ""},
	{7, "User7", ""},
}

func (a *App) GetServers() []server_model {
	return servers
}

func (a *App) GetUsers() []user_model {
	return users
}

func (a *App) GetRooms() []user_model {
	return user_group
}

func (a *App) SendMessage(message string) string {
	fmt.Println(message)
	db.AddMessage(1, 1, message)
	return db.GetMessagesByChatID(1)
}

func (a *App) GetUserUserMessages(user_id int) string {
	return db.GetMessagesByChatID(1)
}
