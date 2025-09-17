package jsonparser

import (
	"bytes"
	"fmt"
	"unsafe"
)

func BytesTrim(b []byte) []byte {
	// insensible aux ""
	if b[0] == '"' && len(b) >= 2 {
		b = b[1 : len(b)-1]
	}
	return b
}

func parseIdBytes(b []byte) (id int) {
	// insensible aux ""
	b = BytesTrim(b)

	for i := 0; i < len(b); i++ {
		c := b[i]
		if c < '0' || c > '9' {
			return -1
		}
		id = id*10 + int(c-'0')
	}

	return id
}

// trouve le membre enfant d'un membre parent
func findChildMember(members []*JsonMember, curr *JsonMember, parentnames ...[]byte) *JsonMember {
outer: // label sur la boucle externe
	for _, m := range members {
		if m.Level == curr.Level+len(parentnames) {
			// parcourir les noms des parents attendus à l'envers
			iter := m
			for i := len(parentnames) - 1; i >= 0; i-- {
				if !bytes.Equal(iter.Name, parentnames[i]) {
					continue outer // pas bon, saute au prochain member
				}
				iter = iter.Parent
			}
			if iter == curr { // c'est le cas sauf si 'members' contient plus que la filiation de 'curr'
				// trouvé
				return m
			}
		}
	}
	return nil
}

// récupère tous les item feature du tableau 'features'
func GetFeaturesNode(members []*JsonMember) []*RawFeature {
	features := []*RawFeature{}

	for _, featureNode := range members {
		if bytes.Equal(featureNode.Name, membername_features) && featureNode.Index >= 0 {
			// un nom + un index = un item de Features
			feature := RawFeature{This: featureNode}
			features = append(features, &feature)
			for _, curr := range members {
				if featureNode.isChildren((curr)) {
					feature.Descendants = append(feature.Descendants, curr)
				}
			}
		}
	}
	return features
}

func (root *JsonMember) ParentChain() string {
	var curName MemberName
	if root.Index >= 0 {
		curName = MemberName(fmt.Sprintf("%s[%d]", root.Name, root.Index))
	} else {
		curName = root.Name
	}
	if root.Parent != nil && root.Parent.Name != nil /*1er node*/ {
		return root.Parent.ParentChain() + "." + string(curName)
	}
	return string(curName)
}

func (root *JsonMember) isChildren(curr *JsonMember) bool {
	if curr.Parent != nil {
		if curr.Parent == root {
			return true
		} else {
			return root.isChildren(curr.Parent)
		}
	}
	return false
}

func (n *JsonMember) GetOffsets(all []byte) (int, int) {
	if n.Index < 0 {

		return int(uintptr(unsafe.Pointer(&n.Name[0])) - 1 /*le "*/ - uintptr(unsafe.Pointer(&all[0]))),
			int(uintptr(unsafe.Pointer(&n.Value[0])) + uintptr(len(n.Value)) - uintptr(unsafe.Pointer(&all[0])))
	} else {
		start := int(uintptr(unsafe.Pointer(&n.Value[0])) - uintptr(unsafe.Pointer(&all[0])))
		return start, start + len(n.Value)
	}
}
