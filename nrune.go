package piscine

func NRune(s string, n int) rune {
	var re rune
	if n >= 0 && n <= len(s) {
		res := []rune(s)
		re = res[n-1]
	}
	return re
}
