package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ProgramName := os.Args[0]
	for index, char := range ProgramName {
		if index > 1 {
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}
