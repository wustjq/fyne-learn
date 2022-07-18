package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(500,500))

	list:=widget.NewList(
		func() int {return 5},
		func() fyne.CanvasObject {return widget.NewLabel("")},
		func(id widget.ListItemID, object fyne.CanvasObject) {object.(*widget.Label).SetText("Hello World")},
	)

	w.SetContent(list)
	w.ShowAndRun()	//运行
}


