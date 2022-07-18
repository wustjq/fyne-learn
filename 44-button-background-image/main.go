package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	w.SetContent(container.NewVBox(red_button(),red_button(),img_buttoon()))
	//运行
	w.ShowAndRun()
}

//给按钮上色
func red_button()(*fyne.Container){
	btn:=widget.NewButton("Visit", func() {

	})

	btn_color:=canvas.NewRectangle(color.NRGBA{R: 255,G: 0,B: 0,A: 255})

	red_contain:=container.New(layout.NewMaxLayout(),btn_color,btn)
	return red_contain
}

//给按钮上图片
func img_buttoon()(*fyne.Container){
	btn:=widget.NewButton("Visit", func() {

	})

	img:=canvas.NewImageFromFile("C:/Users/jq/Desktop/bdyjy.jpg")

	img_contain:=container.New(layout.NewMaxLayout(),img,btn)
	return img_contain
}


