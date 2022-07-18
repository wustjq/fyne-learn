package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("my name is title")

	////指定你需要加载图片路径
	//r,_:=fyne.LoadResourceFromPath("./bdyjy.jpg")
	////设置app 图标
	//w.SetIcon(r)
	//指定你需要加载图片路径
	r,_:=fyne.LoadResourceFromURLString("https://picsum.photos/200")
	//设置app 图标
	w.SetIcon(r)
	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//运行
	w.ShowAndRun()
}

