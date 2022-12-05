package piscine

// ListLast function returns the last element of a linked list l.
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	x := l.Head
	for x.Next != nil {
		x = x.Next
	}
	return x.Data
}