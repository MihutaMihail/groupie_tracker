package DataAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	ArtistURL   string `json:"artists"`
	LocationURL string `json:"locations"`
	DateURL     string `json:"dates"`
	RelationURL string `json:"relation"`
}

var (
	rObjectAPI Response
)

// This function's purpose is to retrieve the data stored in the API. All information will be stored in the RESPONSE struct
func getAPIUrl() {
	rAPI, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rDataAPI, err := ioutil.ReadAll(rAPI.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rDataAPI, &rObjectAPI)
}
