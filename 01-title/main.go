package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my frist title")

	w.ShowAndRun()
}

