package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	f:=20.0
	data:=binding.BindFloat(&f)
	//滑动数据的组件
	slider1:=widget.NewSliderWithData(0.0,100.0,data)

	//根据数据创建label
	//label1:=widget.NewLabelWithData(
	//	binding.FloatToString(data),
	//)

	//w.SetContent(container.NewVBox(slider1,label1))
	w.SetContent(slider1)

	//运行
	w.ShowAndRun()
}

