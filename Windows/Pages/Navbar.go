package pages

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Navbar(w fyne.Window) fyne.CanvasObject {
	// NAVBAR BUTTONS ------------------------------------------------
	BtnHome := widget.NewButton("Home", func() {
		log.Println("BtnHome")
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Home(w)))
	})
	BtnArtistes := widget.NewButton("Artistes", func() {
		log.Println("BtnArtistes")
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, ArtistList(w)))
	})
	BtnLieux := widget.NewButton("Lieux", func() {
		log.Println("BtnLieux")
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Lieux()))
	})

	// NAVBAR ---------------------------------------------------------
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		fyne.NewContainerWithLayout(layout.NewGridLayout(7),
			layout.NewSpacer(),
			BtnHome,
			layout.NewSpacer(),
			BtnArtistes,
			layout.NewSpacer(),
			BtnLieux,
			layout.NewSpacer()))

	return nav
}
