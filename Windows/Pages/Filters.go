package pages

import (
	"Groupie-Tracker/DataAPI"
	utility "Groupie-Tracker/Utility"
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
	oldestCreationDate      = 9999
	youngestCreationDate    = 0
	highestNbOfMembers      = 0
	listOfShowMembers       []int
	listOfFirstAlbumDates   []string
	listOfLocationsConcerts []string
	initialValueSlider      = 2000.0 // random float
	boolDisableSlider       = false
	firstAlbumDate          string
	locationConcert         string
)

func Filters(w fyne.Window) fyne.CanvasObject {
	content := container.NewBorder(showFilters(w), nil, nil, nil, ArtistList(0, false, nil, "", "", false, w))
	return content
}

func showFilters(w fyne.Window) fyne.CanvasObject {
	listOfShowMembers = nil

	// SLIDER CREATION DATE  ------------------------------------------------------------------
	bindingValueSlider := binding.BindFloat(&initialValueSlider)
	sliderCreationDate := widget.NewSliderWithData(float64(getOldestCreationDate()), float64(getYoungestCreationDate()), bindingValueSlider)
	numberCreationDate := widget.NewLabelWithData(binding.FloatToStringWithFormat(bindingValueSlider, "%0.0f"))
	textCreationDate := canvas.NewText("Career Starting Year", color.White)

	// CHECKBOX DISABLE SLIDER ----------------------------------------------------------------
	checkboxDisableSlider := widget.NewCheck("Don't use slider", func(disableSlider bool) {
		if disableSlider {
			boolDisableSlider = true
		} else {
			boolDisableSlider = false
		}
	})

	// CHECKBOX NUMBERS MEMBERS ----------------------------------------------------------------
	checkboxContainer := container.New(layout.NewHBoxLayout())
	textNbMembers := canvas.NewText("N° Members", color.White)

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

	// FIRST ALBUM DATE -----------------------------------------------------------------
	textFirstAlbumDate := canvas.NewText("First Album Date", color.White)

	for _, artist := range DataAPI.GetArtistsData() {
		listOfFirstAlbumDates = append(listOfFirstAlbumDates, artist.FirstAlbum)
	}
	listOfFirstAlbumDates = FilterDouble(listOfFirstAlbumDates)
	selectFirstAlbumDate := widget.NewSelect(utility.SortDates(listOfFirstAlbumDates), func(firstAlbumDate string) {})
	selectFirstAlbumDate.OnChanged = func(choice string) {
		firstAlbumDate = choice
	}
	// LOCATIONS OF CONCERTS -----------------------------------------------------------------
	textLocationConcert := canvas.NewText("Location Concert", color.White)

	listOfLocationsConcerts = GetLocationList(DataAPI.GetLocationsData())
	selectLocationConcert := widget.NewSelect(listOfLocationsConcerts, func(locationConcert string) {})
	selectLocationConcert.OnChanged = func(choice string) {
		locationConcert = choice
	}

	// NAVBAR ITEMS -----------------------------------------------------------
	sliderInfo := container.NewHBox(container.NewCenter(container.NewHBox(textCreationDate, numberCreationDate, checkboxDisableSlider)))
	checkboxInfo := container.NewHBox(container.NewCenter(container.NewHBox(textNbMembers, checkboxContainer)))
	firstAlbumInfo := container.NewHBox(container.NewCenter(container.NewHBox(textFirstAlbumDate, selectFirstAlbumDate)))
	locationConcertInfo := container.NewHBox(container.NewCenter(container.NewHBox(textLocationConcert, selectLocationConcert)))

	nav := container.NewMax(canvas.NewRectangle(color.RGBA{R: 31, G: 31, B: 35, A: 1}),
		container.New(layout.NewGridLayout(3),
			sliderCreationDate,
			sliderInfo,
			checkboxInfo,
			firstAlbumInfo,
			locationConcertInfo))

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
