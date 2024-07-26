package main

import (
	"context"
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

type room_model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var rooms = []room_model{
	{1, "Study"},
	{2, "Game"},
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

func (a *App) GetRooms() []room_model {
	return rooms
}