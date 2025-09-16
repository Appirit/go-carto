package internal

type NodeName []byte

var (
	nodenameFeatures    NodeName = []byte("features")
	nodenameGeometry    NodeName = []byte("geometry")
	nodenameType        NodeName = []byte("type")
	nodenameCoordinates NodeName = []byte("coordinates")
)

type JsonNode struct {
	Name     NodeName
	Index    int // en cas de tableau
	Level    int
	Value    []byte
	Terminal bool
	Parent   *JsonNode
}

type GeometryType string

const (
	GeometryPoint           GeometryType = "Point"
	GeometryLineString      GeometryType = "LineString"
	GeometryPolygon         GeometryType = "Polygon"
	GeometryMultiPoint      GeometryType = "MultiPoint "
	GeometryMultiLineString GeometryType = "MultiLineString"
	GeometryMultiPolygon    GeometryType = "MultiPolygon"
)

type Feature struct {
	Type        GeometryType
	This        *JsonNode
	Descendants []*JsonNode
}
