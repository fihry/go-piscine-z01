package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argv := os.Args
	argc := len(argv)
	for i := 1; i < argc; i++ {
		for j := i + 1; j < argc; j++ {
			if argv[i] > argv[j] {
				argv[i], argv[j] = argv[j], argv[i]
			}
		}
	}
	for i := 1; i < argc; i++ {
		for j := 0; j < len(argv[i]); j++ {
			z01.PrintRune(rune(argv[i][j]))
		}
		z01.PrintRune('\n')
	}
}
