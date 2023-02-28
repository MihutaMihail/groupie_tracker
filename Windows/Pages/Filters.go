package pages

import (
	"Groupie-Tracker/DataAPI"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	oldestCreationDate   = 9999
	youngestCreationDate = 0
	highestNbOfMembers   = 0
	listOfShowMembers    []int
)

func Filters(w fyne.Window) fyne.CanvasObject {
	content := container.NewBorder(showFilters(w), nil, nil, nil, Home(w))
	return content
}

func showFilters(w fyne.Window) fyne.CanvasObject {
	listOfShowMembers = nil

	// SLIDER ------------------------------------------------------------------
	initialValue := 2000.0 // random float
	dataSliderNumber := binding.BindFloat(&initialValue)
	sliderCreationDate := widget.NewSliderWithData(float64(getOldestCreationDate()), float64(getYoungestCreationDate()), dataSliderNumber)
	dataLabelCreationDate := widget.NewLabelWithData(binding.FloatToStringWithFormat(dataSliderNumber, "%0.0f"))

	// CHECKBOX ----------------------------------------------------------------
	checkboxContainer := container.New(layout.NewHBoxLayout())

	for i := 1; i <= getHighestNumberOfMembers(); i++ {
		checkboxMembers := widget.NewCheck(strconv.Itoa(i), nil)
		checkboxLabel, _ := strconv.Atoi(checkboxMembers.Text)
		checkboxMembers.OnChanged = func(b bool) {
			if b {
				listOfShowMembers = append(listOfShowMembers, checkboxLabel)
			} else {
				for index, number := range listOfShowMembers {
					if number == checkboxLabel {
						listOfShowMembers = append(listOfShowMembers[:index], listOfShowMembers[index+1:]...)
					}
				}
			}
		}
		checkboxContainer.Add(checkboxMembers)
	}

	// SUBMIT -----------------------------------------------------------------
	BtnSubmit := widget.NewButton("SubmitFilters", func() {
		w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, ArtistList(int(initialValue), listOfShowMembers, true, w)))
	})

	// NAVBAR ITEMS -----------------------------------------------------------
	sliderLabelCheckbox := container.NewHBox(container.NewCenter(container.NewHBox(dataLabelCreationDate)), container.NewCenter(checkboxContainer))

	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		container.New(layout.NewGridLayout(3),
			sliderCreationDate,
			sliderLabelCheckbox,
			BtnSubmit))

	return nav
}

// FUNCTIONS ----------------------------------------------------------------------
// Function that returns the oldest creation date of an album
func getOldestCreationDate() int {
	artists := DataAPI.GetArtistsData()

	for _, artist := range artists {
		if artist.CreationDate < oldestCreationDate {
			oldestCreationDate = artist.CreationDate
		}
	}

	return oldestCreationDate
}

// Function that returns the newest creation date of an album
func getYoungestCreationDate() int {
	artists := DataAPI.GetArtistsData()

	for _, artist := range artists {
		if artist.CreationDate > youngestCreationDate {
			youngestCreationDate = artist.CreationDate
		}
	}

	return youngestCreationDate
}

// Function that returns the highest amount of members from all the groups
func getHighestNumberOfMembers() int {
	nbOfMembers := 0
	artists := DataAPI.GetArtistsData()

	for _, artist := range artists {
		for i := 0; i < len(artist.Members); i++ {
			nbOfMembers += 1
		}
		if nbOfMembers > highestNbOfMembers {
			highestNbOfMembers = nbOfMembers
		}
		nbOfMembers = 0
	}

	return highestNbOfMembers
}
