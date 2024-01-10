package main

import "github.com/01-edu/z01"

func main() {
	for n := '0'; n <= '9'; n++ {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
