package bl

type Handler func(*Event)

type Event struct {
	Name string
	This, Target *Node
}

func NewEvent(name string, target *Node) *Event {
	if name == "" {
		panic("Event must have a name.")
	}

	if target == nil {
		panic("Event target must not be nil.")
	}

	event := &Event{ Name: name, This: target, Target: target }

	return event
}

func (e *Event) Emit() {
	node := e.Target

	e.bubble(node)
}

func (e *Event) bubble(node *Node) {
	handler, exists := node.handlers[e.Name]

	if exists {
		e.This = node
		handler(e)
	}

	if node.Parent != nil {
		e.bubble(node.Parent)
	}
}

