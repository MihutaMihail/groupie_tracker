package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
	"strconv"
	"strings"
)

func SearchBar(DataSearchBar string, w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()
	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))
	DataSearchBar = IsAutocompletion(DataSearchBar)

	if len(DataSearchBar) == 0 {
		return ArtistList(0, nil, false, w)
	} else {

		// Check pour le nom
		for _, artist := range artists {
			AlredyInside := false
			if strings.ToLower(artist.Name) == strings.ToLower(DataSearchBar) { // Cas ou le nom est exactememnt pareil
				fmt.Println("Trouvé cet unique artist " + artist.Name)
				return fyne.CanvasObject(Artist(artist, w))

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

		// Check pour la date de création
		for _, artist := range artists {
			AlreadyInside := false
			if len(strconv.Itoa(artist.CreationDate)) == len(DataSearchBar) {
				if strconv.Itoa(artist.CreationDate) == DataSearchBar {
					fmt.Println("Trouvé cet artist " + artist.Name)

					btn := widget.NewButton(artist.Name, nil)
					btn.OnTapped = func() {
						FindArtist(btn.Text, artists, w)
					}
					listContainer.Add(btn)
				}
			} else if len(DataSearchBar) < len(strconv.Itoa(artist.CreationDate)) {
				for i := 0; i <= len(strconv.Itoa(artist.CreationDate))-len(DataSearchBar); i++ { //-len(DataSearchBar)
					if strconv.Itoa(artist.CreationDate)[i:i+len(DataSearchBar)] == DataSearchBar && !AlreadyInside {
						fmt.Println("Trouvé cet artist " + artist.Name)

						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
						AlreadyInside = true
					}
				}
			}
		}
	}

	fmt.Println()
	return listContainer
}

func IsArtistName() {

}

func IsAutocompletion(DataSearchBar string) string {
	NewDataSearchBar := DataSearchBar

	for i := 0; i < len(DataSearchBar); i++ {
		if string(DataSearchBar[i]) == "(" {
			NewDataSearchBar = DataSearchBar[0 : i-1]
		}
	}
	return NewDataSearchBar
}

// Fonctions pour l'autocompletion
func Autocompletion(s string, entry *fynex.CompletionEntry, artists []DataAPI.Artist) {
	var results []string
	if len(s) < 1 {
		entry.HideCompletion()
		return
	}
	results = AutoIsArtistName(s, artists, results)
	results = AutoIsCreationDate(s, artists, results)

	if len(results) == 0 {
		entry.HideCompletion()
		return
	}
	entry.SetOptions(results)
	entry.ShowCompletion()
}

func AutoIsArtistName(s string, artists []DataAPI.Artist, results []string) []string {
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
	return results
}

func AutoIsCreationDate(s string, artists []DataAPI.Artist, results []string) []string {
	tmp, _ := strconv.Atoi(s)

	for _, artist := range artists {
		AlredyInside := false

		if artist.CreationDate == tmp {
			results = append(results, artist.Name+" (Creation Date)")
		} else if len(s) <= len(strconv.Itoa(artist.CreationDate)) {
			for i := 0; i < len(strconv.Itoa(artist.CreationDate))-len(s)+1; i++ {
				if strconv.Itoa(artist.CreationDate)[i:i+len(s)] == s && !AlredyInside {
					results = append(results, artist.Name+" (Creation Date)")
					AlredyInside = true
				}
			}
		}
	}
	return results
}
