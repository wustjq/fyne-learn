package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个window
	w := a.NewWindow("my title")

	//改变初始窗口大小
	w.Resize(fyne.NewSize(300,300))

	//创建一个按钮button
	button:=widget.NewButton("my name is button", func() {
		fmt.Println("button is pressed")
	})

	//将装置添加到窗口
	w.SetContent(button)

	//运行
	w.ShowAndRun()
}
