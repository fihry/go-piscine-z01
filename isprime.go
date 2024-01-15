package piscine

func IsPrime(nb int) bool {
	var res bool
	if nb == 2 {
		return true
	}
	if nb > 2 {
		for i := 2; i < nb; i++ {
			if nb%i != 0 {
				res = true
			} else {
				res = false
			}
		}
		return res
	} else if nb == 1 {
		res = false
	} else if nb < 1 {
		res = false
	}
	return res
}
