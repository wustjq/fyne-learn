package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)
func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my name is winodows title")

	//设置默认窗口打下
	w.Resize(fyne.NewSize(300,300))

	//画布里面创建矩形
	canvasRect:=canvas.NewRectangle(color.NRGBA{B: 255,G: 0,R: 0,A: 255})
	//边框颜色 边框线宽
	canvasRect.StrokeColor=color.Black
	canvasRect.StrokeWidth=3

	//添加到窗口
	w.SetContent(canvasRect)

	//运行
	w.ShowAndRun()
}
