package piscine

func StringToIntSlice(str string) []int {
	res := []int{}
	if str == "" {
		return []int(nil)
	}
	for _, i := range str {
		res = append(res, int(i))
	}
	return res
}
