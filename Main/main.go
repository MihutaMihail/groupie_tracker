package main

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	// créé l'appli et une fenêtre
	a := app.New()
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(400, 400))

	// NAVBAR
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}), canvas.NewText("navbar", color.White))

	// IMAGE IMPORT
	r, err := fyne.LoadResourceFromURLString("https://groupietrackers.herokuapp.com/api/images/imagineDragons.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	// ITEMS
	title := canvas.NewText("title", color.Black)
	image1 := canvas.NewImageFromResource(r)
	image2 := canvas.NewText("image2", color.Black)
	text1 := canvas.NewText("text1", color.Black)
	text2 := canvas.NewText("text2", color.Black)

	// CONTENT
	body := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}), container.New(layout.NewGridLayout(5),
		layout.NewSpacer(), layout.NewSpacer(), title, layout.NewSpacer(), layout.NewSpacer(),
		layout.NewSpacer(), image1, layout.NewSpacer(), image2, layout.NewSpacer(),
		layout.NewSpacer(), text1, layout.NewSpacer(), text2, layout.NewSpacer()))

	//
	content := container.NewBorder(nav, nil, nil, nil, body)

	// TEMP affche les datas en terminal, Mihail
	DataAPI.GetArtistsData()

	// garder à la fin ; run et affiche la fenêtre, quand elle est fermé, stop l'appli
	w.SetContent(content)
	w.ShowAndRun()
}
