package piscine

// ListSize function returns the number
// of elements in a linked list l.
func ListSize(l *List) int {
	i := 0
	for l.Head != nil {
		i++
		l.Head = l.Head.Next
	}
	return i
}
