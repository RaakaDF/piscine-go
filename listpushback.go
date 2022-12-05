package piscine

// NodeL type
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List type
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushBack function inserts a new element NodeL
// at the end of the list l while using the structure List.
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}
