package gpxparser

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
	"time"

	"github.com/Appirit/go-carto/internal/geotools"
	"github.com/tkrajina/gpxgo/gpx"
)

func timeConvert(timestamp time.Time) float64 {
	// Unix time en millisecondes
	return float64(timestamp.UnixNano() / int64(time.Second))
}

func Round2(f float64) float64 {
	return math.Round(f*100) / 100
}

// pour les coordonnées : 6 décimales → ~0,11 m (11 cm)
func Round6(f float64) float64 {
	return math.Round(f*1000000) / 1000000
}

// Conversion d'un point GPX en Position GeoJSON
func gpxPointToPosition(p gpx.GPXPoint, referenceTime float64) geotools.Position {
	pos := geotools.Position{Round6(p.Longitude), Round6(p.Latitude)}

	hasElevation := p.Elevation.NotNull()
	if hasElevation {
		pos = append(pos, Round2(p.Elevation.Value()))
	}

	if !p.Timestamp.IsZero() {
		if !hasElevation {
			pos = append(pos, 0)
		}
		pos = append(pos, timeConvert(p.Timestamp)-referenceTime)
	}

	return pos
}

// Conversion d'un segment GPX en geotools.LineString
func gpxSegmentToLineString(seg gpx.GPXTrackSegment) geotools.LineString {
	referenceTime := timeConvert(seg.Points[0].Timestamp)
	line := make(geotools.LineString, len(seg.Points))
	for i, pt := range seg.Points {
		line[i] = gpxPointToPosition(pt, referenceTime)
	}
	return line
}

// Conversion d'un track GPX en geotools.MultiLineString
func gpxTrackToMultiLineString(trk gpx.GPXTrack) geotools.MultiLineString {
	mls := make(geotools.MultiLineString, len(trk.Segments))
	for i, seg := range trk.Segments {
		mls[i] = gpxSegmentToLineString(seg)
	}
	return mls
}

// Crée une geotools.FeatureMultiLineString à partir d'un GPXTrack
func GpxTrackToFeature(trk gpx.GPXTrack) *geotools.FeatureMultiLineString {
	return &geotools.FeatureMultiLineString{
		FeatureCommon: geotools.FeatureCommon{
			Type:       "Feature",
			Properties: map[string]interface{}{"name": trk.Name},
		},
		Geometry: &geotools.GeometryMultiLineString{
			Type:        "geotools.MultiLineString",
			Coordinates: gpxTrackToMultiLineString(trk),
		},
	}
}

// Conversion d'un GPXPoint (waypoint) en geotools.FeaturePoint GeoJSON
func GpxWptToFeature(wpt gpx.GPXPoint, referenceTime float64) *geotools.FeaturePoint {
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

	return &geotools.FeaturePoint{
		FeatureCommon: geotools.FeatureCommon{
			Type:       "Feature",
			Properties: props,
		},
		Geometry: &geotools.GeometryPoint{
			Type:        "Point",
			Coordinates: gpxPointToPosition(wpt, referenceTime),
		},
	}
}

func NewGeojsonFromGpxfile(filename string) (*geotools.FeaturesCollection, error) {

	gpxFile, err := gpx.ParseFile(filename)
	if err != nil {
		// fmt.Printf("Erreur lors de l'analyse du fichier GPX : %v\n", err)
		return nil, err
	}

	res := geotools.FeaturesCollection{Name: strings.TrimSuffix(filename, filepath.Ext(filename))}

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
