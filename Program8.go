package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
			fmt.Println(num)
		}
	}
	return true
}

func countNonPrimeOddNumbers(Number int) int {
	count := 0
	for i := 1; i < Number; i += 2 {
		if !isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	Number := 10
	fmt.Printf("Count non-prime odd numbers below %d is %d\n", Number, countNonPrimeOddNumbers(Number))
}
