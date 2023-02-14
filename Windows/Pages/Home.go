package pages

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Home(w fyne.Window) fyne.CanvasObject {

	// TITRE ------------------------------
	titleText := canvas.NewText("GROUPIE TRACKER", color.Black)
	titleText.TextStyle.Bold = true
	titleText.TextSize = 25
	subTitleText := canvas.NewText("Peut-être serez-vous intéresser par ces artistes ?", color.Black)
	subTitleText.TextSize = 15
	title := container.NewCenter(container.NewVBox(layout.NewSpacer(), titleText, layout.NewSpacer(), subTitleText))

	// ARTISTGRID ------------------------
	artistsGrid := artistGrid(w)

	// CREDITS ---------------------------
	creditsText := canvas.NewText("By Mihail M., Fabien A., Danny L.", color.Black)
	credits := container.NewCenter(container.NewVBox(layout.NewSpacer(), creditsText, layout.NewSpacer()))

	// BODY ------------------------------
	body := container.NewBorder(title, credits, nil, nil, artistsGrid)

	content := container.NewMax(canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 231, A: 1}), body)
	return content
}

func artistGrid(w fyne.Window) *fyne.Container {
	var alreadyInt []int
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	artists := DataAPI.GetArtistsData()

	content := container.NewAdaptiveGrid(4)

	for i := 0; i < 8; i++ {
		randomInt := random.Intn(len(artists) - 1)
		double := false
		for _, num := range alreadyInt {
			if randomInt == num {
				double = true
				break
			}
		}
		if double {
			i--
		} else {
			alreadyInt = append(alreadyInt, randomInt)
			artistToMake := artists[randomInt]
			content.Add(artistCard(artistToMake, artists, w))
		}
	}

	contentCentered := container.NewCenter(content)
	return contentCentered
}

func artistCard(artist DataAPI.Artist, artists []DataAPI.Artist, w fyne.Window) *fyne.Container {
	r, err := fyne.LoadResourceFromURLString(artist.Image)
	if err != nil {
		fmt.Println(err)
	}

	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	image := container.NewCenter(img)

	/*
		nameText := canvas.NewText(artist.Name, color.White)
		nameText.TextStyle.Bold = true
		nameText.TextSize = 20
	*/

	btn := widget.NewButton(artist.Name, nil)
	btn.OnTapped = func() {
		FindArtist(btn.Text, artists, w)
	}
	name := container.NewCenter(btn)

	creationText := canvas.NewText(strconv.Itoa(artist.CreationDate), color.White)
	creationText.TextStyle.Italic = true
	creationText.TextSize = 15
	creation := container.NewCenter(creationText)

	card := container.NewVBox(image, name, creation)

	cardFinal := container.NewMax(canvas.NewRectangle(color.Black), card)
	return cardFinal
}
