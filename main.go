package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed .version
var version string

func main() {
	println("wdemo", version)
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	println("exited")
}

func run() error {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:             "wdemo",
		Width:             800,
		Height:            600,
		MinWidth:          800,
		MinHeight:         600,
		HideWindowOnClose: true,
		Mac: &mac.Options{
			WebviewIsTransparent:          true,
			WindowBackgroundIsTranslucent: true,
			TitleBar:                      mac.TitleBarHiddenInset(),
			Menu:                          app.appMenu,
			ActivationPolicy:              mac.NSApplicationActivationPolicyAccessory,
		},
		LogLevel: logger.DEBUG,
		Startup:  app.startup,
		Shutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
