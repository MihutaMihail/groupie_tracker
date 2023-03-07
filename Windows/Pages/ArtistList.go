package pages

import (
	"Groupie-Tracker/DataAPI"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

	titleNoResulsFound := canvas.NewText("NO RESULTS FOUND", color.White)
	titleNoResulsFound.TextSize = 25
	titleNoResulsFound.TextStyle.Bold = true
	contentNoResultsFound := container.New(layout.NewCenterLayout(), titleNoResulsFound)

	if (disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert == "") {
		filtersOn = false
	}

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
			cleanStart()
			// SLIDER + MEMBERS
		} else if (!disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate == "") && (locationConcert == "") {
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
			cleanStart()
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
			cleanStart()
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
			cleanStart()
			// SLIDER + MEMBERS + FIRST ALBUM
		} else if (!disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate != "") && (locationConcert == "") {
			for _, artist := range artists {
				for _, number := range nbOfMembers {
					if artist.CreationDate == numberSlider && len(artist.Members) == number && artist.FirstAlbum == firstAlbumDate {
						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
					}
				}
			}
			cleanStart()
			// SLIDER + MEMBERS + LOCATION CONCERT
		} else if (!disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				if artist.CreationDate == numberSlider {
					for _, number := range nbOfMembers {
						for _, idLocation := range listIDLocations {
							if artist.Id == idLocation && len(artist.Members) == number {
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
			cleanStart()
			// SLIDER + FIRST ALBUM + LOCATION CONCERT
		} else if (!disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate != "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				if artist.CreationDate == numberSlider {
					for _, idLocation := range listIDLocations {
						if artist.Id == idLocation && artist.FirstAlbum == firstAlbumDate {
							btn := widget.NewButton(artist.Name, nil)
							btn.OnTapped = func() {
								FindArtist(btn.Text, artists, w)
							}
							listContainer.Add(btn)
						}
					}
				}
			}
			cleanStart()
			// LOCATION CONCERT
		} else if (disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
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
			cleanStart()
			// LOCATION CONCERT + FIRST ALBUM
		} else if (disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				for _, idLocation := range listIDLocations {
					if artist.Id == idLocation && artist.FirstAlbum == firstAlbumDate {
						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
					}
				}
			}
			cleanStart()
			// LOCATION CONCERT + MEMBERS
		} else if (disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				for _, number := range nbOfMembers {
					for _, idLocation := range listIDLocations {
						if artist.Id == idLocation && len(artist.Members) == number {
							btn := widget.NewButton(artist.Name, nil)
							btn.OnTapped = func() {
								FindArtist(btn.Text, artists, w)
							}
							listContainer.Add(btn)
						}
					}
				}
			}
			cleanStart()
			// LOCATION CONCERT + MEMBERS + FIRST ALBUM
		} else if (disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate == "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				for _, number := range nbOfMembers {
					for _, idLocation := range listIDLocations {
						if artist.Id == idLocation && len(artist.Members) == number && artist.FirstAlbum == firstAlbumDate {
							btn := widget.NewButton(artist.Name, nil)
							btn.OnTapped = func() {
								FindArtist(btn.Text, artists, w)
							}
							listContainer.Add(btn)
						}
					}
				}
			}
			cleanStart()
			// FIRST ALBUM
		} else if (disableSlider) && (len(nbOfMembers) == 0) && (firstAlbumDate != "") && (locationConcert == "") {
			for _, artist := range artists {
				if artist.FirstAlbum == firstAlbumDate {
					btn := widget.NewButton(artist.Name, nil)
					btn.OnTapped = func() {
						FindArtist(btn.Text, artists, w)
					}
					listContainer.Add(btn)
				}
			}
			cleanStart()
			// FIRST ALBUM + MEMBERS
		} else if (disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate != "") && (locationConcert == "") {
			for _, artist := range artists {
				for _, number := range nbOfMembers {
					if artist.FirstAlbum == firstAlbumDate && len(artist.Members) == number {
						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
					}
				}
			}
			cleanStart()
			// MEMBERS
		} else if (disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate == "") && (locationConcert == "") {
			for _, artist := range artists {
				for _, number := range nbOfMembers {
					if len(artist.Members) == number {
						btn := widget.NewButton(artist.Name, nil)
						btn.OnTapped = func() {
							FindArtist(btn.Text, artists, w)
						}
						listContainer.Add(btn)
					}
				}
			}
			cleanStart()
			// SLIDER + MEMBERS + FIRST ALBUM + LOCATION CONCERT
		} else if (!disableSlider) && (len(nbOfMembers) > 0) && (firstAlbumDate != "") && (locationConcert != "") {
			listIDLocations = getArtistsbyLocation(locationConcert)

			for _, artist := range artists {
				for _, number := range nbOfMembers {
					for _, idLocation := range listIDLocations {
						if artist.Id == idLocation && len(artist.Members) == number && artist.FirstAlbum == firstAlbumDate && artist.CreationDate == numberSlider {
							btn := widget.NewButton(artist.Name, nil)
							btn.OnTapped = func() {
								FindArtist(btn.Text, artists, w)
							}
							listContainer.Add(btn)
						}
					}
				}
			}
			cleanStart()
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

	if len(listContainer.Objects) == 0 {
		return contentNoResultsFound
	} else {
		return listContainer
	}
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
			if locationArray == LocationToBase(locationFind) {
				idLocations = append(idLocations, location.Id)
			}
		}
	}

	return idLocations
}

func cleanStart() {
	initialValueSlider = 2000
	boolDisableSlider = false
	listOfShowMembers = nil
	firstAlbumDate = ""
	locationConcert = ""
}
