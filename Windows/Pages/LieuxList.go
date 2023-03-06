package pages

import (
	"Groupie-Tracker/DataAPI"
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeLieuxList(w fyne.Window) *fyne.Container {
	locationsDATA := DataAPI.GetLocationsData()
	listContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3))

	allLocations := GetLocationList(locationsDATA)

	for _, location := range allLocations {
		btn := widget.NewButton(location, nil)
		btn.OnTapped = func() {
			FindLocation(btn.Text, locationsDATA, w)
		}
		listContainer.Add(btn)
	}

	final := container.NewScroll(listContainer)
	final.SetMinSize(fyne.NewSize(1000, 600))

	return container.NewCenter(final)
}

// return a list of all Location, sorted and cleared of doubles
func GetLocationList(locationsDATA []DataAPI.Location) []string {
	var allLocations []string
	isDouble := false

	for _, locations := range locationsDATA {
		for _, location := range locations.Locations {
			location = LocationToReadable(location)
			for _, locationPassed := range allLocations {
				if locationPassed == location {
					isDouble = true
				}
			}
			if !isDouble {
				allLocations = append(allLocations, location)
			}
			isDouble = false
		}
	}

	sort.Strings(allLocations)
	return allLocations
}

// filtre une liste de string pour supprimer les doublons
func FilterDouble(entryList []string) []string {
	var finalList []string
	isDouble := false

	for _, entry := range entryList {
		for _, entryPassed := range finalList {
			if entryPassed == entry {
				isDouble = true
			}
		}
		if !isDouble {
			finalList = append(finalList, entry)
		}
		isDouble = false
	}

	return finalList
}

// Go to the location page with the readable name
func FindLocation(name string, loc []DataAPI.Location, w fyne.Window) {
	nameBase := LocationToBase(name)
	for _, location := range loc {
		for _, locationString := range location.Locations {
			if locationString == nameBase {
				// Lance la navbar la page Artist, modifé avec la data correspondante

				log.Println("Went to " + name + " (location) page")
				w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Lieux(nameBase)))
			}
		}
	}
}
