package pages

import (
	"Groupie-Tracker/DataAPI"
	"Groupie-Tracker/Search"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Navbar(w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()
	var DataSearchBar string
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

	SearchText := widget.NewEntry()
	SearchText.SetPlaceHolder("Faire une recherche")
	BtnSubmit := widget.NewButton("Submit", func() {
		DataSearchBar = SearchText.Text
		SearchText.Text = ""
		Search.SearchBar(DataSearchBar, artists, w)
		//fmt.Println(DataSearchBar)
	})

	// NAVBAR ---------------------------------------------------------
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			layout.NewSpacer(),
			BtnHome,
			layout.NewSpacer(),
			BtnArtistes,
			layout.NewSpacer(),
			BtnLieux,
			layout.NewSpacer(),
			SearchText,
			BtnSubmit,
			layout.NewSpacer()))

	return nav
}
