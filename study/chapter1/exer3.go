package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	echo1()
	fmt.Println(time.Since(start))
	start = time.Now()
	echo2()
	fmt.Println(time.Since(start))
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
