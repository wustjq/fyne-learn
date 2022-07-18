package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my name is winodows title")

	//设置默认窗口打下
	w.Resize(fyne.NewSize(300,300))


	//画布里面创建图片
	img:=canvas.NewImageFromFile("C:/Users/jq/Desktop/1.jpg")
	//添加到窗口
	w.SetContent(img)

	//运行
	w.ShowAndRun()
}
