package piscine

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}
	return res
}
