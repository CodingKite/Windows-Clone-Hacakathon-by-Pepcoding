package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var num_article int = 1

// var news News

func showBitcoin(w fyne.Window) {
	// a := app.NewWithID("")
	// w := a.NewWindow("Currency Pice")
	// w.Resize(fyne.NewSize(300, 400))

	URL := fmt.Sprintf("https://api.coindesk.com/v1/bpi/currentprice.json")

	//API
	res, _ := http.Get(URL)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	coinDisk, _ := UnmarshalCoinDisk(body)

	fmt.Println(coinDisk)
	label1 := canvas.NewText(coinDisk.ChartName+" Chart", color.White)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextStyle = fyne.TextStyle{Bold: true}

	s2 := fmt.Sprintf("Code: %s", coinDisk.Bpi.Usd.Code)
	s4 := fmt.Sprintf("Rate: %s", coinDisk.Bpi.Usd.Rate)
	s5 := fmt.Sprintf("Desc: %s", coinDisk.Bpi.Usd.Description)
	s6 := fmt.Sprintf("RateFloat: %.2f", coinDisk.Bpi.Usd.RateFloat)

	grid := widget.NewTextGridFromString(s2 + "\n" + s4 + "\n" + s5 + "\n" + s6 + "\n")
	grid.SetStyleRange(4, 0, 7, 7, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})
	grid.Rows[1].Style = &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 64, A: 128}}
	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	r2 := fmt.Sprintf("Code: %s", coinDisk.Bpi.Eur.Code)

	r4 := fmt.Sprintf("Rate: %s", coinDisk.Bpi.Eur.Rate)
	r5 := fmt.Sprintf("Desc: %s", coinDisk.Bpi.Eur.Description)
	r6 := fmt.Sprintf("RateFloat: %.2f", coinDisk.Bpi.Eur.RateFloat)

	grid2 := widget.NewTextGridFromString(r2 + "\n" + r4 + "\n" + r5 + "\n" + r6 + "\n")
	grid2.SetStyleRange(4, 0, 7, 7, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 64, B: 192, A: 128}})
	grid2.Rows[1].Style = &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}}
	grid2.ShowLineNumbers = true
	grid2.ShowWhitespace = true

	v1 := fmt.Sprintf("Code: %s", coinDisk.Bpi.Gbp.Code)
	v4 := fmt.Sprintf("Rate: %s", coinDisk.Bpi.Gbp.Rate)
	v5 := fmt.Sprintf("Desc: %s", coinDisk.Bpi.Gbp.Description)
	v6 := fmt.Sprintf("RateFloat: %.2f", coinDisk.Bpi.Gbp.RateFloat)

	grid3 := widget.NewTextGridFromString(v1 + "\n" + v4 + "\n" + v5 + "\n" + v6 + "\n")
	grid3.SetStyleRange(4, 0, 7, 7, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 0, G: 64, B: 192, A: 128}})
	grid3.Rows[1].Style = &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 192, B: 64, A: 128}}
	grid3.ShowLineNumbers = true
	grid3.ShowWhitespace = true

	img := canvas.NewImageFromFile("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\newfile\\images\\price.png")
	// img.Resize(fyne.NewSize(150, 150))
	img.FillMode = canvas.ImageFillOriginal
	e := container.NewVBox(label1, grid, grid2, grid3)
	e.Resize(fyne.NewSize(300, 300))
	c := container.NewBorder(img, nil, nil, nil, e)

	

	myWin.SetContent(container.NewBorder(nil,panelContent,nil,nil,c))

	w.Show()

}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    coinDisk, err := UnmarshalCoinDisk(bytes)
//    bytes, err = coinDisk.Marshal()

func UnmarshalCoinDisk(data []byte) (CoinDisk, error) {
	var r CoinDisk
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CoinDisk) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CoinDisk struct {
	Time       Time   `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        Bpi    `json:"bpi"`
}

type Bpi struct {
	Usd Eur `json:"USD"`
	Gbp Eur `json:"GBP"`
	Eur Eur `json:"EUR"`
}

type Eur struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

type Time struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	Updateduk  string `json:"updateduk"`
}

