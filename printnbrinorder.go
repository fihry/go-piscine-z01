package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n > 9 {
		z01.PrintRune(rune(n%10) + '0')
		PrintNbrInOrder(n / 10)
	} else {
		z01.PrintRune(rune(n%10) + '0')
	}
}
