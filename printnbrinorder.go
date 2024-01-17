package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	digits := []int{}

	if n == 0 {
		z01.PrintRune(rune(n + '0'))
	}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	for _, digit := range digits {
		z01.PrintRune(rune(digit + '0'))
	}
}
