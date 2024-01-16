package piscine

func NbIsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb%2 == 0 {
		return false
	}
	for i := 2; i*i <= nb; i += 2 {
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
	for i := 1; i <= nb; i++ {
		if NbIsPrime(nb) == true {
			return nb
		}
		nb += 1
	}
	return res
}
