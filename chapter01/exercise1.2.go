package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for index, argument := range os.Args[1:] {
		fmt.Println(index)
		fmt.Println(argument)
		s += sep + argument
		sep = " "
	}
	fmt.Println(s)
}
