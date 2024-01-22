package main

import "github.com/01-edu/z01"

var k = "x = 42, y = 21\n"

func main() {
	for _, i := range k {
		z01.PrintRune(i)
	}
}

// result most be :
