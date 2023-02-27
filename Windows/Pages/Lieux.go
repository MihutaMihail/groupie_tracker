package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Lieux(location string) fyne.CanvasObject {
	content := canvas.NewText(location, color.White)
	return content
}
