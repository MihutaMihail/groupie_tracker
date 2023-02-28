package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
	"strings"
)

func SearchBar(DataSearchBar string, w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()
	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))

	if len(DataSearchBar) == 0 {
		return ArtistList(w)
	} else {
		for _, artist := range artists {
			AlredyInside := false
			if strings.ToLower(artist.Name) == strings.ToLower(DataSearchBar) { // Cas ou le nom est exactememnt pareil
				fmt.Println("Trouvé cet unique artist " + artist.Name)
				return fyne.CanvasObject(Artist(artist))

			} else if len(DataSearchBar) <= len(artist.Name) { // cas ou le terme cherhcer est plus cours que les noms
				for i := 0; i < len(artist.Name)-len(DataSearchBar)+1; i++ { //-len(DataSearchBar)
					if strings.ToLower(artist.Name[i:i+len(DataSearchBar)]) == strings.ToLower(DataSearchBar) && !AlredyInside {
						fmt.Println("Trouvé cet artist " + artist.Name)

						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
						AlredyInside = true
					}
				}
			}
		}
	}
	fmt.Println()
	return listContainer
}

func Autocompletion(s string, entry *fynex.CompletionEntry, artists []DataAPI.Artist) {
	var results []string
	if len(s) < 3 {
		entry.HideCompletion()
		return
	}
	for _, artist := range artists {
		AlredyInside := false
		if artist.Name == s {
			results = append(results, artist.Name)
		} else if len(s) <= len(artist.Name) {
			for i := 0; i < len(artist.Name)-len(s)+1; i++ {
				if strings.ToLower(artist.Name[i:i+len(s)]) == strings.ToLower(s) && !AlredyInside {
					results = append(results, artist.Name)
					AlredyInside = true
				}
			}
		}
	}
	if len(results) == 0 {
		entry.HideCompletion()
		return
	}
	entry.SetOptions(results)
	entry.ShowCompletion()
}
