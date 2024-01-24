package piscine

func DescendAppendRange(max, min int) []int {
	var result []int
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result

}
