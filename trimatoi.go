package piscine

func TrimAtoi(s string) int {
	sign := 1
	res := 0
	count := 0
	for _, v := range s {
		if v == '-' && res == 0 && count == 0 {
			count++
			sign = -1
		}
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		}
	}
	return res * sign
}
