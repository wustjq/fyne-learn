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
	//画布里面创建梯度
	canvasgradient:=canvas.NewHorizontalGradient(color.White,color.Black)
	//canvas.NewLinearGradient()
	//canvas.NewRadialGradient()
	//添加到窗口
	w.SetContent(canvasgradient)

	//运行
	w.ShowAndRun()
}

