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
	newItem:=fyne.NewMenuItem("New",nil)
	editItem:=fyne.NewMenuItem("Edit",nil)
	saveItem:=fyne.NewMenuItem("Save",nil)

	//创建menu   这个相当于是折叠起来先显示的
	menu1:=fyne.NewMenu("File",newItem,editItem,saveItem)
	menu2:=fyne.NewMenu("Config",newItem,editItem,saveItem)
	menu3:=fyne.NewMenu("Content",newItem,editItem,saveItem)

	//创建mainmenu
	mainmenu:=fyne.NewMainMenu(menu1,menu2,menu3)

	//设置在窗口上
	w.SetMainMenu(mainmenu)
	//运行
	w.ShowAndRun()
}

