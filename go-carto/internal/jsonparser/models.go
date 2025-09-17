package jsonparser

type MemberName []byte

var (
	membername_features    MemberName = []byte("features")
	membername_geometry    MemberName = []byte("geometry")
	membername_type        MemberName = []byte("type")
	membername_coordinates MemberName = []byte("coordinates")
	membername_properties  MemberName = []byte("properties")
	membername_number      MemberName = []byte("number")
)

type GeometryType []byte

var (
	GeometryPoint           GeometryType = []byte("Point")
	GeometryLineString      GeometryType = []byte("LineString")
	GeometryPolygon         GeometryType = []byte("Polygon")
	GeometryMultiPoint      GeometryType = []byte("MultiPoint ")
	GeometryMultiLineString GeometryType = []byte("MultiLineString")
	GeometryMultiPolygon    GeometryType = []byte("MultiPolygon")
)

// membre JSON (name/value pair) ou item de tableau après parse
type JsonMember struct {
	Name     MemberName  // nom du membre
	Index    int         // -1 ou index dans le tableau si ce membre est un item de tableau
	Level    int         // profondeur du membre
	Value    []byte      // valeur =tout si élément de tableau
	Terminal bool        // est ce une feuille de l'arborescence
	Parent   *JsonMember // lien arrièrevvers le parent
}

type RawFeature struct {
	Type        GeometryType
	This        *JsonMember
	Descendants []*JsonMember
}
