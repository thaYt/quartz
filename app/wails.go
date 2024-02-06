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

var (
	Events = make(map[string]func(...any))
)

func sendPlayer(data ...any) {
	name := data[0].(string)
	runtime.LogDebug(WailsApp.Ctx, "SendStats: "+name)
	Emit("addPlayer", api.GetBWStats(name, api.Cubelify))
}

// NewApp creates a new App application struct
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
	Events["SendPlayer"] = sendPlayer
	for v, f := range Events {
		On(v, f)
	}
}

func (a *App) GetVersion() string {
	return Version
}

func Emit(s string, opts ...any) {
	runtime.EventsEmit(WailsApp.Ctx, s, opts...)
}

func On(s string, f func(...any)) {
	runtime.EventsOn(WailsApp.Ctx, s, f)
}
