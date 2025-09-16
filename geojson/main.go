package main

import (
	"fmt"
	"os"

	geojson "github.com/paulmach/go.geojson"
)

func main() {
	fmt.Println("Hello, World!")
	// Feature Collection

	rawFeatureJSON, _ := os.ReadFile("D0576.json")
	// rawFeatureJSON := []byte(`
	// { "type": "FeatureCollection",
	//   "features": [
	//     { "type": "Feature",
	//       "geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
	//       "properties": {"prop0": "value0"}
	//     }
	//   ]
	// }`)

	fc1, _ := geojson.UnmarshalFeatureCollection(rawFeatureJSON)
	fmt.Printf("%v\n", fc1)
	for _, feat := range fc1.Features {
		fmt.Printf("%v\n", feat)

	}
}
