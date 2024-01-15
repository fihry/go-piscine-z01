package piscine

func IsPrime(nb int) bool {
	res := true
	if nb <= 1 {
		res = false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			res = false
		}
	}
	return res
}
