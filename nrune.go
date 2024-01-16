package piscine

func NRune(s string, n int) rune {
	var r rune
	if n >= 0 && n <= len(s) {
		res := []rune(s)
		r = res[n-1]
	}
	return r
}
