package jsonparser

func (feature *RawFeature) GetGeometryType() GeometryType {
	member := findChildMember(feature.Descendants, feature.This, membername_geometry, membername_type)
	if member != nil {
		return GeometryType(BytesTrim(member.Value))
	}
	return nil
}

func (feature *RawFeature) GetId() int {
	member := findChildMember(feature.Descendants, feature.This, membername_properties, membername_number)
	if member != nil {
		return parseIdBytes(member.Value)
	}
	return -1
}
