package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个window
	w := a.NewWindow("my title")

	//修改初始显示窗口大小
	w.Resize(fyne.NewSize(300,300))

	w.SetContent(widget.NewLabel("title label"))

	//运行
	w.ShowAndRun()
}
