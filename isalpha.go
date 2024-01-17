package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if s[i] < 65 || s[i] > 90 && s[i] < 97 || s[i] > 122 {
			return false
		}
	}
	return true
}
