package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Artist(artist DataAPI.Artist) fyne.CanvasObject {
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
	image1.FillMode = canvas.ImageFillContain

	// Bottom Left Text
	StartYText := canvas.NewText("Date de création : "+strconv.Itoa(artist.CreationDate), color.Black)
	StartYText.TextSize = 20
	StartY := container.New(layout.NewCenterLayout(), StartYText)
	FirstAlbumText := canvas.NewText("Premier album : "+artist.FirstAlbum, color.Black)
	FirstAlbumText.TextSize = 20
	FirstAlbum := container.New(layout.NewCenterLayout(), FirstAlbumText)
	BLText := container.NewVBox(StartY, FirstAlbum)

	// Bottom Right Text
	BRText := container.NewVBox(layout.NewSpacer(), DateScreen(artist.Id), layout.NewSpacer())

	//

	// CONTENT -------------------------------------------------------
	body := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}),
		container.NewBorder(
			container.New(layout.NewCenterLayout(), TopText), nil, nil, nil,
			container.New(layout.NewAdaptiveGridLayout(5),
				layout.NewSpacer(), image1, layout.NewSpacer(), BRText, layout.NewSpacer(),
				layout.NewSpacer(), BLText, layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())))

	return body
}

func DateScreen(id int) *fyne.Container {
	locations := getLocationsByID(id)
	relations := getRelationByID(id)

	BRText := container.NewVBox()

	for _, location := range locations.Locations {
		text := canvas.NewText(location+" : "+relations.DatesLocations[location][0], color.Black)
		text.TextSize = 20
		textCentered := container.New(layout.NewCenterLayout(), text)
		BRText.Add(textCentered)
	}

	return BRText
}

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
