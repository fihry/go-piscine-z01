package piscine

func StringToIntSlice(str string) []int {
	res := []int{}
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' {
			res = append(res, int(str[i]))
		}
	}
	return res
}
