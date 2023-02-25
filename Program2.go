/*Implement the above in a non-recursive way. Computation should be in a separate function */

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter the Factorial Number : ")
	fmt.Scan(&number)
	fmt.Printf("Factorial %d number is %d", number, factorial_nonrecursive(number))
}

func factorial_nonrecursive(number int) int {

	if number == 0 {
		return 1
	}
	value := 1
	for i := 2; i < number+1; i++ {
		value = value * i
	}
	return value
}

// var number int
// 	fmt.Print("Enter the Factorial Number : ")
// 	fmt.Scan(&number)
// 	factorial := 1
// 	for i := 1; i <= number; i++ {
// 		factorial *= i
// 	}
// 	fmt.Printf("Factorial %d number is %d", number, factorial)
