package piscine

func NbIsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	var res int
	if NbIsPrime(nb) == true {
		return nb
	} else if nb <= 2 {
		return 2
	}
	for i := 1; nb >= i; i++ {
		if NbIsPrime(nb) == true {
			res = nb
			break
		}
		nb += 1
	}
	return res
}
