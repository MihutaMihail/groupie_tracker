package Geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Declaration of variables
var (
	openCageKeyAPI      = "7fa0cecf85af4b48bb0f64bc929f7070" // Mihail Key
	coordinatesLocation []float64
	longitude           float64
	latitude            float64
	valuesGeocoding     map[string]interface{}
)

// Function that will return the geographical coordinates of a place
func GetGeocodeCoordinates(locationMaps string) []float64 {
	coordinatesLocation = nil

	locationMaps = makeLocationURLValid(locationMaps)
	urlOpenCage := ("https://api.opencagedata.com/geocode/v1/json?q=" + locationMaps + "&key=" + openCageKeyAPI)
	fmt.Println("Geographical Coordinates URL = " + urlOpenCage)

	response, err := http.Get(urlOpenCage)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &valuesGeocoding)

	// This block of code will search through the API's response to get the important information
	// Usually, the API will found multiple locations so we'll take the first one (which seems to be the most accurate) and get out of the loop
out:
	for _, v := range valuesGeocoding["results"].([]interface{}) {
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

// This function will the location name and transforming it by replacing '_' or '-' or ' ' by a '+'
// If not, the URL will not be valid and we will not get a location
func makeLocationURLValid(locationName string) string {
	replacer := strings.NewReplacer("_", "+", "-", "+", " ", "+")

	return replacer.Replace(locationName)
}
