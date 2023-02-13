package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Artist(artist DataAPI.Artist) fyne.CanvasObject {
	// IMAGE IMPORT
	r, err := fyne.LoadResourceFromURLString("https://groupietrackers.herokuapp.com/api/images/imagineDragons.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	// ITEMS ---------------------------------------------------------
	title := canvas.NewText(artist.Name, color.Black)
	title.TextSize = 25
	image1 := canvas.NewImageFromResource(r)
	image2 := canvas.NewText("Map", color.Black)
	text1 := canvas.NewText("text1", color.Black)
	text2 := canvas.NewText("text2", color.Black)

	// CONTENT -------------------------------------------------------
	body := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}), container.New(layout.NewGridLayout(5),
		layout.NewSpacer(), layout.NewSpacer(), title, layout.NewSpacer(), layout.NewSpacer(),
		layout.NewSpacer(), image1, layout.NewSpacer(), image2, layout.NewSpacer(),
		layout.NewSpacer(), text1, layout.NewSpacer(), text2, layout.NewSpacer()))

	//

	return body
}
