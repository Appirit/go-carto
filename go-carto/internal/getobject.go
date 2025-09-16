package internal

import (
	"bytes"
)

// récupère tous les item feature du tableau 'features'
func GetFeaturesNode(nodes []*JsonNode) []*Feature {
	features := []*Feature{}

	for _, featureNode := range nodes {
		if bytes.Equal(featureNode.Name, nodenameFeatures) && featureNode.Index >= 0 {
			// un nom + un index = un item de Features
			feature := Feature{This: featureNode}
			features = append(features, &feature)
			for _, curr := range nodes {
				if featureNode.isChildren((curr)) {
					feature.Descendants = append(feature.Descendants, curr)
				}
			}
		}
	}
	return features
}

func (feature *Feature) GetGeometryType() GeometryType {
	for _, n := range feature.Descendants {
		if n.Level == feature.This.Level+2 && bytes.Equal(n.Name, nodenameType) && bytes.Equal(n.Parent.Name, nodenameGeometry) {
			return GeometryType(string(n.Value[1 : len(n.Value)-2]))
		}
	}
	return ""
}
