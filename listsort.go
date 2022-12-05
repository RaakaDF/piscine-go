package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	head := l
	if head == nil {
		return nil
	}
	head.Next = ListSort(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}
	return head
}

func move(l *NodeI) *NodeI {
	paris := l
	nb := l.Next
	retro := nb

	for nb != nil && l.Data > nb.Data {
		paris = nb
		nb = nb.Next
	}
	paris.Next = l
	l.Next = nb
	return retro
}
