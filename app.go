package main

import (
	"changeme/action"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	action *action.Ollama操作
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.action = action.NeOllama操作()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	println("收到js的调用信息")
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) E按钮1被点击() string {
	println("E按钮1被点击")
	return fmt.Sprintf("E按钮1被点击")
}

func (a *App) E获取模型列表() string {
	println("E获取模型列表")
	return a.action.E获取模型列表()
}

func (a *App) E下载模型(name string) string {
	println("E下载模型")
	data := a.action.E下载模型(name, func(总大小, 已下载, 总进度 string) {
		println(总大小, 已下载, 总进度)
		jsonstring := fmt.Sprintf(`{"total_size":"%s","downloaded":"%s","progress":"%s"}`, 总大小, 已下载, 总进度)
		runtime.EventsEmit(a.ctx, "downInfo", jsonstring)
	})
	println(data)
	return data
}

func (a *App) E停止下载() string {
	a.action.E停止下载()
	return "ok"
}
func (a *App) E删除模型(name string) string {

	data := a.action.E删除模型(name)
	return data
}
func (a *App) E模型改名(name string, newName string) string {

	data := a.action.E复制模型(name, newName)
	return data
}

func (a *App) E对话(name string, 内容 string) string {

	data := a.action.E对话(name, 内容)
	return data
}
