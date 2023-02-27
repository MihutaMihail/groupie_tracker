package pages

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
)

func Navbar(w fyne.Window) fyne.CanvasObject {
	//artists := DataAPI.GetArtistsData()
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
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, MakeLieuxList(w)))
	})

	entry := fynex.NewCompletionEntry([]string{})
	fmt.Println(entry)

	SearchText := widget.NewEntry()
	SearchText.SetPlaceHolder("Faire une recherche")
	BtnSubmit := widget.NewButton("Submit", func() {
		log.Println("BtnSubmit")
		DataSearchBar = SearchText.Text
		SearchText.Text = ""
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, SearchBar(DataSearchBar, w)))
		//SearchBar2(DataSearchBar, w)
		//fmt.Println(DataSearchBar)
	})

	// NAVBAR ---------------------------------------------------------
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		container.New(layout.NewGridLayout(5),
			BtnHome,
			BtnArtistes,
			BtnLieux,
			SearchText,
			BtnSubmit))

	return nav
}
