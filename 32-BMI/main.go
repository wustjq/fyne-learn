package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

func main() {
	//创建app
	a:=app.New()

	//创建一个windows
	w:=a.NewWindow("Rand Code")

	//修改默认大小
	w.Resize(fyne.NewSize(300,300))

	//显示文本
	title:=canvas.NewText("BMI Calculate",color.Black)

	//指定两个输入
	weightInput:=widget.NewEntry()
	weightInput.SetPlaceHolder("Pleace input your weight(kg)")

	heightInput:=widget.NewEntry()
	heightInput.SetPlaceHolder("Pleace input your height(m)")



	//输出文本
	textAns:=canvas.NewText("",color.Black)
	textAns.TextSize=20

	height,_:=strconv.Atoi(heightInput.Text)
	heightvar:=float64(height)

	weight,_:=strconv.Atoi(weightInput.Text)
	weighttvar:=float64(weight)

	//校核
	if weighttvar<=0 || weighttvar>=250 || heightvar<=0 || heightvar>=2.5{
		textAns.Text="Pleace Enter Again ....."
		textAns.Color=color.NRGBA{R: 255,B: 0,G: 0,A: 255}

		//窗口显示
		w.SetContent(container.NewVBox(title,weightInput,heightInput,textAns))
	}else{
		//计算按钮
		btn:=widget.NewButton("Calculate", func() {
			textAns.Text=CalculateBmi(heightvar,weighttvar)
			textAns.Refresh()
		})
		//窗口显示
		w.SetContent(container.NewVBox(title,weightInput,heightInput,textAns,btn))
	}
	//运行
	w.ShowAndRun()
}

//计算BMI
func CalculateBmi(height,weight float64)string{
	//这里体重是kg  身高是m
	bmi:=weight/(height*height)
	if bmi>=40{
		return "Extremely Obesity"
	}else if bmi>=30{
		return "Severe Obesity"
	}else if bmi>=28{
		return "Obesity"
	}else if bmi>=24{
		return "Overweight"
	}else if bmi>=18.5{
		return "Normal"
	}else{
		return "Thin"
	}
}