package Windows

import (
	"Groupie-Tracker/DataAPI"
	pages "Groupie-Tracker/Windows/Pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MainWindow(a fyne.App) {
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(750, 400))
	w.SetMaster()

	body := pages.Home()
	nav := pages.Navbar(w)
	content := container.NewBorder(nav, nil, nil, nil, body)

	// TEMP affche la datas en terminal, Mihail
	DataAPI.GetArtistsData()

	// garder à la fin ; run et affiche la fenêtre, quand elle est fermé, stop l'appli
	w.SetContent(content)
	w.Show()
}
