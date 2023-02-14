package main

import (
	"Groupie-Tracker/DataAPI"
	"Groupie-Tracker/Windows"

	"fyne.io/fyne/v2/app"
)

var (
	artists             []DataAPI.Artist
	coordinatesResponse []float64
)

func main() {
	// créé l'appli et une fenêtre
	a := app.New()
	Windows.MainWindow(a)
	//Windows.TestWindow(a)

	a.Run()
}
