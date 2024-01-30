package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	v := l
	for v != nil {
		if count == pos {
			return v
		}
		v = v.Next
		count++
	}
	return nil
}
