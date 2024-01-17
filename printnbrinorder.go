package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	digits := []int{}

	if n == 0 {
		z01.PrintRune(rune(n + '0'))
	}
	if n < 0 {
		z01.PrintRune(rune('-'))
		n = -n
	}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	SortNb(digits)
	for _, v := range digits {
		z01.PrintRune(rune(v + '0'))
	}
}

func SortNb(digits []int) {
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
}
