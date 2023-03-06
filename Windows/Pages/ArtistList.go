package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	listIDLocations []int
)

func ArtistList(numberSlider int, disableSlider bool, nbOfMembers []int, firstAlbumDate string, locationConcert string, filtersOn bool, w fyne.Window) fyne.CanvasObject {
	artists := DataAPI.GetArtistsData()

	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(5))

	if (disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert == "") {
		filtersOn = false
	}

	// REMOVE ONCE TESTS ARE COMPLETED
	fmt.Println(numberSlider)
	fmt.Println(disableSlider)
	fmt.Println(nbOfMembers)
	fmt.Println(firstAlbumDate)
	fmt.Println(locationConcert)
	fmt.Println(filtersOn)
	// REMOVE ONCE TESTS ARE COMPLETED

	// FILTERS ON -----------------------------------------------------------
	if filtersOn {
		// SLIDER
		if (!disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert == "") {
			for _, artist := range artists {
				if artist.CreationDate == numberSlider {
					btn := widget.NewButton(artist.Name, nil)
					btn.OnTapped = func() {
						FindArtist(btn.Text, artists, w)
					}
					listContainer.Add(btn)
				}
			}
			// SLIDER + MEMBERS
		} else if (!disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert == "") {
			for _, artist := range artists {
				for _, number := range nbOfMembers {
					if artist.CreationDate == numberSlider && len(artist.Members) == number {
						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
					}
				}
			}
			// SLIDER + FIRST ALBUM
		} else if (!disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate != "") && (locationConcert == "") {
			for _, artist := range artists {
				if artist.CreationDate == numberSlider && artist.FirstAlbum == firstAlbumDate {
					btn := widget.NewButton(artist.Name, nil)
					btn.OnTapped = func() {
						FindArtist(btn.Text, artists, w)
					}
					listContainer.Add(btn)
				}
			}
			// SLIDER + LOCATION CONCERT
		} else if (!disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				if artist.CreationDate == numberSlider {
					for _, idLocation := range listIDLocations {
						if artist.Id == idLocation {
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
	} else {
		// FILTERS OFF -----------------------------------------------------------
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
	}

	return listContainer
}

func FindArtist(name string, artists []DataAPI.Artist, w fyne.Window) {
	for _, artist := range artists {
		if artist.Name == name {
			// Lance la navbar la page Artist, modifé avec la data correspondante
			w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Artist(artist, w)))
			log.Println("Went to " + name + " (artist) page")
		}
	}
}

// Get all IDs (artists) that play in a location
func getArtistsbyLocation(locationFind string) []int {
	locations := DataAPI.GetLocationsData()
	idLocations := []int{}

	for _, location := range locations {
		for _, locationArray := range location.Locations {
			if locationArray == locationFind {
				idLocations = append(idLocations, location.Id)
			}
		}
	}

	return idLocations
}
