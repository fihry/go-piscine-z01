package piscine

func ListPushFront(l *List, data interface{}) {
	v := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = v
		l.Tail = v
	} else {
		l.Head.Next = v
		l.Head = v
	}
}
