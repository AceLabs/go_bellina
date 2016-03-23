package bl

import (
	adt "adt/clown_stack"
	"fmt"
)

var root *Node
var current *Node
var nodeStack *adt.Stack

func init() {
	Clear()
}

func Clear() {
	root = NewNode()
	root.Name = "root"

	nodeStack = adt.NewStack()
}

func Root() {
	current = root
	nodeStack.Push(current)
}

func Div() {
	node := NewNode()

	current.Kids.PushBack(node)

	current = node
	nodeStack.Push(current)
}

func End() {
	nodeStack.Pop()

	if nodeStack.Len() > 0 {
		current = nodeStack.Peek().(*Node)
	} else {
		current = nil
	}
}

func GetCurrent() *Node {
	return current
}

func On(eventName string, handler Handler) {
	if current.handlers == nil {
		current.handlers = make(map[string] Handler)
	}

	if eventName == "" {
		panic("On cannot have empty eventName.")
	}

	if handler == nil {
		panic("On cannot have nil handler.")
	}

	_, exists := current.handlers[eventName]

	if exists {
		panic(fmt.Sprintf("On '%s' has already been registered for this node.", eventName))
	}

	current.handlers[eventName] = handler
}


