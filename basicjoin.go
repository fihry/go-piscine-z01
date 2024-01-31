package piscine

func BasicJoin(elems []string) string {
	str := ""
	for i := 0; i < len(elems); i++ {
		for j := 0; j < len(elems[i]); j++ {
			str += string(elems[i][j])
		}
	}
	return str
}
