package piscine

func Atoi(s string) int {
	sign := 1
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				digit := int(s[i] - '0')
				result = result*10 + digit
			} else {
				return 0
			}
		}
	}
	return result * sign
}
