package piscine

func NbIsPrime(nb int) bool {
	if nb == 2 {
		return true
	} else if nb <= 1 {
		return false
	} else if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for !NbIsPrime(nb) {
		nb++
	}
	return nb
}
