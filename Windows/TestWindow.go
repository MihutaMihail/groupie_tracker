package Windows

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TestWindow(a fyne.App) {
	w := a.NewWindow("TEST WINDOW")
	w.Resize(fyne.NewSize(400, 400))

	btn1 := widget.NewButton("artist1", func() {
		log.Println("btn1")
		// Code
	})
	btn2 := widget.NewButton("artist2", func() {
		log.Println("btn2")
		// Code
	})
	btn3 := widget.NewButton("artist3", func() {
		log.Println("btn3")
		// Code
	})
	btn4 := widget.NewButton("artist4", func() {
		log.Println("btn4")
		// Code
	})
	btn5 := widget.NewButton("artist5", func() {
		log.Println("btn5")
		// Code
	})
	btn6 := widget.NewButton("artist6", func() {
		log.Println("btn6")
		// Code
	})
	btn7 := widget.NewButton("artist7", func() {
		log.Println("btn7")
		// Code
	})
	btn8 := widget.NewButton("artist8", func() {
		log.Println("btn8")
		// Code
	})

	content := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2), btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)

	w.SetContent(content)

	w.Show()
}
