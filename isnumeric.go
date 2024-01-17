package piscine

func IsNumeric(s string) bool {
	for j := range s {
		if s[j] < '0' || s[j] > '9' {
			return false
		}
	}
	return true
}
