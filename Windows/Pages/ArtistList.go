package pages

import (
	"Groupie-Tracker/DataAPI"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ArtistList(w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()

	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(5))

	// création des buttons
	for _, artist := range artists {
		btn := widget.NewButton(artist.Name, nil)
		// Artist(artist) doesn't work, all button ends up the same (only last artist is remembered when you press the button)
		// So we need to find again the artist with the btn.Text
		btn.OnTapped = func() {
			FindArtist(btn.Text, artists, w)
		}
		listContainer.Add(btn)
	}
	return listContainer
}

func FindArtist(name string, artists []DataAPI.Artist, w fyne.Window) {
	for _, artist := range artists {
		if artist.Name == name {
			// Lance la navbar la page Artist, modifé avec la data correspondante
			w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Artist(artist)))
			log.Println("Went to " + name + " (artist) page")
		}
	}
}
