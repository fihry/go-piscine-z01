package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	v := l.Head
	for v != nil {
		if comp(v.Data, ref) {
			return &v.Data
		}
		v = v.Next
	}
	return nil
}
