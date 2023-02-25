/*Write a program which computes the factorial of a number. Number is taken as
input. Factorial implementation should be in a recursive way and should be in a
function.
Sample input: 3
Desired Output : 6*/

package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number : ")
	fmt.Scan(&number)
	fmt.Printf("factorial number %d is %d\n", number, factorial(number))
	// factorial(number)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
