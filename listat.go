package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	v := l
	for v != nil {
		v = v.Next
		count++
		if count == pos {
			return v
		}
	}
	return nil
}
