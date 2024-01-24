package piscine

func StringToIntSlice(str string) []int {
	res := []int{}
	if str == "" {
		return []int(nil)
	}
	for i := 0; i < len(str); i++ {
		res = append(res, int(str[i]))
	}
	return res
}
