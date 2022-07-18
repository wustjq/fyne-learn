package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Rand Code")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	title:=canvas.NewText("Password Generator",color.Black)

	//客户输入密码长度的位置
	input:=widget.NewEntry()
	input.SetPlaceHolder("Enter Password length")

	//生成结果的位置  初始结果为空
	textAns:=canvas.NewText("",color.Black)
	textAns.TextSize=20
	//运行的按钮
	btn:=widget.NewButton("Generate", func() {
		//获取输入的数字结果，转成int类型
		length,_:=strconv.Atoi(input.Text)
		//验证码
		ans:=RandGenCode(length)
		textAns.Text=ans
		textAns.Refresh()
	})

	w.SetContent(container.NewVBox(title,input,textAns,btn))
	//运行
	w.ShowAndRun()
}

func RandGenCode(length int)(ans string){
	rand.Seed(time.Now().UnixNano())
	//数字
	strNum:="0123456789"
	//小写
	strLow:="abcdefghijklmnopqrstuvwxyz"
	//大写
	strUp:="ABCDEFGHIJKLMNOPQRSTUVWSYZ"
	for i:=0;i<length;i++{
		randstr:=rand.Intn(3)
		switch randstr {
		case 0:
			rand1:=rand.Intn(len(strNum))
			ans+=string(strNum[rand1])

		case 1:
			rand2:=rand.Intn(len(strLow))
			ans+=string(strLow[rand2])

		case 2:
			rand3:=rand.Intn(len(strUp))
			ans+=string(strUp[rand3])

		}
	}
	return ans
}