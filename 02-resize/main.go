package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个window
	w:=a.NewWindow("my title")

	//设置初始窗口大小
	w.Resize(fyne.NewSize(600,600))

	//运行
	w.ShowAndRun()
}
