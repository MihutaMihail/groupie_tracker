package utility

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CanvasTextWrap(upperLimit int, entryText *canvas.Text, color color.Color) *fyne.Container {
	final := container.NewVBox()
	text := entryText.Text
	string1 := ""
	string2 := ""

	// if it's too long
	if len(text) > upperLimit {

		// on parcours le texte a partir de la limite
		for i := upperLimit - 10; i < len(text); i++ {
			// quand on trouve un espace
			if text[i] == ' ' {
				string1 = text[0 : i-1]
				string2 = text[i:]
				break
			}
		}

		// recursif
		final.Add(canvas.NewText(string1, color))
		nextLines := canvas.NewText(string2, color)
		final.Add(CanvasTextWrap(upperLimit, nextLines, color))

	} else {
		// si ce n'est pas trop long
		final.Add(entryText)
	}
	return final
}
