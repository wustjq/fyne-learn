package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))
	
	btn:=widget.NewButton("button", func() {
		fmt.Println("button is pressed")
	})

	//btn.Disable()
	
	w.SetContent(container.NewVBox(btn))
	//运行
	w.ShowAndRun()
}

