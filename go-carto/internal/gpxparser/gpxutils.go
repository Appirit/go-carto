package gpxparser

import (
	"fmt"
	"math"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

func timeConvert(timestamp time.Time) float64 {
	// Unix time en millisecondes
	return float64(timestamp.UnixNano() / int64(time.Second))
}

func Round2(f float64) float64 {
	return math.Round(f*100) / 100
}

func Round6(f float64) float64 {
	return math.Round(f*1000000) / 1000000
}

// Conversion d'un point GPX en Position GeoJSON
func gpxPointToPosition(p gpx.GPXPoint, referenceTime float64) Position {
	pos := Position{Round6(p.Longitude), Round6(p.Latitude)}

	if p.Elevation.NotNull() {
		pos = append(pos, Round2(p.Elevation.Value()))
	}

	if !p.Timestamp.IsZero() {
		pos = append(pos, timeConvert(p.Timestamp)-referenceTime)
	}

	return pos
}

// Conversion d'un segment GPX en LineString
func gpxSegmentToLineString(seg gpx.GPXTrackSegment) LineString {
	referenceTime := timeConvert(seg.Points[0].Timestamp)
	line := make(LineString, len(seg.Points))
	for i, pt := range seg.Points {
		line[i] = gpxPointToPosition(pt, referenceTime)
	}
	return line
}

// Conversion d'un track GPX en MultiLineString
func gpxTrackToMultiLineString(trk gpx.GPXTrack) MultiLineString {
	mls := make(MultiLineString, len(trk.Segments))
	for i, seg := range trk.Segments {
		mls[i] = gpxSegmentToLineString(seg)
	}
	return mls
}

// Crée une FeatureMultiLineString à partir d'un GPXTrack
func GpxTrackToFeature(trk gpx.GPXTrack) *FeatureMultiLineString {
	return &FeatureMultiLineString{
		FeatureCommon: FeatureCommon{
			Type:       "Feature",
			Properties: map[string]interface{}{"name": trk.Name},
		},
		Geometry: &GeometryMultiLineString{
			Type:        "MultiLineString",
			Coordinates: gpxTrackToMultiLineString(trk),
		},
	}
}

// Conversion d'un GPXPoint (waypoint) en FeaturePoint GeoJSON
func GpxWptToFeature(wpt gpx.GPXPoint, referenceTime float64) *FeaturePoint {
	props := map[string]interface{}{}
	if wpt.Symbol != "" {
		props["symbol"] = wpt.Symbol
	}
	if wpt.Comment != "" {
		props["cmt"] = wpt.Comment
	}
	if wpt.Name != "" {
		props["name"] = wpt.Name
	}

	return &FeaturePoint{
		FeatureCommon: FeatureCommon{
			Type:       "Feature",
			Properties: props,
		},
		Geometry: &GeometryPoint{
			Type:        "Point",
			Coordinates: gpxPointToPosition(wpt, referenceTime),
		},
	}
}

func ConvertToGeojson(filename string) (*FeaturesContainer, error) {

	gpxFile, err := gpx.ParseFile(filename)
	if err != nil {
		// fmt.Printf("Erreur lors de l'analyse du fichier GPX : %v\n", err)
		return nil, err
	}

	res := FeaturesContainer{}

	// Parcourir toutes les traces (tracks)
	fmt.Printf("%d Tracks\n", len(gpxFile.Tracks))
	for _, track := range gpxFile.Tracks {
		res.MultiLineString = append(res.MultiLineString, GpxTrackToFeature(track))
	}

	// Parcourir toutes les points
	referenceTime := float64(0)
	if len(gpxFile.Tracks) != 0 && len(gpxFile.Tracks[0].Segments) != 0 {

		referenceTime = timeConvert(gpxFile.Tracks[0].Segments[0].Points[0].Timestamp)
	}
	for _, wpt := range gpxFile.Waypoints {
		res.Points = append(res.Points, GpxWptToFeature(wpt, referenceTime))
	}
	return &res, nil
}
