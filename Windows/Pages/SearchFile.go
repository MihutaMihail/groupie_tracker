package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func SearchBar(DataSearchBar string, w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()
	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))

	if len(DataSearchBar) == 0 {
		return ArtistList(w)
	} else {
		for _, artist := range artists {
			if strings.ToLower(artist.Name) == strings.ToLower(DataSearchBar) { // Cas ou le nom est exactememnt pareil
				fmt.Println("Trouvé cet unique artiste" + artist.Name)
				listContainer = fyne.NewContainer(fyne.CanvasObject(Artist(artist)))

			} else if len(DataSearchBar) <= len(artist.Name) { // cas ou le terme cherhcer est plus cours que les noms
				for i := 0; i < len(artist.Name)-len(DataSearchBar)+1; i++ { //-len(DataSearchBar)
					if strings.ToLower(artist.Name[i:i+len(DataSearchBar)]) == strings.ToLower(DataSearchBar) {
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
	}
	fmt.Println()
	return listContainer
}
