package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	n.Next = nil

	if l == nil || l.Data >= n.Data {
		n.Next = l
		return n
	}
	temp2 := l
	for temp2.Next != nil && temp2.Next.Data < n.Data {
		temp2 = temp2.Next
	}
	n.Next = temp2.Next
	temp2.Next = n

	return l
}
