package piscine

func ToLower(s string) string {
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s = s[:i] + string(s[i]+32) + s[i+1:]
		}
	}
	return s
}
