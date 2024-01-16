package main

import (
	"fmt"
	// "piscine"
)

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("", ""))
}

func Index(s string, toFind string) int {
	if len(toFind) != 0 {
		for i := 0; i <= len(s)-1; i++ {
			if s[i] == toFind[0] {
				return i
			}
		}
		return -1
	}
	return 0
}
