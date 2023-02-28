package pages

import (
	"Groupie-Tracker/Geocoding"
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Lieux(location string) fyne.CanvasObject {
	r, err := fyne.LoadResourceFromURLString(Geocoding.GetGeolocalisationMap(Geocoding.GetGeocodeCoordinates(location)))
	if err != nil {
		fmt.Println(err)
	}

	mapItem := canvas.NewImageFromResource(r)
	mapItem.FillMode = canvas.ImageFillOriginal

	locationReadable := LocationToReadable(location)
	title := canvas.NewText(locationReadable, color.Black)
	title.TextSize = 25
	title.TextStyle.Bold = true

	content := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}),
		container.NewBorder(
			container.NewCenter(container.NewVBox(layout.NewSpacer(), title)), nil, nil, nil,
			container.NewCenter(container.NewGridWithColumns(1, mapItem))))
	return content
}

// make the name of location from base to readable
func LocationToReadable(loc string) string {
	locationSplit := strings.Split(loc, "-")
	location := locationSplit[1] + " (" + locationSplit[0] + ")"
	return location
}

// make the name of location from readable to base
func LocationToBase(loc string) string {
	locationSplit := strings.Split(loc, " (")
	location := locationSplit[1][:len(locationSplit[1])-1] + "-" + locationSplit[0]
	return location
}
