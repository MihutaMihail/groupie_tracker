package main

import (
	"Groupie-Tracker/DataAPI"
	"Groupie-Tracker/Geocoding"
	"Groupie-Tracker/Windows"
	"fmt"

	"fyne.io/fyne/v2/app"
)

var (
	artists     []DataAPI.Artist
	coordinates []float64
)

func main() {
	artists = DataAPI.GetArtistsData()
	fmt.Println(artists[0].Name)

	// créé l'appli et une fenêtre
	a := app.New()
	Windows.MainWindow(a)
	Windows.TestWindow(a)

	a.Run()

	coordinates = Geocoding.GetGeocodeLocation("north_carolina-usa")
	fmt.Println(coordinates)
}
