package main

import (
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type app struct {
	runtime *wails.Runtime
	appMenu *menu.Menu
}

func NewApp() *app {
	app := &app{}
	return app
}

// startup is called at application startup
func (a *app) startup(runtime *wails.Runtime) {
	a.runtime = runtime
	runtime.Window.SetTitle("wails-demo")
}

// shutdown is called at application termination
func (a *app) shutdown() {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *app) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
