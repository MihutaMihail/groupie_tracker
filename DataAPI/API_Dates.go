package DataAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseDates struct {
	Dates []Date `json:"index"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var (
	rObjectDates ResponseDates
)

// This function will get all the dates contained in the API
func getDatesAPI() {
	getAPIUrl()

	rDates, err := http.Get(rObjectAPI.DateURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rDataDates, err := ioutil.ReadAll(rDates.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rDataDates, &rObjectDates)
}

// This function will show (only terminal) all the dates
func GetDatesAPI() []Date {
	getDatesAPI()
	return rObjectDates.Dates

	/*
		for i := 0; i < len(rObjectDates.Dates); i++ {
			fmt.Print(rObjectDates.Dates[i].Id)
			fmt.Print("   ")
			fmt.Println(rObjectDates.Dates[i].Dates)
		}
	*/
}
