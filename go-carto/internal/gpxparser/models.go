package gpxparser

// Point représente un point [lon, lat(, alt(), timestamp))].
type Position = []float64

// LineString ou MultiPoint
type LineString = []Position

// MultiLineString ou Polygon
type MultiLineString = [][]Position

type MultiPolygon = [][][]Position

type FeatureCommon struct {
	Type       string                 `json:"type"`                 // "Feature"
	Properties map[string]interface{} `json:"properties,omitempty"` // propriétés arbitraires
	// ID         interface{}            `json:"id,omitempty"`         // string|number optionnel
	// BBox       []float64              `json:"bbox,omitempty"`
}

//----------------------------------------------------------------------------
// Geometry GeoJSON générique mais ici on s'attend à MultiLineString.
type GeometryMultiLineString struct {
	Type        string          `json:"type"`                  // "MultiLineString"
	Coordinates MultiLineString `json:"coordinates,omitempty"` // coordonnées
	// BBox        []float64       `json:"bbox,omitempty"`        // optionnel
}

// Feature GeoJSON minimal compatible avec spec.
type FeatureMultiLineString struct {
	FeatureCommon
	Geometry *GeometryMultiLineString `json:"geometry,omitempty"` // objet geometry
}

//----------------------------------------------------------------------------
// Geometry GeoJSON générique mais ici on s'attend à LineString
type GeometryLineString struct {
	Type        string     `json:"type"`                  // "MultiLineString"
	Coordinates LineString `json:"coordinates,omitempty"` // coordonnées
	// BBox        []float64  `json:"bbox,omitempty"`        // optionnel
}

// Feature GeoJSON minimal compatible avec spec.
type FeatureLineString struct {
	FeatureCommon
	Geometry *GeometryLineString `json:"geometry,omitempty"` // objet geometry
}

//----------------------------------------------------------------------------
// Geometry GeoJSON générique mais ici on s'attend à Point
type GeometryPoint struct {
	Type        string   `json:"type"`                  // "MultiLineString"
	Coordinates Position `json:"coordinates,omitempty"` // coordonnées
	// BBox        []float64 `json:"bbox,omitempty"`        // optionnel
}

// Feature GeoJSON minimal compatible avec spec.
type FeaturePoint struct {
	FeatureCommon
	Geometry *GeometryPoint `json:"geometry,omitempty"` // objet geometry
}

type FeaturesContainer struct {
	Points          []*FeaturePoint
	LineString      []*FeatureLineString
	MultiLineString []*FeatureMultiLineString
}
