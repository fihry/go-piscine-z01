package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Argv := os.Args
	Argc := len(Argv)
	for i := Argc - 1; 0 < i; i-- {
		for j := 0; j < len(Argv[i]); j++ {
			z01.PrintRune(rune(Argv[i][j]))
		}
		z01.PrintRune('\n')
	}
}
