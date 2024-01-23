package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		msg, _ := ioutil.ReadAll(os.Stdin)
		for _, ch := range string(msg) {
			z01.PrintRune(ch)
		}
	} else {
		for i := 1; i < len(os.Args); i++ {
			msg, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				for _, ch := range "ERROR: " + err.Error() {
					z01.PrintRune(ch)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			for _, ch := range string(msg) {
				z01.PrintRune(ch)
			}
		}
	}
}
