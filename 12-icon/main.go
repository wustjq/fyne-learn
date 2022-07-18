package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)
func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my name is winodows title")

	//设置默认窗口打下
	w.Resize(fyne.NewSize(300,300))

	//定义一个组件  cards
	widgetCard:=widget.NewCard("title","sub title",canvas.NewCircle(color.Black))

	//添加到窗口
	w.SetContent(widgetCard)

	//运行
	w.ShowAndRun()
}

