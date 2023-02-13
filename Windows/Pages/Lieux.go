package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Lieux() fyne.CanvasObject {
	content := canvas.NewText("PAGE DES LIEUX", color.White)
	return content
}
