package piscine

func IterativePower(nb int, power int) int {
	if power > 0 {
		res := nb
		for i := 1; power >= i; i++ {
			res = nb * power
		}
		return res

	} else {
		return 0
	}
}
