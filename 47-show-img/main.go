package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	w.Resize(fyne.NewSize(500,500))

	btn:=widget.NewButton("Open .jpg & .png", func() {
		fileDialog:=dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			data,_:=ioutil.ReadAll(r)
			res:=fyne.NewStaticResource(r.URI().Name(),data)

			img:=canvas.NewImageFromResource(res)
			win:=fyne.CurrentApp().NewWindow(r.URI().Name())

			win.SetContent(img)
			win.Resize(fyne.NewSize(500,500))
			win.Show()
		},w)

		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".jpg",".png"}))
		fileDialog.Show()
	})

	w.SetContent(btn)

	//运行
	w.ShowAndRun()
}

