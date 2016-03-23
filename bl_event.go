package bl

type Handler func(*Event)

type Event struct {
	Name string
	Target *Node
}

func NewEvent(name string, target *Node) *Event {
	if name == "" {
		panic("Event must have a name.")
	}

	if target == nil {
		panic("Event target must not be nil.")
	}

	event := &Event{ Name: name, Target: target }

	return event
}

func (e *Event) Emit() {
	node := e.Target

	handler, exists := node.handlers[e.Name]

	if exists {
		handler(e)
	}
}

