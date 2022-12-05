package piscine

// ListAt function takes a pointer to the first element of list l and an int pos as parameters.
// The function prints the NodeL in the position pos of the linked list l.
func ListAt(l *NodeL, pos int) *NodeL {
	y := 0
	for l != nil {
		if pos == y {
			return l
		}
		y++
		l = l.Next
	}
	return nil
}
