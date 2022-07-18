package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	entry:=widget.NewMultiLineEntry()
	//创建按钮用来储存文件
	btn:=widget.NewButton("Save .txt file", func() {
		//第一个是个函数  第二个是父窗口
		fileDialog:=dialog.NewFileSave(func(wc fyne.URIWriteCloser, err error) {
			//从输入获取数据信息
			data:=entry.Text
			//写到文件去
			wc.Write([]byte(data))
		},w)

		//设置文件名
		fileDialog.SetFileName("anyFileName.txt")
		fileDialog.Show()
		fileDialog.Refresh()
	})

	//窗口显示
	w.SetContent(container.NewVBox(entry,btn))
	//运行
	w.ShowAndRun()
}

