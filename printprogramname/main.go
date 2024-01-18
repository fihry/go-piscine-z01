package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ProgramName := os.Args[0]
	for _, char := range ProgramName {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
