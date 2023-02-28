package pages

import (
	"Groupie-Tracker/DataAPI"
	utility "Groupie-Tracker/Utility"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Artist(artist DataAPI.Artist, w fyne.Window) fyne.CanvasObject {
	// IMAGE IMPORT
	r, err := fyne.LoadResourceFromURLString(artist.Image)
	if err != nil {
		fmt.Println(err)
	}

	// ITEMS ---------------------------------------------------------
	// top text
	titleText := canvas.NewText(artist.Name, color.Black)
	titleText.TextSize = 25
	title := container.New(layout.NewCenterLayout(), titleText)
	members := canvas.NewText("", color.Black)
	for i, member := range artist.Members {
		members.Text += member
		//si ce n'est pas le dernier elm
		if i != len(artist.Members)-1 {
			members.Text += ", "
		}
	}
	members.TextSize = 20
	TopText := container.NewVBox(title, members)

	// Image
	image1 := canvas.NewImageFromResource(r)
	image1.FillMode = canvas.ImageFillOriginal

	// Bottom Left Text
	StartYText := canvas.NewText("Date de création : "+strconv.Itoa(artist.CreationDate), color.Black)
	StartYText.TextSize = 20
	StartY := container.New(layout.NewCenterLayout(), StartYText)
	FirstAlbumText := canvas.NewText("Premier album : "+artist.FirstAlbum, color.Black)
	FirstAlbumText.TextSize = 20
	FirstAlbum := container.New(layout.NewCenterLayout(), FirstAlbumText)
	BLText := container.NewVBox(StartY, FirstAlbum)

	// ImageSide
	LeftSide := container.NewVBox(image1, BLText)

	// DateGrid
	DateGrid := container.NewVBox(layout.NewSpacer(), container.NewMax(canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 115, A: 1}), DateScreen(artist.Id, w)), layout.NewSpacer())

	//

	// CONTENT -------------------------------------------------------
	body := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}),
		container.NewBorder(
			container.New(layout.NewCenterLayout(), TopText), nil, nil, nil,
			container.NewCenter(container.NewGridWithColumns(3, LeftSide, layout.NewSpacer(), DateGrid))))

	return body
}

func DateScreen(id int, w fyne.Window) *fyne.Container {
	final := container.NewVBox()
	// get data
	locations := getLocationsByID(id)
	relations := getRelationByID(id)

	// make rows one by one
	for _, location := range locations.Locations {
		// add the label/ Location
		locationReadable := LocationToReadable(location)
		locationText := locationReadable
		locationBtn := widget.NewButton(locationText, nil)
		locationBtn.OnTapped = func() {
			FindLocation(locationBtn.Text, DataAPI.GetLocationsData(), w)
		}

		final.Add(layout.NewSpacer())
		final.Add(locationBtn)

		// add the dates
		final.Add(makeLocationDateList(id, location, relations))
		final.Add(layout.NewSpacer())

	}
	finalScroll := container.NewScroll(container.NewCenter(final))
	finalScroll.SetMinSize(fyne.NewSize(300, 300))
	return container.NewVBox(finalScroll)
}

func makeLocationDateList(id int, location string, relations DataAPI.Relation) *fyne.Container {
	final := container.NewHBox()
	neededRelation := relations.DatesLocations[location]
	neededRelation = utility.SortDates(neededRelation)
	textFinal := ""

	// make columns one by one
	for i, date := range neededRelation {
		textFinal += date

		if i != len(neededRelation)-1 {
			textFinal += ", "
		}
	}
	textUnwrapped := canvas.NewText(textFinal, color.Black)
	textWrapped := utility.CanvasTextWrap(45, textUnwrapped, color.Black)
	final.Add(textWrapped)

	return final
}

// get the locations DATA with corresponding ID
func getLocationsByID(Id int) DataAPI.Location {
	locations := DataAPI.GetLocationsData()
	for _, location := range locations {
		if location.Id == Id {
			return location
		}
	}
	log.Println("ERROR : Searched for " + strconv.Itoa(Id) + " (date) and NOT FOUND")
	return locations[0]
}

// get the locations DATA with corresponding ID
func getRelationByID(Id int) DataAPI.Relation {
	relations := DataAPI.GetRelationsAPI()
	for _, relation := range relations {
		if relation.Id == Id {
			return relation
		}
	}
	log.Println("ERROR : Searched for " + strconv.Itoa(Id) + " (relation) and NOT FOUND")
	return relations[0]
}
