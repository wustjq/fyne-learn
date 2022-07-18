package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//创建多个item
	item1:=fyne.NewMenuItem("New",nil)
	item2:=fyne.NewMenuItem("Edit",nil)
	item3:=fyne.NewMenuItem("Save",nil)

	item2.ChildMenu=fyne.NewMenu("label",
		fyne.NewMenuItem("Add",nil),
		fyne.NewMenuItem("Sub",nil),
		fyne.NewMenuItem("Div",nil),
	)
	//创建menu   这个相当于是折叠起来先显示的
	menu:=fyne.NewMenu("File",item1,item2,item3)

	//创建mainmenu
	mainmenu:=fyne.NewMainMenu(menu)

	//设置在窗口上
	w.SetMainMenu(mainmenu)
	//运行
	w.ShowAndRun()
}

