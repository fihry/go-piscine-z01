package piscine

func SortIntegerTable(table []int) {
	s := len(table)
	for j := s - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if table[i] > table[i+1] {
				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}
}
