package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			return false
		} else if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return true
}
