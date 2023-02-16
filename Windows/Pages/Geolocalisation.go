package pages

import (
	"Groupie-Tracker/Geocoding"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// TEST Geolocalisation (not permanent)
func Geolocalisation() fyne.CanvasObject {
	// For actual use, replace "Germany Mainz" by the actual location
	r, err := fyne.LoadResourceFromURLString(Geocoding.GetGeolocalisationMap(Geocoding.GetGeocodeCoordinates("Germany Mainz")))
	if err != nil {
		fmt.Println(err)
	}
	image1 := canvas.NewImageFromResource(r)
	image1.FillMode = canvas.ImageFillContain

	content := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}),
		container.NewBorder(
			container.New(layout.NewCenterLayout()), nil, nil, nil,
			container.New(layout.NewAdaptiveGridLayout(5),
				layout.NewSpacer(), image1, layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())))

	return content
}
