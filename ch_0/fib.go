package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var maxItrs, err = strconv.ParseInt(os.Args[len(os.Args)-1], 0, 0)
	if err != nil {
		panic(err)
	}
	var fib1, fib2 int64 = 0, 1
	fmt.Println(fib1, fib2)
	for i := 1; int64(i) < maxItrs; i++ {
		fib1, fib2 = fib2, fib1+fib2
		fmt.Println(fib2)
	}
}
