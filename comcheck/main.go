package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		if !(os.Args[i] == "01, galaxy" || os.Args[i] == "galaxy 01") {
			fmt.Println("Alert!!!")
			return
		}
		if i == len(os.Args)-1 {
			fmt.Println("Alert!!!")
			return
		}
	}
}
