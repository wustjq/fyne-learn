package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建一个app
	a:=app.New()

	//设置主题
	//a.Settings().SetTheme(theme.LightTheme())
	//a.Settings().SetTheme(theme.DarkTheme())

	//创建一个windows
	w:=a.NewWindow("my name is title")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//文字说明
	label:=widget.NewLabel("Fyne Theme")

	//亮色主题
	btn1:=widget.NewButton("Light Theme", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	//暗色主题
	btn2:=widget.NewButton("Dark Theme", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	//退出
	btn3:=widget.NewButton("Exit", func() {
		a.Quit()
	})

	//添加到窗口
	w.SetContent(container.NewVBox(label,btn1,btn2,btn3))
	//运行
	w.ShowAndRun()
}

