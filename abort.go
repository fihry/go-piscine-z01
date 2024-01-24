package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr[2]
}
