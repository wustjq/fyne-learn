package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

//func init() {
//	fontPaths := findfont.List()
//	for _, path := range fontPaths {
//		//楷体 "simkai.ttf"
//		//黑体 "simhei.ttf"
//		if strings.Contains(path, "simkai.ttf") {
//			os.Setenv("FYNE_FONT",path)
//			break
//		}
//	}
//}

func init() {
	os.Setenv("FYNE_FONT","msyhl.ttc")
}

func main(){
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("我是窗口")

	w.Resize(fyne.NewSize(500,500))

	label:=widget.NewLabel("这是标签")
	w.SetContent(label)
	//运行
	w.ShowAndRun()
}



