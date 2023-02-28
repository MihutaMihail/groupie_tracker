package Windows

import (
	pages "Groupie-Tracker/Windows/Pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MainWindow(a fyne.App) {
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(1450, 700))
	w.SetMaster()

	// à désactiver pour les tests et l'accès console
	//w.SetFullScreen(true)

	body := pages.Home(w)
	nav := pages.Navbar(w)
	content := container.NewBorder(nav, nil, nil, nil, body)

	// garder à la fin ; run et affiche la fenêtre, quand elle est fermé, stop l'appli
	w.SetContent(content)
	w.Show()
}
