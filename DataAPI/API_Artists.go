package DataAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var (
	rObjectArtists []Artist
)

// This function will get all the artists contained in the API
func getArtistsAPI() {
	getAPIUrl()

	rArtists, err := http.Get(rObjectAPI.ArtistURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rDataArtists, err := ioutil.ReadAll(rArtists.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rDataArtists, &rObjectArtists)

}

// This function will show (only terminal) all the artists
func GetArtistsData() []Artist {
	getArtistsAPI()

	return rObjectArtists

	// Get all artists
	/*for num := range rObjectArtists {
		fmt.Print(rObjectArtists[num].Id)
		fmt.Print(rObjectArtists[num].Image)
		fmt.Print(rObjectArtists[num].Name)
		fmt.Print(rObjectArtists[num].Members)
		fmt.Print(rObjectArtists[num].CreationDate)
		fmt.Print(rObjectArtists[num].FirstAlbum)
		fmt.Println()
	}*/
}
