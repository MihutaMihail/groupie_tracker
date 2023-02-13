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
	r, err := fyne.LoadResourceFromURLString(artist.Image)
	if err != nil {
		fmt.Println(err)
	}

	// ITEMS ---------------------------------------------------------
	// top text
	titleText := canvas.NewText(artist.Name, color.Black)
	titleText.TextSize = 25
	title := container.New(layout.NewCenterLayout(), titleText)
	members := canvas.NewText("", color.Black)
	for i, member := range artist.Members {
		members.Text += member
		//si ce n'est pas le dernier elm
		if i != len(artist.Members)-1 {
			members.Text += ", "
		}
	}
	members.TextSize = 20
	TopText := container.NewVBox(title, members)

	// Image
	image1 := canvas.NewImageFromResource(r)
	image1.FillMode = canvas.ImageFillContain
	image2 := canvas.NewText("Map", color.Black)
	text1 := canvas.NewText("text1", color.Black)
	text2 := canvas.NewText("text2", color.Black)

	//

	// CONTENT -------------------------------------------------------
	body := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}),
		container.NewBorder(
			container.New(layout.NewCenterLayout(), TopText), nil, nil, nil,
			container.New(layout.NewAdaptiveGridLayout(5),
				layout.NewSpacer(), image1, layout.NewSpacer(), image2, layout.NewSpacer(),
				layout.NewSpacer(), text1, layout.NewSpacer(), text2, layout.NewSpacer())))

	return body
}
