package main

import (
	"downloader/backend"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := backend.NewApp(
		"One-Studio",           // Developer 开发者
		"csgo-demo-downloader", // AppName   应用名称
		"0.0.1",                // Version   版本
		false,                  // Debug     是否开启debug日志输出
	)

	err := wails.Run(&options.App{
		Title:             "CSGO Demo 下载器",
		Width:             720,
		Height:            480,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.Startup,
		OnDomReady:        app.DomReady,
		OnBeforeClose:     app.BeforeClose,
		OnShutdown:        app.Shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 true,
			DisableFramelessWindowDecorations: false,
		},
		Mac: &mac.Options{
			// WebviewIsTransparent: true,
			// WindowIsTranslucent:  true,
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			About: &mac.AboutInfo{
				Title:   "CSGO工具箱",
				Message: "a tool for parsing demo sharecode and downloading gotv demos",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
