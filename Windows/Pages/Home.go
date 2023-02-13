package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Home() fyne.CanvasObject {
	content := canvas.NewText("HOME PAGE", color.White)
	return content
}
