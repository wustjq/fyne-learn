package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(500,500))
	label:=widget.NewLabel("")
	sel:=widget.NewSelect([]string{"one","two","three","four"}, func(s string) {
		fmt.Printf("selected is %s\n",s)
		label.Text=s
		label.Refresh()
	})

	w.SetContent(container.NewVBox(sel,label))
	//运行
	w.ShowAndRun()
}


