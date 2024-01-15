package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 && nb < 64 {
		res := nb
		for i := 1; nb > i; i++ {
			res = res * (nb - i)
		}
		return res
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
