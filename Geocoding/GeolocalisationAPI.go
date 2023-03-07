package Geocoding

import (
	"fmt"
	"strconv"
)

// Declaration of variables
var (
	geoapifyKeyAPI = "37fcff924dd44dee820c00a295f16864" // Mihail Key
)

// This function uses the coordinates and creates an image that indicates where the location is placed on the map
func GetGeolocalisationMap(coordinates []float64) string {
	latitude := strconv.FormatFloat(coordinates[0], 'f', -1, 64)
	longitude := strconv.FormatFloat(coordinates[1], 'f', -1, 64)

	urlMapMarker := ("https://maps.geoapify.com/v1/staticmap?style=osm-carto&width=800&height=600&" +
		"center=lonlat:" + longitude + "," + latitude + "&zoom=13.8135&" +
		"marker=lonlat:" + longitude + "," + latitude + ";color:red;" +
		"size:medium&apiKey=" + geoapifyKeyAPI)

	fmt.Println("Map Marker URL = " + urlMapMarker)

	return urlMapMarker
}
