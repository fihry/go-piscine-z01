package piscine

func ToUpper(s string) string {
	for i := range s {
		if s[i] >= 97 && s[i] <= 122 {
			s = s[:i] + string(s[i]-32) + s[i+1:]
		}
	}

	return s
}
