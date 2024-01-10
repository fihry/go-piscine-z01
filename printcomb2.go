package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := 1; b <= 99; b++ {
			last := !(a == 98 && b == 99)
			z01.PrintRune(rune('0' + a/10))
			z01.PrintRune(rune('0' + a%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + b/10))
			z01.PrintRune(rune('0' + b%10))
			if last {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
