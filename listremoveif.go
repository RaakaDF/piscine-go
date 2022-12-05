package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	nb := l.Head
	pr := nb

	for nb != nil && nb.Data == data_ref {
		l.Head = nb.Next
		nb = l.Head
	}
	for nb != nil {
		for nb != nil && nb.Data != data_ref {
			pr = nb
			nb = nb.Next
		}
		if nb == nil {
			return
		}
		pr.Next = nb.Next
		nb = n.Next
	}
}
