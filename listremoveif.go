package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	pr := n

	for n != nil && n.Data == data_ref {
		l.Head = n.Next
		n = l.Head
	}
	for n != nil {
		for n != nil && n.Data != data_ref {
			pr = n
			n = n.Next
		}
		if n == nil {
			return
		}
		pr.Next = n.Next
		n = n.Next
	}
}
