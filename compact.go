package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, v := range *ptr {
		if v != "" {
			count++
		}
	}
	arr := make([]string, count)
	count = 0
	for _, v := range *ptr {
		if v != "" {
			arr[count] = v
			count++
		}
	}
	*ptr = arr
	return count
}
