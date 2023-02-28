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

	var allLocations []string
	isDouble := false

	for _, locations := range locationsDATA {
		for _, location := range locations.Locations {
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

func FindLocation(name string, loc []DataAPI.Location, w fyne.Window) {
	for _, location := range loc {
		for _, locationString := range location.Locations {
			if locationString == name {
				// Lance la navbar la page Artist, modifé avec la data correspondante
				w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Lieux(locationString)))
				log.Println("Went to " + name + " (location) page")
			}
		}
	}
}
