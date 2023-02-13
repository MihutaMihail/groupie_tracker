package pages

import (
	"Groupie-Tracker/DataAPI"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ArtistList(w fyne.Window) fyne.CanvasObject {
	//content := canvas.NewText("PAGE DES ARTISTES", color.White)

	artists := DataAPI.GetArtistsData()

	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))

	for _, artist := range artists {
		btn := widget.NewButton(artist.Name, func() {
			//findArtist(btn.name)
		})
		listContainer.Add(btn)
	}
	return listContainer
}

func findArtist(name string) {

}
