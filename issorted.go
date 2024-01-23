package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sort1 := true
	sort2 := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			sort1 = false
		}
		if f(a[i], a[i+1]) < 0 {
			sort2 = false
		}
	}
	return sort1 || sort2
}
