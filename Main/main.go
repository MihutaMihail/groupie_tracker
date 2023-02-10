package main

import (
	"Groupie-Tracker/DataAPI"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	// créé l'appli et une fenêtre
	a := app.New()
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(400, 400))
	content := container.NewVSplit(
		// NAVBAR
		container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}), canvas.NewText("navbar", color.White)),

		// BODY
		container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}), canvas.NewText("body", color.Black)),
	)

	// TEMP affche les datas en terminal, Mihail
	DataAPI.GetArtistsData()

	// garder à la fin ; run et affiche la fenêtre, quand elle est fermé, stop l'appli
	w.SetContent(content)
	w.ShowAndRun()
}
