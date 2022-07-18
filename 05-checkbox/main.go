package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my name is winodows title")

	//设置初始窗口打下
	w.Resize(fyne.NewSize(300,300))

	//定义一个装置  checkbox
	check:=widget.NewCheck("men", func(b bool) {
		//勾选上  就为true
		if b{
			fmt.Println("我是男性")
		}else{
			fmt.Println("我是女性")
		}
	})

	//将装置添加到窗口
	w.SetContent(check)

	//运行
	w.ShowAndRun()
}