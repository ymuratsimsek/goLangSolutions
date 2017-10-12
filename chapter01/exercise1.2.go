package main

import (
	"fmt"
	"os"
)

func main() {
	for index, argument := range os.Args[1:] {
		fmt.Println(index, argument)
	}
}
