package DataAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseLocations struct {
	Locations []Location `json:"index"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

var (
	rObjectLocations ResponseLocations
)

// This function will get all the locations contained in the API
func getLocationsAPI() {
	getAPIUrl()

	rLocations, err := http.Get(rObjectAPI.LocationURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rDataLocations, err := ioutil.ReadAll(rLocations.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rDataLocations, &rObjectLocations)
}

// This function will show (only terminal) all the locations
func GetLocationsData() []Location {
	getLocationsAPI()
	return rObjectLocations.Locations

	/*
		for i := 0; i < len(rObjectLocations.Locations); i++ {
			fmt.Print(rObjectLocations.Locations[i].Id)
			fmt.Println(rObjectLocations.Locations[i].Locations)
		}
	*/
}
