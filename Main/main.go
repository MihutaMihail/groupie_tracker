package main

import (
	"Groupie-Tracker/Windows"

	"fyne.io/fyne/v2/app"
)

func main() {
	// créé l'appli et une fenêtre
	a := app.New()
	Windows.MainWindow(a)
	Windows.TestWindow(a)

	a.Run()
}
