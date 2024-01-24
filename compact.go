package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, v := range *ptr {
		if v != " " {
			count++
		}
	}
	return count
}
