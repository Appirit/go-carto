package internal

import "bytes"

func ParseJSON(data []byte) ([]*JsonNode, error) {
	nodes := []*JsonNode{}
	i := 0
	parseValue(data, &i, nil, -1, 0, nil, &nodes, data)
	return nodes, nil
}

func parseValue(data []byte, i *int, name NodeName, idx int, level int, parent *JsonNode, nodes *[]*JsonNode, full []byte) {
	skipSpaces(data, i)
	start := *i
	terminal := false
	// log.Printf("current=%s i=%d, lvl=%d, %v\n", name, *i, level, parent)
	switch data[*i] {
	case '{':
		*i++

		node := &JsonNode{name, idx, level, nil, false, parent}
		parent = node
		end := parseObject(data, i, level, nodes, full, parent)
		if name != nil {
			node.Value = full[start : end+1]
			*nodes = append(*nodes, node)
		}
	case '[':
		*i++
		if bytes.Equal(name, nodenameCoordinates) {
			// On saute entièrement le tableau
			skipArray(data, i)
			end := *i
			if name != nil {
				node := &JsonNode{name, idx, level, full[start:end], false, parent}
				parent = node
				*nodes = append(*nodes, node)
			}
		} else {
			// On descend dans le tableau
			parseArray(data, i, name, level, nodes, full, parent)
			skipSpaces(data, i)
			*i++ // skip ']'
		}
	case '"':
		parseString(data, i)
		terminal = true
	case 't': // true
		*i += 4
		terminal = true
	case 'f': // false
		*i += 5
		terminal = true
	case 'n': // null
		*i += 4
		terminal = true
	default: // number
		for *i < len(data) && (isDigit(data[*i]) || isNumExtra(data[*i])) {
			*i++
		}
		terminal = true
	}

	if terminal && name != nil {
		node := &JsonNode{name, idx, level, full[start:*i], true, parent}
		parent = node
		*nodes = append(*nodes, node)
	}
}

// parse un objet depuis un '{'
func parseObject(data []byte, i *int, level int, nodes *[]*JsonNode, full []byte, parent *JsonNode) int {
	for {
		skipSpaces(data, i)
		if data[*i] == '}' {
			*i++
			return *i - 1
		}
		key := readNodeName(data, i)
		skipSpaces(data, i)
		*i++ // skip ':'
		parseValue(data, i, key, -1, level+1, parent, nodes, full)
		skipSpaces(data, i)
		if data[*i] == '}' {
			*i++
			return *i - 1
		}
		*i++ // skip ','
	}
}

func parseArray(data []byte, i *int, parentName NodeName, level int, nodes *[]*JsonNode, full []byte, parent *JsonNode) {
	idx := 0
	for {
		skipSpaces(data, i)
		if data[*i] == ']' {
			return
		}
		// name := fmt.Sprintf("%s[%d]", parentName, idx)
		parseValue(data, i, parentName, idx, level+1, parent, nodes, full)
		idx++
		skipSpaces(data, i)
		if data[*i] == ']' {
			return
		}
		*i++ // skip ','
	}
}

func skipArray(data []byte, i *int) {
	level := 1
	for *i < len(data) && level > 0 {
		switch data[*i] {
		case '[':
			level++
		case ']':
			level--
		case '"':
			parseString(data, i)
			continue
		}
		*i++
	}
}

func skipSpaces(data []byte, i *int) {
	for *i < len(data) {
		switch data[*i] {
		case ' ', '\n', '\r', '\t':
			*i++
		default:
			return
		}
	}
}

func parseString(data []byte, i *int) {
	*i++ // skip opening "
	for *i < len(data) {
		if data[*i] == '\\' {
			*i += 2
			continue
		}
		if data[*i] == '"' {
			*i++
			break
		}
		*i++
	}
}

func readNodeName(data []byte, i *int) NodeName {
	*i++ // skip opening "
	start := *i
	for *i < len(data) {
		if data[*i] == '\\' {
			*i += 2
			continue
		}
		if data[*i] == '"' {
			s := data[start:*i]
			*i++
			return s
		}
		*i++
	}
	return nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNumExtra(b byte) bool {
	return b == '+' || b == '-' || b == '.' || b == 'e' || b == 'E'
}
