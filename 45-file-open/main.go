package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	btn:=widget.NewButton("Open .txt files", func() {
		//使用对话框打开文件
		//第一个参数是函数体，第二个参数是父窗口
		file_dial:=dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			//读取文件
			data,_:=ioutil.ReadAll(r)

			//得到结果
			result:=fyne.NewStaticResource("name",data)

			//展示文本在标签
			entry:=widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))

			//显示  title就设置为文件名，
			w:=fyne.CurrentApp().NewWindow(string(result.StaticName))

			w.SetContent(container.NewVBox(entry))

			w.Show()
		},w)

		//进行过滤
		file_dial.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		file_dial.Show()
	})

	w.SetContent(container.NewVBox(btn))
	//运行
	w.ShowAndRun()
}

