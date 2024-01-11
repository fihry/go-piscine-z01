package piscine

func UltimateDivMod(a *int, b *int) {
	var swp int
	var swp1 int
	swp = *a

	swp1 = *b
	*a = swp / swp1
	*b = swp % swp1
}
