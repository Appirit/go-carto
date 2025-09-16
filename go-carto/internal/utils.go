package internal

import (
	"fmt"
	"unsafe"
)

func (root *JsonNode) ParentChain() string {
	var curName NodeName
	if root.Index >= 0 {
		curName = NodeName(fmt.Sprintf("%s[%d]", root.Name, root.Index))
	} else {
		curName = root.Name
	}
	if root.Parent != nil && root.Parent.Name != nil /*1er node*/ {
		return root.Parent.ParentChain() + "." + string(curName)
	}
	return string(curName)
}

func (root *JsonNode) isChildren(curr *JsonNode) bool {
	if curr.Parent != nil {
		if curr.Parent == root {
			return true
		} else {
			return root.isChildren(curr.Parent)
		}
	}
	return false
}

func (n *JsonNode) GetOffsets(all []byte) (int, int) {
	if n.Index < 0 {

		return int(uintptr(unsafe.Pointer(&n.Name[0])) - 1 /*le "*/ - uintptr(unsafe.Pointer(&all[0]))),
			int(uintptr(unsafe.Pointer(&n.Value[0])) + uintptr(len(n.Value)) - uintptr(unsafe.Pointer(&all[0])))
	} else {
		start := int(uintptr(unsafe.Pointer(&n.Value[0])) - uintptr(unsafe.Pointer(&all[0])))
		return start, start + len(n.Value)
	}
}
