package pages

import (
	"Groupie-Tracker/DataAPI"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
	"image/color"
	"log"
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
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, MakeLieuxList(w)))
	})

	entry := fynex.NewCompletionEntry([]string{})
	entry.SetPlaceHolder("Coucou")
	entry.OnChanged = func(s string) {
		Autocompletion(s, entry, artists)
	}

	BtnSubmit2 := widget.NewButton("Submit", func() {
		log.Println("BtnSubmit")
		DataSearchBar = entry.Text
		entry.Text = ""
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, SearchBar(DataSearchBar, w)))
	})

	// NAVBAR ---------------------------------------------------------
	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		container.New(layout.NewGridLayout(5),
			BtnHome,
			BtnArtistes,
			BtnLieux,
			entry,
			BtnSubmit2))

	return nav
}
