package geotools

import "fmt"

func (met *ElevationMetrics) ToString() string {
	return fmt.Sprintf("asc. %dm (%d - %d) max côte=%d", met.Ascent, met.MinElevation, met.MaxElevation, met.MaxClimb)
}

func MaxClimb(alt []float64) int {
	if len(alt) < 2 {
		return 0
	}
	var maxClimb, currClimb float64
	for i := 1; i < len(alt); i++ {
		diff := alt[i] - alt[i-1]
		if diff > 0 {
			currClimb += diff
			if currClimb > maxClimb {
				maxClimb = currClimb
			}
		} else {
			currClimb = 0 // descente → on repart à zéro
		}
	}
	return int(maxClimb)
}

// Calcule les éléments lié à l'altitude
func NewElevationMetrics(elevations []float64) *ElevationMetrics {
	met := ElevationMetrics{MaxClimb: MaxClimb(elevations), MaxElevation: -100000, MinElevation: 100000}

	//
	var asc, desc, previousElevation float64
	for i, ele := range elevations {
		if ele == 0 {
			// alt 0 = même que la précédente
			ele = previousElevation // correction
			continue
		}
		ele_m := int(ele)
		if met.MaxElevation < ele_m {
			met.MaxElevation = ele_m
		}
		if met.MinElevation > ele_m {
			met.MinElevation = ele_m
		}

		if i != 0 {
			dif := ele - previousElevation
			if dif > 0 {
				asc += dif
			} else if dif < 0 {
				desc += dif
			}
		}
		previousElevation = ele
	}
	met.Ascent += int(asc)
	met.Descent += int(desc)
	return &met
}

// altitude selon le type de géométrie
// func getRouteTime(feature *geojson.Feature, met *RouteMetrics) {
// 	switch feature.Geometry.Type {

// 	case geojson.GeometryLineString:
// 		elevationFromGeometry(met, feature.Geometry.LineString)

// 	case geojson.GeometryMultiLineString:
// 		if len(feature.Geometry.MultiLineString) > 0 && len(feature.Geometry.MultiLineString[0]) > 0 && len(feature.Geometry.MultiLineString[0][0]) > 0 {
// 			for _, geometry := range feature.Geometry.MultiLineString {
// 				elevationFromGeometry(met, geometry)
// 			}
// 		}
// 	}
// }
