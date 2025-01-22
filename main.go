package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("NEW WINDOW")
	w.Resize(fyne.NewSize(600,400))

	headText := widget.NewLabel("Title")
	headerContainer := container.NewCenter(headText)

	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("ENTRY1...")
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("ENTRY2...")

	formContainer := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Entry111111", Widget: entry1},
			{Text: "Entry2", Widget: entry2}},
		OnSubmit: func() {
			cfgText := container.NewVBox(
				widget.NewLabel("confirm line1"),
				widget.NewLabel(entry1.Text + entry2.Text),
			)
			dialog.ShowCustomConfirm("title","Yes","No", cfgText, func(b bool) {
				if b {
					fmt.Println("confirmed:", entry1.Text + entry2.Text)
				}
			}, w)
			log.Println("1:", entry1.Text)
			log.Println("2:", entry2.Text)
		},
		OnCancel: func() {
			w.Close()
		},
	}

	content := container.NewGridWithColumns(
		1,
		headerContainer,
		formContainer,
		layout.NewSpacer(),
	)

	w.SetContent(content)
	w.ShowAndRun()
}