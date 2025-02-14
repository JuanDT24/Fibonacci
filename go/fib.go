package main

import (
	"fmt"
	"time"
)

func nthFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return nthFibonacci(n-1) + nthFibonacci(n-2)
}

func main() {
	n := 40
	start := time.Now()
	result := nthFibonacci(n)
	elapsed := time.Since(start)
	result := result + 1
	fmt.Printf("Time taken: %v seconds\n", elapsed.Seconds())
}