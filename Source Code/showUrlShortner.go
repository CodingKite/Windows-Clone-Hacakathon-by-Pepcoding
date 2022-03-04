package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showUrlShortner(w fyne.Window){
	// go mod init >> creates go.mod file 
	// go mod tidy >> list all the dpeencies inside imort

	// creates an new app
	// a := app.New()

	// Window in out aapp
	// w := a.NewWindow("Title of the Winode")

	// resize out Window
	// w.Resize(fyne.NewSize(400,400))


	// UI part
	title1 := widget.NewLabel("URL Shortner App")
	title1.Alignment = fyne.TextAlignCenter
	title1.TextStyle = fyne.TextStyle{Bold:true}

	long_url := widget.NewEntry()
	long_url.PlaceHolder="Enter Long URL Here..."

	short_url := widget.NewMultiLineEntry()
	short_url.PlaceHolder="Short URL will be displayerd here.."
	
	btn := widget.NewButton("Shorten",func(){
		// fmt.Print("Button is clicked")

		// Made URL dynamix
		URL := fmt.Sprintf("https://api.1pt.co/addURL?long=%s",long_url.Text)

		// Http response

		res, _ := http.Get(URL)

		defer res.Body.Close()

		// Get the body of the response
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(body)
		fmt.Println("=====================")

		uRLShortner, _ := UnmarshalURLShortner(body)
		fmt.Println(uRLShortner)
		fmt.Println("=====================")
		fmt.Println(uRLShortner.Status)
		fmt.Println(uRLShortner.Message)
		fmt.Println(uRLShortner.Long)
		fmt.Println(uRLShortner.Short)


		t := "1pt.co/"+uRLShortner.Short
		fmt.Println(t)

		short_url.Text ="Here is your Short Url: \n" + t
		short_url.Refresh()


	})


	c := container.NewVBox(title1, long_url,btn,short_url)
myWin.SetContent(container.NewBorder(nil,panelContent,nil,nil,c),)

// Shoe the winode and run the app
	w.Show()
}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    uRLShortner, err := UnmarshalURLShortner(bytes)
//    bytes, err = uRLShortner.Marshal()

func UnmarshalURLShortner(data []byte) (URLShortner, error) {
	var r URLShortner
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *URLShortner) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type URLShortner struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Short   string `json:"short"`
	Long    string `json:"long"`
}

