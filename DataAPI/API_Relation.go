package DataAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseRelation struct {
	Relations []Relation `json:"index"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	rObjectRelation ResponseRelation
)

// This function will get all the relations contained in the API
func getRelationsAPI() {
	getAPIUrl()

	rRelation, err := http.Get(rObjectAPI.RelationURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rDataRelation, err := ioutil.ReadAll(rRelation.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rDataRelation, &rObjectRelation)
}

// This function will show (only terminal) all the relations
func GetRelationsAPI() []Relation {
	getRelationsAPI()
	return rObjectRelation.Relations

	/*
		for i := 0; i < len(rObjectRelation.Relations); i++ {
			fmt.Print(rObjectRelation.Relations[i].Id)
			fmt.Println(rObjectRelation.Relations[i].DatesLocations)
		}
	*/
}
