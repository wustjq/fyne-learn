package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
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

	//播种随机数
	rand.Seed(time.Now().UnixNano())
	//画图创建矩形
	rect:=canvas.NewRectangle(color.NRGBA{R: 0,G: 0,B: 0,A: 255})
	rect.SetMinSize(fyne.NewSize(150,150))

	//随机颜色
	btn1:=widget.NewButton("Ranndom Color", func() {
		rect.FillColor=color.NRGBA{R: uint8(rand.Intn(255)),B: uint8(rand.Intn(255)),G: uint8(rand.Intn(255)),A: 255}
		rect.Refresh()
	})

	//随机红色
	btnRed:=widget.NewButton("Ranndom Red", func() {
		rect.FillColor=color.NRGBA{R: uint8(rand.Intn(255)),B:0,G: 0,A: 255}
		rect.Refresh()
	})

	//随机绿
	btnGreen:=widget.NewButton("Ranndom Green", func() {
		rect.FillColor=color.NRGBA{R: 0,B:0,G: uint8(rand.Intn(255)),A: 255}
		rect.Refresh()
	})

	//随机蓝色
	btnBlue:=widget.NewButton("Ranndom Blue", func() {
		rect.FillColor=color.NRGBA{R: 0,B:uint8(rand.Intn(255)),G: 0,A: 255}
		rect.Refresh()
	})

	w.SetContent(container.NewVBox(rect,btn1,btnRed,btnBlue,btnGreen))
	//运行
	w.ShowAndRun()
}

