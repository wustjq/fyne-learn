package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"net/http"
)

func main() {
	//创建一个app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("my name is title")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//创建组件label
	labelTitle:=widget.NewLabel("what is my ip?")
	ip:=widget.NewLabel("your ip is ...")
	value:=widget.NewLabel("...")
	button:=widget.NewButton("Run", func() {
		value.Text=myIp()
		value.Refresh()
	})

	//添加组件到窗口
	w.SetContent(container.NewVBox(labelTitle,ip,value,button))

	//运行
	w.ShowAndRun()
}

type IP struct {
	query string
}

func myIp()string{
	req,err:=http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println("http get() err:",err)
		return err.Error()
	}
	defer req.Body.Close()

	body,err:=ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll() err:",err)
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body,&ip)
	return ip.query
}