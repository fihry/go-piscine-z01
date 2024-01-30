package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	v := l
	for v != nil {
		v = v.Next
		counter++
		if counter == pos {
			return v
		}
	}
	return nil
}
