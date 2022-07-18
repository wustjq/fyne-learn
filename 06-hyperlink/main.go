package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
)
func main() {
	//创建一个app
	a:=app.New()

	//创建一个窗口
	w:=a.NewWindow("my name is winodows title")

	//设置默认窗口打下
	w.Resize(fyne.NewSize(300,300))

	////定义一个装置  hyperlink
	url,_:=url.Parse("https://www.cnblogs.com/wustjq/p/16426471.html")
	widgetUrl:=widget.NewHyperlink("my name is hyperlink",url)

	//添加至窗口
	w.SetContent(widgetUrl)

	//运行
	w.ShowAndRun()
}
