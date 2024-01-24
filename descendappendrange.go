package piscine

func DescendAppendRange(max, min int) []int {
	var result []int
	if max > min {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	}
	return result

}
