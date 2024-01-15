package piscine

func IterativePower(nb int, power int) int {
	if power > 0 {
		res := nb
		for i := 1; power > i; i++ {
			res *= nb
		}
		return res

	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
