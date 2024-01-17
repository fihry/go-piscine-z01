package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if s[i] < 'A' || s[i] > 'Z' && s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}
