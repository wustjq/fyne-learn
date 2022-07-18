package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

	//创建图片组件  初始显示1.png
	canvasImg:=canvas.NewImageFromFile("C:/Users/jq/Desktop/dice/1.png")
	//以原始图片样式显示    必须有
	canvasImg.FillMode=canvas.ImageFillOriginal

	//随机数种子
	rand.Seed(time.Now().UnixNano())

	//创建一个组件button
	button:=widget.NewButton("run", func() {
		//1-6
		rand1:=rand.Intn(6)+1
		//修改图片路径
		canvasImg.File=fmt.Sprintf("C:/Users/jq/Desktop/dice/%d.png",rand1)
		//刷新
		canvasImg.Refresh()
	})

	//创建竖直box容器
	vBox:=container.NewVBox(canvasImg,button)

	//在窗口显示
	w.SetContent(vBox)
	//运行
	w.ShowAndRun()
}