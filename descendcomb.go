package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	var st int = 0
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := i; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if i == k && (j < l || j == l) {
						continue
					}
					if st > 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					st = 1
				}
			}
		}
	}
}
