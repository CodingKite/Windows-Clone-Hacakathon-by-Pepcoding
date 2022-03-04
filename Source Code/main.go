package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWin fyne.Window = myApp.NewWindow("my OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main(){



	// Weather ing
	weatherFile , err := os.Open("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\weatherIcon.png")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(weatherFile)
	b , err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	
// Calc Image
	calcFile , err := os.Open("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\calcIcon.png")
	if err != nil {
		log.Fatal(err)
	}

	v := bufio.NewReader(calcFile)
	c , err := ioutil.ReadAll(v)
	if err != nil {
		log.Fatal(err)
	}

// Editor image
	editFile , err := os.Open("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\editorIcon.png")
	if err != nil {
		log.Fatal(err)
	}

	e := bufio.NewReader(editFile)
	d , err := ioutil.ReadAll(e)
	if err != nil {
		log.Fatal(err)
	}
	


// // gallery image
	galFile , err := os.Open("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\galleryIcon.png")
	if err != nil {
		log.Fatal(err)
	}

	p := bufio.NewReader(galFile)
	g , err := ioutil.ReadAll(p)
	if err != nil {
		log.Fatal(err)
	}


	// bitcoin image
	bitFile , err := os.Open("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\bitcoinIcon.png")
	if err != nil {
		log.Fatal(err)
	}

	bi := bufio.NewReader(bitFile)
	cu , err := ioutil.ReadAll(bi)
	if err != nil {
		log.Fatal(err)
	}
	


	// myApp.Settings().SetTheme(theme.LightTheme())
	img  = canvas.NewImageFromFile("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\images\\windows.jpg")

	btn1 = widget.NewButtonWithIcon("",fyne.NewStaticResource("Weather",b),func(){
		showWeatherApp(myWin)
	})
	
	btn2 = widget.NewButtonWithIcon("",fyne.NewStaticResource("Calc",c),func(){
		showCalc()
	})
	
	btn3 = widget.NewButtonWithIcon("",fyne.NewStaticResource("gallery",g),func(){
		showGalleryApp(myWin)
	})
	
	btn4 = widget.NewButtonWithIcon("",fyne.NewStaticResource("Editor",d),func(){
		showTextEditor()
	})

	btn5 = widget.NewButtonWithIcon("",fyne.NewStaticResource("Bitcoin",cu),func(){
		showBitcoin(myWin)
	})

	btn6 = widget.NewButtonWithIcon("URL Shortner",theme.ComputerIcon(),func(){
		showUrlShortner(myWin)
	})

	DeskBtn =  widget.NewButtonWithIcon("Home",theme.HomeIcon(),func(){
		myWin.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

panelContent := container.NewCenter(container.NewGridWithColumns(9,layout.NewSpacer(),DeskBtn,btn1,btn2,btn3,btn4,btn5,btn6,layout.NewSpacer(),))

	// panelContent = container.NewVBox(taskbar)

	myWin.Resize(fyne.NewSize(924,620))
	myWin.CenterOnScreen()
	myWin.SetContent(container.NewBorder(nil,panelContent,nil,nil,img),)
	myWin.ShowAndRun()
}