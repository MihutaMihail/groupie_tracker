package Geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Declaration of variables
var (
	openCageKeyAPI      = "7fa0cecf85af4b48bb0f64bc929f7070" // Mihail Key
	coordinatesLocation []float64
	longitude           float64
	latitude            float64
	values              map[string]interface{}
)

// Function that will return the geographical coordinates of a place
func GetGeocodeLocation(locationMaps string) []float64 {
	url.QueryEscape(locationMaps)
	urlOpenCage := ("https://api.opencagedata.com/geocode/v1/json?q=" + locationMaps + "&key=" + openCageKeyAPI)
	fmt.Println("Geocode URL " + urlOpenCage)

	response, err := http.Get(urlOpenCage)
	if err != nil {
		panic("Error retrieving response")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Error retrieving response")
	}

	json.Unmarshal(body, &values)

	// This block of code will search through the API's response to get the important information
	// Usually, there will be more than 1 location found so we'll take the first one and get out of the loop
out:
	for _, v := range values["results"].([]interface{}) {
		for i2, v2 := range v.(map[string]interface{}) {
			if i2 == "geometry" {
				latitude = v2.(map[string]interface{})["lat"].(float64)
				longitude = v2.(map[string]interface{})["lng"].(float64)
				break out
			}
		}
	}

	coordinatesLocation = append(coordinatesLocation, latitude, longitude)
	return coordinatesLocation
}
