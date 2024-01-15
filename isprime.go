package piscine

func IsPrime(nb int) bool {
	var res bool
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i != 0 {
				res = true
			} else {
				return false
			}
		}
		return res
	} else if nb < 2 {
		return false
	}
	return res
}
