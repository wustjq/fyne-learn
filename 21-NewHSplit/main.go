package main

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("my name is title")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//画布中创建文本
	labels1:=canvas.NewText("text1",color.Black)
	labels2:=canvas.NewText("text2",color.Black)

	//创建一个标签组件
	w1:=widget.NewLabel("widget label")
	//添加到窗口
	w.SetContent(container.NewVSplit(container.NewHSplit(labels1,labels2),w1))

	//运行
	w.ShowAndRun()
}

