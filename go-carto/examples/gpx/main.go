package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Appirit/go-carto/internal/geotools"
	"github.com/Appirit/go-carto/internal/gpxparser"
)

func main() {
	log.Println("***********")
	log.Println("*** GPX ***")
	log.Println("***********")
	os.Chdir(must(filepath.Abs(`../../../testdata`))) // tous les outils nécessaires sont dans ce dossier

	geogson := must(gpxparser.NewGeojsonFromGpxfile("Comps-Forerunner-Site.gpx"))
	// geogson, _ := gpxparser.ConvertToGeojson("carroux.gpx")

	for _, multiPoints := range geogson.MultiLineString {
		for _, points := range multiPoints.Geometry.Coordinates {
			geotools.Measure(geogson.Name, points)
		}
	}

	// Sérialisation JSON
	// data := must(json.MarshalIndent(geogson, "", "  "))
	// fmt.Println(string(data))
}

func must[T any](val T, err error) T {
	if err != nil {
		log.Println("***********************************************************")
		log.Fatal(err)
	}
	return val
}
