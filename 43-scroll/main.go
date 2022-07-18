package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//red rect
	redRect:=canvas.NewRectangle(color.NRGBA{R: 255,G: 0,B: 0,A: 255})
	redRect.SetMinSize(fyne.NewSize(200,200))

	//blue rect
	blueRect:=canvas.NewRectangle(color.NRGBA{R: 0,G: 0,B: 255,A: 255})
	blueRect.SetMinSize(fyne.NewSize(200,200))

	c:=container.NewHBox(redRect,blueRect)

	//创建scroll
	scroll:=container.NewScroll(c)

	scroll.Direction=container.ScrollHorizontalOnly
	w.SetContent(scroll)
	//运行
	w.ShowAndRun()
}


