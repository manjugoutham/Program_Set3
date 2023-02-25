package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 1 {
		return []int{0, 1}
	}

	s := fibonacci(n - 1)
	return append(s, s[n-1]+s[n-2])
}

func main() {
	n := 5
	series := fibonacci(n)

	fmt.Printf("Fibonacci series of size %d: %v\n", n, series)
}
