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

func (a *App) GetUsers() []string {
	return []string{"10.0.0.10:5173", "10.0.0.12:8080", "192.168.1.23:80", "10.223.0.2:443"}
}

func (a *App) GetRooms() []string {
	return []string{"Study", "Game"}
}

func (a *App) GetServers() []string {
	return []string{"Sept", "Microsoft", "News", "Google", "Contoso", "Midjourney", "Apple"}
}
