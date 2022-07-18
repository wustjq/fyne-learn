package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(500,500))

	btn:=widget.NewButton("Button", func() {})
	btn.Resize(fyne.NewSize(40,100))
	btn.Move(fyne.Position{0,0})
	fmt.Println("button:",btn.Position())

	red_Rect:=canvas.NewRectangle(color.NRGBA{R: 255,G: 0,B: 0,A: 255})
	red_Rect.Resize(fyne.NewSize(40,100))
	fmt.Println("red_Rect:",red_Rect.Position())


	w.SetContent(container.NewWithoutLayout(btn,red_Rect))
	w.ShowAndRun()	//运行
}

