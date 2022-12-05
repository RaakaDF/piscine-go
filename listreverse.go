package piscine

// ListReverse function reverses the order
// of the elements of a given linked list l.
func ListReverse(l *List) {
	lnouv := &List{}
	for l.Head != nil {
		ListPushFront(lnouv, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head, l.Tail = lnouv.Head, lnouv.Tail
}
