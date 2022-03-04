package main

import (
  "io/ioutil"
  "strconv"
  "fyne.io/fyne/v2"
//   "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/storage"
  "fyne.io/fyne/v2/widget"
)

var count int = 0

func showTextEditor() {
  content := container.NewVBox(
    container.NewHBox(
      widget.NewLabel("Text Editor"),
    ),
  )

content.Add(widget.NewButton("Add New File", func() {
	content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
	count++
}))

input := widget.NewMultiLineEntry()
input.SetPlaceHolder("Enter text...")
input.Resize(fyne.NewSize(400, 400))

// func ShowFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window)
// save button functionality
saveBtn := widget.NewButton("Save", func() {
saveFileDialog := dialog.NewFileSave(
	func(sv fyne.URIWriteCloser, _ error) {
	textData := []byte(input.Text)
	sv.Write(textData)
	}, myWin)
saveFileDialog.SetFileName("New File " + strconv.Itoa(count) + ".txt")
saveFileDialog.Show()
})

  //Open Button functionality
  openBtn := widget.NewButton("Open", func() {
    openFileDialog := dialog.NewFileOpen(
      func(re fyne.URIReadCloser, _ error) {
        ReadData, _ := ioutil.ReadAll(re)

        output := fyne.NewStaticResource("New File ", ReadData)

        viewData := widget.NewMultiLineEntry()
        viewData.SetText(string(output.StaticContent))

        saveBtn := widget.NewButton("Save", func() {
          saveFileDialog := dialog.NewFileSave(
            func(sv fyne.URIWriteCloser, _ error) {
              textData := []byte(input.Text)
              sv.Write(textData)
            }, myWin)
          saveFileDialog.SetFileName("New File " + strconv.Itoa(count) + ".txt")
          saveFileDialog.Show()
        })
        w := fyne.CurrentApp().NewWindow(string(output.StaticName))
        w.SetContent(container.NewBorder(nil,saveBtn, nil,nil, container.NewScroll(viewData)))
        w.Resize(fyne.NewSize(400, 400))
        w.Show()

      }, myWin)
      openFileDialog.SetFilter(storage.NewExtensionFileFilter([] string{".txt"}))
      openFileDialog.Show()
  })


  textEditorContainer:= container.NewVBox(
      content,
      input,
      container.NewHBox(
        saveBtn,
        openBtn,
      ),
    )

myWin := myApp.NewWindow("Text Editor")
myWin.Resize(fyne.NewSize(600, 600))
myWin.SetContent(container.NewBorder(nil,DeskBtn,nil,nil,textEditorContainer),)
myWin.Show()
}
