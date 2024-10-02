package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	peer "github.com/piheta/sept/backend"
	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// cli args
	bFlag := flag.Bool("b", false, "Only start the backend")
	flag.Parse()
	if *bFlag {
		fmt.Println("The -b flag was passed")
		peer.Peer()
	}

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

	app := NewApp(
		user_repo,
		chat_repo,
		userchat_repo,
		message_repo,

		auth_service,
	)
	err := wails.Run(&options.App{
		Width:  700,
		Height: 512,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
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
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
