package backend

import (
	"context"
	"downloader/backend/config"
	"downloader/backend/replay"
	"downloader/backend/tool"
	"log"
	"os"
	"path"
	"time"

	// log "github.com/sirupsen/logrus"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx           context.Context
	cfg           config.CFG
	logf          *os.File
	debug         bool
	AppDeveloper  string
	AppName       string
	Version       string
	cfgSaved      bool
	backendReady  bool
	frontendReady bool
}

// NewApp 创建一个新的 App 应用程序
func NewApp(appDeveloper string, appName string, version string, debug bool) *App {
	var a = App{AppDeveloper: appDeveloper, AppName: appName, Version: version, debug: debug}

	return &a
}

// FrontendInited 前端初始化完成
func (a *App) FrontendInited() {
	a.frontendReady = true
	log.Println(a.frontendReady)
}

// Startup 在应用程序启动时调用
func (a *App) Startup(ctx context.Context) {
	// 执行初始化设置
	a.ctx = ctx

	// 释放附件
	go func() {
		if err := a.releaseAssets(); err != nil {
			log.Println(err)
		}
		log.Println("附件释放完成")
	}()

	// 读取设置
	var err error
	if a.cfg, err = config.ReadConfig(a.GetConfigPath("config.json")); err != nil {
		log.Println(err)
		a.cfg = config.DefaultCFG()
		log.Println("使用默认设置")
	}

	// 处理相关设置
	a.cfg = config.CheckConfig(a.cfg)
	a.cfg.Version = a.Version

	// 设置默认demo存放路径
	if a.cfg.DemoDir == "" {
		a.cfg.DemoDir = a.GetConfigPath("replay")
		if err := os.MkdirAll(a.cfg.DemoDir, os.ModePerm); err != nil {
			log.Println(err)
		}
	}

	a.backendReady = true
	log.Println("后端准备完成")
}

// DomReady 在前端Dom加载完毕后调用
func (a *App) DomReady(ctx context.Context) {
	log.Println("前端Dom加载完毕")

	for !a.backendReady || !a.frontendReady {
		time.Sleep(time.Millisecond * 100)
		log.Println(a.frontendReady)
	}

	rt.EventsEmit(a.ctx, "inited")
	log.Println("前后端初始化均完成")
}

// Shutdown 在前端Dom销毁 应用程序终止时被调用
func (a *App) Shutdown(ctx context.Context) {
	// 在此处做一些资源释放的操作
	log.Println("后端结束之前")

	err := config.SaveConfig(a.cfg, a.GetConfigPath("config.json"))
	if err != nil {
		log.Println(err)
	}

}

// BeforeClose 关闭应用程序之前回调
func (a *App) BeforeClose(ctx context.Context) bool {
	// 同步前端的设置
	log.Println("关闭程序之前")

	a.cfgSaved = false
	rt.EventsEmit(ctx, "shutdown")

	time.Sleep(time.Millisecond * 500)

	log.Println("设置已保存")
	return false
}

func (a *App) GetConfigPath(cfgName string) (cfgPath string) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("获取应用配置目录失败: " + err.Error())
	}

	return path.Join(tool.FormatPath(configDir), a.AppDeveloper, a.AppName, cfgName)
}

// 前端获取设置
func (a *App) GetCFG() config.CFG {
	log.Println("getCFG", a.cfg)

	return a.cfg
}

// 后端读取前端发送的设置
func (a *App) SetCFG(cfg config.CFG) {
	a.cfg = cfg

	a.cfgSaved = true

	log.Println("setCFG", a.cfg, a.cfgSaved)
}

// 功能
func (a *App) ParseShareCode(shareLink string) {
	urls := replay.GetDemoLink(shareLink, a.GetConfigPath(""))

	for _, url := range urls {
		rt.BrowserOpenURL(a.ctx, url)
	}
}

// 通知消息
func (a *App) info(title string, content string) {
	rt.EventsEmit(a.ctx, "info", title, content)
}

func (a *App) success(title string, content string) {
	rt.EventsEmit(a.ctx, "success", title, content)
}

func (a *App) warning(title string, content string) {
	rt.EventsEmit(a.ctx, "warning", title, content)
}

func (a *App) error(title string, content string) {
	rt.EventsEmit(a.ctx, "error", title, content)
}
