package main

import (
	"fmt"
)

func main() {
	n := 10
	dynamic := true

	if dynamic {
		memo = make([]float64, n)
		fmt.Println(fib2(n))
	} else {
		fmt.Println(fib1(n))
	}
}

func fib1(n int) float64 {
	if n <= 2 {
		return 1
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}

var memo []float64

func fib2(n int) float64 {
	var f float64
	if memo[n-1] != 0 {
		return memo[n-1]
	} else if n <= 2 {
		f = 1
	} else {
		f = fib2(n-1) + fib2(n-2)
		memo[n-1] = f
	}

	return f
}
