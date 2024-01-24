package piscine

func StringToIntSlice(str string) []int {
	res := []int{}
	for i := 0; i < len(str); i++ {
		res = append(res, int(str[i]))
	}
	return res
}
