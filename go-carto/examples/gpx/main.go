package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Appirit/go-carto/internal/gpxparser"
)

func main() {
	log.Println("*************")
	log.Println("*** GPX ***")
	log.Println("*************")
	os.Chdir(must(filepath.Abs(`../../../testdata`))) // tous les outils nécessaires sont dans ce dossier

	geogson, _ := gpxparser.ConvertToGeojson("activity_19441849900.gpx")
	// geogson, _ := gpxparser.ConvertToGeojson("carroux.gpx")
	// Sérialisation JSON
	data, _ := json.MarshalIndent(geogson, "", "  ")
	fmt.Println(string(data))
}

func must[T any](val T, err error) T {
	if err != nil {
		log.Println("***********************************************************")
		log.Fatal(err)
	}
	return val
}
