package Search

import (
	"Groupie-Tracker/DataAPI"
	"fmt"
	"fyne.io/fyne/v2"
)

func SearchBar(DataSearchBar string, artists []DataAPI.Artist, w fyne.Window) {
	for _, artist := range artists {
		if artist.Name == DataSearchBar {
			fmt.Println("Trouvé cet artist " + artist.Name)
			//FindArtist(DataSearchBar, artists)
			//TODO afficher la fenetre

		} else if len(DataSearchBar) < len(artist.Name) {
			for i := 0; i < len(artist.Name)-len(DataSearchBar)+1; i++ { //-len(DataSearchBar)
				if artist.Name[i:i+len(DataSearchBar)] == DataSearchBar {
					fmt.Println("Trouvé cet artist " + artist.Name)
				}
				//if artist.Name[i:len(DataSearchBar)-1] == DataSearchBar {
				//	fmt.Println("Trouvé cet artist " + artist.Name)
				//	fmt.Println(artist.Name[i : len(DataSearchBar)-1])
				//}
			}
		}
	}
	fmt.Println()
}

//func FindArtist(name string, artists []DataAPI.Artist, w fyne.Window) {
//	for _, artist := range artists {
//		if artist.Name == name {
//			// Lance la navbar la page Artist, modifé avec la data correspondante
//			w.SetContent(container.NewBorder(Navbar(w), nil, nil, nil, Artist(artist)))
//			log.Println("Went to " + name + " (artist) page")
//		}
//	}
//}
