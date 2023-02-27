package pages

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Lieux(location string) fyne.CanvasObject {
	content := canvas.NewText(location, color.White)
	return content
}

func LocationToReadable(loc string) string {
	locationSplit := strings.Split(loc, "-")
	location := locationSplit[0] + " (" + locationSplit[1] + ")"
	return location
}
