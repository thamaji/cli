package cli

import "strings"

const (
	T_Command = 1 + iota
	T_Option
	T_Value
)

type Node struct {
	Type  int
	Value interface{}
}

func parse(keywords map[string]*Node, args []string) ([]*Node, []string) {
	nodes := []*Node{}

	for i, arg := range args {
		// long option
		if len(arg) >= 3 && arg[0:2] == "--" {
			name := arg
			value := ""

			if j := strings.Index(name, "="); j >= 0 {
				value = name[j+1:]
				name = name[:j]
			}

			if node, ok := keywords[name]; ok && node.Type == T_Option {
				nodes = append(nodes, node)

				if value != "" {
					nodes = append(nodes, &Node{Type: T_Value, Value: value})
				}

				continue
			}
		}

		// short option
		if len(arg) >= 2 && arg[0:1] == "-" {
			names := arg[1:]

			ok := true
			subNodes := []*Node{}

			for _, name := range names {
				if node, ok := keywords["-"+string(name)]; ok && node.Type == T_Option {
					subNodes = append(subNodes, node)
					continue
				}

				ok = false
				break
			}

			if ok {
				nodes = append(nodes, subNodes...)
				continue
			}
		}

		// sub command
		if node, ok := keywords[arg]; ok && node.Type == T_Command {
			nodes = append(nodes, node)
			return nodes, args[i:]
		}

		// argument
		nodes = append(nodes, &Node{Type: T_Value, Value: arg})
	}

	return nodes, nil
}
