package piscine

func ListLast(l *List) interface{} {
	current := l.Head
	if current == nil {
		return nil
	}
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
