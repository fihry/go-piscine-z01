package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Argv := os.Args
	Argc := len(Argv)
	for i := 1; i < Argc; i++ {
		for j := 0; j < len(Argv[i]); j++ {
			z01.PrintRune(rune(Argv[i][j]))
		}
		z01.PrintRune('\n')
	}
}

// argument := os.Args[1:]

// for _, oneArg := range argument {
// 	for _, char := range oneArg {
// 		z01.PrintRune(char)
// 	}
// 	z01.PrintRune('\n')
// }
