package main

import (
	"Groupie-Tracker/DataAPI"
	"Groupie-Tracker/Geocoding"
	"fmt"
)

var (
	artists []DataAPI.Artist
)

func main() {
	/*artists = DataAPI.GetArtistsData()
	fmt.Println(artists[0].Name)

	// créé l'appli et une fenêtre
	a := app.New()
	Windows.MainWindow(a)
	Windows.TestWindow(a)

	a.Run()*/

	coordinates := Geocoding.GetGeocodeLocation("north_carolina-usa")
	fmt.Println(coordinates)
}
