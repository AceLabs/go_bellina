package bl

import "container/list"

type Node struct {
	Name string
	Parent *Node
	Kids *list.List
	handlers map[string] Handler
}

func NewNode(parent *Node) *Node {
	node := new(Node)

	node.Name = "-"
	node.Parent = parent
	node.Kids = list.New()

	return node
}