package piscine

func NRune(s string, n int) rune {
	if n >= 0 && n <= len(s) {
		res := []rune(s)
		return res[n-1]
	}
	return 0
}
