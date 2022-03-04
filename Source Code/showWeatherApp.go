package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showWeatherApp(myWin fyne.Window){
// Creating App
	// myApp := app.New()

// Creating NewWindow
	// myWin := myApp.NewWindow("Weather App")

// Resizing the NewWindow
	// myWin.Resize(fyne.NewSize(500,500))

// UI Part

	// Image
		
	imgWeather := canvas.NewImageFromFile("C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\newfile\\images\\weather.png")
		imgWeather.FillMode = canvas.ImageFillOriginal

	// Head Label
		headingLabel := canvas.NewText("Weather Details",color.White)
		headingLabel.TextStyle = fyne.TextStyle{Bold:true}
		headingLabel.Alignment = fyne.TextAlignCenter

	// Weather Label

		countryLabel := canvas.NewText("",color.White)
		windLabel := canvas.NewText("",color.White)
		tempLabel := canvas.NewText("",color.White)
		humidityLabel:= canvas.NewText("",color.White)

	// Alignng Labels and Setting them to Bold
		countryLabel.Alignment = fyne.TextAlignCenter
		countryLabel.TextStyle = fyne.TextStyle{Bold:true}
		windLabel.Alignment = fyne.TextAlignCenter
		windLabel.TextStyle = fyne.TextStyle{Bold:true}
		tempLabel.Alignment = fyne.TextAlignCenter
		tempLabel.TextStyle = fyne.TextStyle{Bold:true}
		humidityLabel.Alignment = fyne.TextAlignCenter
		humidityLabel.TextStyle = fyne.TextStyle{Bold:true}

	// Dropdown
	
		var setLocation = []string{"Mumbai","Kolkata","Hyderabad","Delhi","Jaipur",""}
		selectDropdown := widget.NewSelect(setLocation, func (location string, )  {
		// fmt.Printf("my city %s is awesome \n",location)

	// API Part
			// Changing the String URL to be passed
			apiCall :="https://api.openweathermap.org/data/2.5/weather?q="+location+"&appid=d8ebaf6142f3daf9ecb1eded890116c6"
			// fmt.Println("API CALL MODIFIED: "+apiCall)

			// Sending Request API call and Getting response along with whole body
			res, err :=  http.Get(apiCall)
				if err!= nil{
					fmt.Println(err)
				}

			// Cloasing th Body
			defer res.Body.Close()

			// Reading all the Data in the Body of the Response
			body, err := ioutil.ReadAll(res.Body)
			if err!= nil{
				fmt.Println(err)
			}

			// Converting body to type string
			// sb := string(body)
			// fmt.Println(sb)

			// Calling the UnmarshalWeather function and passing the body into it
			weather, err := UnmarshalWeather(body)
			if err!= nil{
				fmt.Println(err)
			}
			
			// Setting the different Values into Labels
			countryLabel.Text = fmt.Sprintf("Country %s",weather.Sys.Country)
			windLabel.Text = fmt.Sprintf("Wind Speed %.2f",weather.Wind.Speed)
			tempLabel.Text = fmt.Sprintf("Temperature %.2f",weather.Main.Temp)
			humidityLabel.Text = fmt.Sprintf("Humidity %d",weather.Main.Humidity)

			// Refreshing the Labels for change in value 
			countryLabel.Refresh()
			windLabel.Refresh()
			tempLabel.Refresh()
			humidityLabel.Refresh()
	})
		// Setting the PlaceHolder in the selectEntry
	selectDropdown.PlaceHolder= "Select your City"
	
weatherContainer := container.NewVBox(
		headingLabel,
		container.NewHBox(layout.NewSpacer(),imgWeather,layout.NewSpacer()),
		// container.NewWithoutLayout(layout.NewSpacer(),imgWeather,layout.NewSpacer()),
		// container.NewWithoutLayout(layout.NewSpacer(),selectDropdown,layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(),selectDropdown,layout.NewSpacer()),
		countryLabel,
		windLabel,
		tempLabel,
		humidityLabel,
		layout.NewSpacer(),
	)

	// Organizing the Layout
	myWin.SetContent(container.NewBorder(nil,panelContent,nil,nil,weatherContainer),)

	// Show the window and Run the app
	myWin.Show()
}


// DEFINING THE STRUCT

//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

// clouds.go

type Clouds struct {
	All int64 `json:"all"`
}

// coord.go

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// main.go

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

// sys.go

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

// weatherelement.go

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

// wind.go

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}




