package geotools

// // charge un geojson depuis un fichier
// func Load(filePath string) (*geojson.FeatureCollection, error) {
// 	rawFeatureJSON, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	features, err := geojson.UnmarshalFeatureCollection(rawFeatureJSON)
// 	return features, err
// }

// func GetRouteMetrics(feature *geojson.Feature) (Metrics, error) {
// 	met := Metrics{minEle: 99999, maxEle: -99999}

// 	getRouteElevation(feature, &met)
// 	getRouteTime(feature, &met)
// 	return met, nil
// }

// func SimplifyNoPause(feature *geojson.Feature, speedMini_ms float64) (numPointDel int) {
// 	switch feature.Geometry.Type {

// 	case geojson.GeometryLineString:
// 		numPointDel += len(feature.Geometry.LineString) - len(simplifyNoPause(feature.Geometry.LineString, speedMini_ms))

// 	case geojson.GeometryMultiLineString:

// 		if len(feature.Geometry.MultiLineString) > 0 && len(feature.Geometry.MultiLineString[0]) > 0 && len(feature.Geometry.MultiLineString[0][0]) > 0 {
// 			for _, LineString := range feature.Geometry.MultiLineString {
// 				numPointDel += len(feature.Geometry.LineString) - len(simplifyNoPause(LineString, speedMini_ms))
// 			}
// 		}
// 	}
// 	return
// }

// func ToString(feature *geojson.Feature) string {
// 	switch feature.Geometry.Type {

// 	case geojson.GeometryMultiLineString:
// 		// fmt.Printf("%v\n", feature.Geometry.MultiLineString)
// 		var Properties GeoProperties
// 		if err := mapstructure.Decode(feature.Properties, &Properties); err != nil {
// 			log.Println(err)
// 			Properties = GeoProperties{}
// 		}
// 		return fmt.Sprintf("%s, %s-%dpts", Properties.Name, feature.Geometry.Type, len(feature.Geometry.MultiLineString[0]))
// 	}
// 	return string(feature.Geometry.Type)
// }
