package piscine

func IsUpper(s string) bool {
	for i := range s {
		if s[i] < 65 || s[i] > 90 {
			return false
		}
	}
	return true
}
