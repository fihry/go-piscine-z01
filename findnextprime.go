package piscine

func NbPrime(nb int) bool {
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
	if NbPrime(nb) == true {
		return nb
	}
	for i := 1; nb >= i; i++ {
		if nb%i != 0 {
			res = nb
			break
		}
		nb += 1
	}
	return res
}
