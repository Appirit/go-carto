package geotools

// trouvé en comparant l'analyse de Garmin
const (
	pauseMini        = Duree(60) // une pause dure au moins 1mn
	amplitudeLissage = 3         // moyenne glissante des alt.
)

const (
	lon = iota
	lat
	ele_m
	time_sec
)

// Point représente un point [lon, lat(, alt(), timestamp))].
type Position = []float64

// LineString ou MultiPoint
type LineString = []Position

// MultiLineString ou Polygon
type MultiLineString = [][]Position

type MultiPolygon = [][][]Position

type IntervalDistance = float64

type Duree = float64

// type RouteSegments []RouteSegment

// // Mesures d'un segment de trajet
// type RouteSegment struct {
// 	ID           int     // identifiant du segment
// 	Distance_m   float64 // distance en mètres
// 	Duration_sec float64 // temps en secondes
// 	Ascent       float64 // dénivelé positif
// }

// Mesures issues d'un trajet cartographique
type ElevationMetrics struct {
	// Distance      float64 // distqnce en m
	Ascent       int // dénivelé positif cumulé
	Descent      int // dénivelé négatif cumulé
	MinElevation int // altitude minimale
	MaxElevation int // altitude maximale
	MaxClimb     int // ascension maximale

	// DurationTotal float64 // durée totale
	// DurationMove  float64 // durée en mouvement
}

type GeoProperties struct {
	Date string  `mapstructure:"date"`
	Cmt  string  `mapstructure:"cmt"`
	Desc string  `mapstructure:"desc"`
	Len  float64 `mapstructure:"len"`
	Name string  `mapstructure:"name"`
	Sym  string  `mapstructure:"sym"`

	Numero        string  `mapstructure:"numero"`
	NumeroPrefixe string  `mapstructure:"numeroPrefixe"`
	Notation      string  `mapstructure:"notation"`
	Duration      float64 `mapstructure:"duration"`
	Run           float64 `mapstructure:"run"`
	Speed         float64 `mapstructure:"speed"`
	Elevation     string  `mapstructure:"elevation"`
	Trailhead     string  `mapstructure:"trailhead"`

	Operator  string `mapstructure:"operator"`
	Tel       string `mapstructure:"tel"`
	Mail      string `mapstructure:"mail"`
	Facebook  string `mapstructure:"facebook"`
	Twitter   string `mapstructure:"twitter"`
	Youtube   string `mapstructure:"youtube"`
	Instagram string `mapstructure:"instagram"`
	Pdf       string `mapstructure:"pdf"`
	Parking   string `mapstructure:"parking"`

	OtherLayers string `mapstructure:"otherLayers"`
	AmenityOvp  string `mapstructure:"amenityOvp"`
}

//	type GeoDescription struct {
//		ID         int           `mapstructure:"id"`
//		Name       string        `mapstructure:"name,omitempty"`
//		Properties GeoProperties `mapstructure:"properties,omitempty"`
//	}
type FeatureCommon struct {
	Type       string                 `json:"type"`                 // "Feature"
	Properties map[string]interface{} `json:"properties,omitempty"` // propriétés arbitraires
	ID         interface{}            `json:"id,omitempty"`         // string|number optionnel
	BBox       []float64              `json:"bbox,omitempty"`
}

// ----------------------------------------------------------------------------
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

// ----------------------------------------------------------------------------
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

// ----------------------------------------------------------------------------
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

type FeaturesCollection struct {
	Name            string
	Points          []*FeaturePoint
	LineString      []*FeatureLineString
	MultiLineString []*FeatureMultiLineString
}
