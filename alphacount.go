package piscine

func AlphaCount(s string) int {
	res := 0
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			res += 1
		}
	}
	return res
}
