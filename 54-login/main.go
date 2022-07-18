package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Menu")

	//修改默认大小
	w.Resize(fyne.NewSize(500,500))

	label:=widget.NewLabel("")
	form:=widget.NewForm(
		widget.NewFormItem("UserName",widget.NewEntry()),
		widget.NewFormItem("PassWord",widget.NewPasswordEntry()),
	)
	form.OnCancel= func() {
		label.Text="Cancel"
		label.Refresh()
	}
	form.OnSubmit= func() {
		label.Text="Submit"
		label.Refresh()
	}

	w.SetContent(container.NewVBox(form,label))
	w.ShowAndRun()	//运行
}


