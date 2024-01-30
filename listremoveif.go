package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	v := l.Head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}
	for v != nil {
		if v.Next != nil && v.Next.Data == data_ref {
			if v.Next == l.Tail {
				l.Tail = v
			}
			v.Next = v.Next.Next
		} else {
			v = v.Next
		}
	}
}
