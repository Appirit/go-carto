package internal

import (
	"fmt"
)

type Node struct {
	Name     string
	Level    int
	Value    []byte
	Terminal bool
}

func ParseJSON(data []byte) ([]Node, error) {
	nodes := []Node{}
	i := 0
	parseValue(data, &i, "", 0, &nodes, data, "")
	return nodes, nil
}

func parseValue(data []byte, i *int, name string, level int, nodes *[]Node, full []byte, parent string) {
	skipSpaces(data, i)
	start := *i
	terminal := false

	switch data[*i] {
	case '{':
		*i++
		end := parseObject(data, i, name, level, nodes, full)
		if name != "" {
			*nodes = append(*nodes, Node{name, level, full[start : end+1], false})
		}
	case '[':
		*i++
		if name == "coordinates" {
			// On saute entièrement le tableau
			skipArray(data, i)
			end := *i
			if name != "" {
				*nodes = append(*nodes, Node{name, level, full[start:end], false})
			}
		} else {
			// On descend dans le tableau
			parseArray(data, i, name, level, nodes, full)
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

	if terminal && name != "" {
		*nodes = append(*nodes, Node{name, level, full[start:*i], true})
	}
}

func parseObject(data []byte, i *int, parent string, level int, nodes *[]Node, full []byte) int {
	for {
		skipSpaces(data, i)
		if data[*i] == '}' {
			*i++
			return *i - 1
		}
		key := readString(data, i)
		skipSpaces(data, i)
		*i++ // skip ':'
		parseValue(data, i, key, level+1, nodes, full, parent)
		skipSpaces(data, i)
		if data[*i] == '}' {
			*i++
			return *i - 1
		}
		*i++ // skip ','
	}
}

func parseArray(data []byte, i *int, parent string, level int, nodes *[]Node, full []byte) {
	idx := 0
	for {
		skipSpaces(data, i)
		if data[*i] == ']' {
			return
		}
		name := fmt.Sprintf("%s[%d]", parent, idx)
		parseValue(data, i, name, level+1, nodes, full, parent)
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

func readString(data []byte, i *int) string {
	*i++ // skip opening "
	start := *i
	for *i < len(data) {
		if data[*i] == '\\' {
			*i += 2
			continue
		}
		if data[*i] == '"' {
			s := string(data[start:*i])
			*i++
			return s
		}
		*i++
	}
	return ""
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNumExtra(b byte) bool {
	return b == '+' || b == '-' || b == '.' || b == 'e' || b == 'E'
}
