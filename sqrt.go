package piscine

func Sqrt(nb int) int {
	if nb > 1 {
		var v int
		for i := 1; i <= nb; i++ {
			if i*i == nb {
				v = i
			}
		}
		return v
	} else if nb == 1 {
		return 1
	} else {
		return 0
	}
}
