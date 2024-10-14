package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/piheta/sept/backend/controllers"
	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	execDir := filepath.Dir(execPath)

	// Create the folder inside the Resources directory of the app bundle
	resourcesPath := filepath.Join(execDir, "..", "Resources", "Data")
	err = os.MkdirAll(resourcesPath, 0o755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	db.SEPT_DATA = resourcesPath
	fmt.Printf(db.SEPT_DATA)
	//
	// startup logic
	// try to init db with jwt
	// if this fails, init repos with null
	// on login, repos will be correctly initialized from inside of "app" controller
	//
	auth_service := services.NewAuthSerivce()
	auth_service.LogInWithExistingJwt()

	user_repo := repos.NewUserRepo(db.DB)
	chat_repo := repos.NewChatRepo(db.DB)
	userchat_repo := repos.NewUserchatRepo(db.DB)
	message_repo := repos.NewMessageRepo(db.DB)

	app := controllers.NewApp(
		user_repo,
		chat_repo,
		userchat_repo,
		message_repo,

		auth_service,
	)
	err = wails.Run(&options.App{
		Width:     700,
		Height:    512,
		MinWidth:  400,
		MinHeight: 250,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Sept",
				Message: "© 2024 Piheta",
			},
			Preferences: &mac.Preferences{
				TabFocusesLinks:        mac.Enabled,
				TextInteractionEnabled: mac.Enabled,
				FullscreenEnabled:      mac.Disabled,
			},
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
