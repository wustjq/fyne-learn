package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//multiline 多行
	multilineEntry:=widget.NewMultiLineEntry()
	const loren="The sunset and the solitary bird fly together, and the autumn water is the same color as the sky"
	multilineEntry.SetText(loren)
	multilineEntry.Wrapping=fyne.TextTruncate

	w.SetContent(multilineEntry)

	//运行

	w.ShowAndRun()
}

