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
	// TEST Geolocalisation (not permanent)
	BtnGeoLocalisation := widget.NewButton("Geolocalisation", func() {
		log.Println("BtnGeoLocalisation")
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Geolocalisation()))
	})
	// TEST Geolocalisation (not permanent)

	BtnClose := widget.NewButton("X", func() {
		if w.FullScreen() == false {
			w.SetFullScreen(true)
		} else {
			w.Close()
		}
	})

	// NAVBAR ---------------------------------------------------------
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
			layout.NewSpacer(),
			BtnHome,
			layout.NewSpacer(),
			BtnArtistes,
			layout.NewSpacer(),
			BtnLieux,
			layout.NewSpacer(),
			// TEST Geolocalisation (not permanent)
			BtnGeoLocalisation,
			layout.NewSpacer(),
			// TEST Geolocalisation (not permanent)
			SearchText,
			BtnSubmit,
			layout.NewSpacer(),
			BtnClose,
			layout.NewSpacer()))

	return nav
}
