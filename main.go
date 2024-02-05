package main

import (
	"context"
	"embed"
	"quartz/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

var log = logger.NewDefaultLogger()

func main() {
	// Create an instance of the app structure
	app.WailsApp = app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "quartz",
		Width:     600,
		Height:    450,
		Frameless: true,
		MinWidth:  400,
		MinHeight: 300,
		// MaxWidth:          1024,
		// MaxHeight:         768,
		HideWindowOnClose: true,
		AssetServer:       &assetserver.Options{Assets: assets},
		// BackgroundColour:  &options.RGBA{R: 25, G: 25, B: 25, A: 25},
		OnStartup: app.WailsApp.Startup,
		// AlwaysOnTop:    AlwaysOnTop,
		OnDomReady: startLogReader,
		Bind: []interface{}{
			app.WailsApp,
		},
		Logger:             log,
		LogLevel:           1,
		LogLevelProduction: 3,
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
			IsZoomControlEnabled: false,
			BackdropType:         windows.Acrylic,
		},
		Mac: &mac.Options{},
		OnBeforeClose: func(ctx context.Context) bool {
			log.Info("Closing...")
			WriteConfig()
			return false
		},
		Debug: options.Debug{OpenInspectorOnStartup: true},

		EnableDefaultContextMenu:         false,
		EnableFraudulentWebsiteDetection: false,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
