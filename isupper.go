package piscine

func IsUpper(s string) bool {
	var res bool
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			res = true
		} else {
			res = false
		}
	}
	return res
}
