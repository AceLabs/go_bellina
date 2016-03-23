package bl

import "container/list"

type Node struct {
	Name string
	Kids *list.List
	handlers map[string] Handler
}

func NewNode() *Node {
	node := new(Node)

	node.Name = "-"

	node.Kids = list.New()

	return node
}