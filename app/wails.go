package app

import (
	"context"
	"quartz/api"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var Version = "0.0.1"

var WailsApp *App

// App struct
type App struct {
	Ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// NewApp creates a new App application struct
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
	runtime.EventsOn(ctx, "SendPlayer", func(optionalData ...interface{}) {
		name := optionalData[0].(string)
		runtime.LogDebug(ctx, "SendStats: "+name)
		Emit("addPlayer", api.GetBWStats(name, api.Cubelify))
	})
}

func (a *App) GetVersion() string {
	return Version
}

func Emit(s string, opts ...any) {
	runtime.EventsEmit(WailsApp.Ctx, s, opts...)
}
