package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("my name is title")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	rand.Seed(time.Now().UnixNano())
	//新建一个标签
	label:=widget.NewLabel("my name is label")

	//创建一个按钮button
	button:=widget.NewButton("Generate", func() {
		rand1:=rand.Intn(100)
		label.Text=fmt.Sprint(rand1)
		label.Refresh()
	})

	//水平box将两个微件装起来
	newHBox:=container.NewVBox(button,label)

	//显示在窗口
	w.SetContent(newHBox)
	//运行
	w.ShowAndRun()
}
