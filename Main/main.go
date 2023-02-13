package main

import (
	"Groupie-Tracker/DataAPI"
	"Groupie-Tracker/Geocoding"
	"Groupie-Tracker/Windows"
	"fmt"

	"fyne.io/fyne/v2/app"
)

var (
	artists             []DataAPI.Artist
	coordinatesResponse []float64
)

func main() {
	// TEST
	coordinatesResponse = Geocoding.GetGeocodeLocation("north_carolina-usa")
	fmt.Println(coordinatesResponse)

	// créé l'appli et une fenêtre
	a := app.New()
	Windows.MainWindow(a)
	Windows.TestWindow(a)

	a.Run()
}
