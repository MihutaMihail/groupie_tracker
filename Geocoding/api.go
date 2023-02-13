package Geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetGeocodeLocation(locationMaps string) []float64 {
	openCageKeyAPI := "7fa0cecf85af4b48bb0f64bc929f7070"
	url.QueryEscape(locationMaps)
	//urlGoogle := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + locationMaps + "&key=" + googleAPIKey)
	urlOpenCage := ("https://api.opencagedata.com/geocode/v1/json?q=" + locationMaps + "&key=" + openCageKeyAPI)
	fmt.Println("Geocode URL " + urlOpenCage)

	response, err := http.Get(urlOpenCage)
	if err != nil {
		panic("Error retrieving response")
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic("Error retrieving response")
	}

	var Coordinates []float64

	var longitude float64
	var latitude float64
	var values map[string]interface{}

	json.Unmarshal(body, &values)
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

	Coordinates = append(Coordinates, latitude, longitude)
	return Coordinates
}
