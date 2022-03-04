package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
// func showCalc(myWin fyne.Window) {
	// Creating an App and adding Window to it
	// myApp := app.New()
	// myWin := myApp.NewWindow("Calculator")
	
	// Resizing the Window Size
	// myWin.Resize(fyne.NewSize(600,400))
	
	// Setting Icon to the App
	// Loading icon from an Path
	icon,_ := fyne.LoadResourceFromPath("C:/Users/rahul/Downloads/Video/Pepcoding fyne vids/Project1 Calculator/calcIcon.png")
	myWin.SetIcon(icon)
	
	// Input & Output
	output:= ""
	input := widget.NewLabel(output)
	

	var historyArray []string 
	historyString :=""
	historyShow := widget.NewLabel(historyString)
	isHistory := false

	// Buttons
		// Misc Button
		historyBtn := widget.NewButton("History",func() {
			if isHistory{
				historyString= ""
			}else{
				for i:= len(historyArray)-1;i>0;i-- {
				historyString += historyArray[i]
				historyString += "\n"
				}
			}
			isHistory =! isHistory
			historyShow.SetText(historyString)
		})


		backBtn := widget.NewButton("Back",func() {
			
			if len(output)>0 {
			output = output[:len(output)-1]
			input.SetText(output)
			}

		})
		clearBtn := widget.NewButton("Clear",func() {
			output = ""
			input.SetText(output)
		})
		equlBtn := widget.NewButton("=",func() {
			
			expression, err := govaluate.NewEvaluableExpression(output);
			if err == nil {
				result, err := expression.Evaluate(nil);
				if err == nil {
					answer := strconv.FormatFloat(result.(float64),'f',-1,64)
					strToAppend := output+ " = "+answer
					historyArray = append(historyArray, strToAppend)
					output = answer
				}else{
					output= "ERROR"
				}
			}else{
					output= "ERROR"
			}
			input.SetText(output)

		})

		// Operator Button
		divideBtn := widget.NewButton("/",func() {
			output += "/"
			input.SetText(output)
		})
		multiplyBtn := widget.NewButton("x",func() {
			output += "*"
			input.SetText(output)
		})
		plusBtn := widget.NewButton("+",func() {
			output += "+"
			input.SetText(output)
		})
		minusBtn := widget.NewButton("--",func() {
			output += "-"
			input.SetText(output)
		})

		// Number Buttons
		
		openBracketBtn := widget.NewButton("(",func() {
			output += "("
			input.SetText(output)
		})

		closeBracketBtn := widget.NewButton(")",func() {
			output += ")"
			input.SetText(output)
		})
		pointBtn := widget.NewButton(".",func(){
			output += "."
			input.SetText(output)
		})

		nineBtn := widget.NewButton("9",func() {
			output += "9"
			input.SetText(output)
		})
		eightBtn := widget.NewButton("8",func() {
			output += "8"
			input.SetText(output)
		})
		sevenBtn := widget.NewButton("7",func() {
			output += "7"
			input.SetText(output)
		})
		sixBtn := widget.NewButton("6",func() {
			output += "6"
			input.SetText(output)
		})
		fiveBtn := widget.NewButton("5",func() {
			output += "5"
			input.SetText(output)
		})
		fourBtn := widget.NewButton("4",func() {
			output += "4"
			input.SetText(output)
		})
		threeBtn := widget.NewButton("3",func() {
			output += "3"
			input.SetText(output)
		})
		twoBtn := widget.NewButton("2",func() {
			output += "2"
			input.SetText(output)
		})
		oneBtn := widget.NewButton("1",func() {
			output += "1"
			input.SetText(output)
		})
		zeroBtn := widget.NewButton("0",func() {
			output += "0"
			input.SetText(output)
		})


	// Grouping Similar Buttons together

	calcContainer := container.NewVBox(
		input,
		historyShow,
		container.NewGridWithColumns(2,historyBtn,backBtn),
		container.NewGridWithColumns(4,clearBtn,openBracketBtn,closeBracketBtn,divideBtn),
		container.NewGridWithColumns(4,sevenBtn,eightBtn,nineBtn,multiplyBtn),
		container.NewGridWithColumns(4,fourBtn,fiveBtn,sixBtn,minusBtn),
		container.NewGridWithColumns(4,oneBtn,twoBtn,threeBtn,plusBtn),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,zeroBtn,pointBtn),equlBtn),
	)

	myWin := myApp.NewWindow("Calc")
	myWin.Resize(fyne.NewSize(500,200))

	myWin.SetContent(
		container.NewBorder(nil,DeskBtn,nil,nil,calcContainer),)
		// myWin.ShowAndRun()
	myWin.Show()
}