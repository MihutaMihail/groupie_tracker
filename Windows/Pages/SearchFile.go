package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SearchBar(DataSearchBar string, w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()

	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))

	for _, artist := range artists {
		if artist.Name == DataSearchBar { // Cas ou le nom est exactememnt pareil
			fmt.Println("Trouvé cet unique artiste" + artist.Name)
			//FindArtist(DataSearchBar, artists)
			//TODO afficher la fenetre

		} else if len(DataSearchBar) <= len(artist.Name) { // cas ou le terme cherhcer est plus cours que les noms
			for i := 0; i < len(artist.Name)-len(DataSearchBar)+1; i++ { //-len(DataSearchBar)
				if artist.Name[i:i+len(DataSearchBar)] == DataSearchBar {
					fmt.Println("Trouvé cet artist " + artist.Name)

					btn := widget.NewButton(artist.Name, nil)
					btn.OnTapped = func() {
						FindArtist(btn.Text, artists, w)
					}
					listContainer.Add(btn)
				}
			}
		}
	}
	fmt.Println()
	return listContainer
}
