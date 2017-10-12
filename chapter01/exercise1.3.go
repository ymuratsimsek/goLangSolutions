package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func withLoop() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func withLoopWithRange() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func withJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	fmt.Println("Execution time withLoop:")
	withLoop()
	fmt.Println("Execution time withLoopWithArgs:")
	withLoopWithRange()
	fmt.Println("Executioon time withJoin:")
	withJoin()
}
