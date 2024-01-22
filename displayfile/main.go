package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
	if len(os.Args) == 2 {
		msg, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(msg))
	}
}
