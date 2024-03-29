package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	v := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = v
		l.Tail = v
	} else {
		l.Tail.Next = v
		l.Tail = v
	}
}
