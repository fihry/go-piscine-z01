package piscine

func Swap(a *int, b *int) {
	var swp int
	swp = *a
	*a = *b
	*b = swp
}
