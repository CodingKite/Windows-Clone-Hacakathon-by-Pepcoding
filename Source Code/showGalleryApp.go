package main

import (
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	
)

func showGalleryApp(myWin fyne.Window) {
	// myApp := app.New()
	// myWin := myApp.NewWindow("Image Gallery")
	// myWin.Resize(fyne.NewSize(800,600))
	tabs := container.NewAppTabs()
	// root_src := "C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project2 Gallery App"
	
	root_src := "C:\\Users\\rahul\\Downloads\\Video\\Pepcoding fyne vids\\Project6\\newfile\\images\\"

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }

	// var picsArray []string;
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false{
			extension := strings.Split(file.Name(),".")[1];
			if extension == "png" || extension == "jpg"{
				// picsArray = append(picsArray,root_src+"\\"+file.Name());
				// picsArray = append(picsArray,);
					image := canvas.NewImageFromFile(root_src+file.Name())
					image.FillMode = canvas.ImageFillContain
					tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
    }
 	
	tabs.SetTabLocation(container.TabLocationLeading)
	// myWin.SetContent(container.NewVBox(tabs))
	// myWin.SetContent(tabs)
	myWin.SetContent(container.NewBorder(nil,panelContent,nil,nil,tabs))

// myWin.ShowAndRun()
	myWin.Show()

}
