package piscine

func IsLower(s string) bool {
	for i := range s {
		if s[i] < 97 || s[i] > 122 {
			return false
		}
	}
	return true
}
