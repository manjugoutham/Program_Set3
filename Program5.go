/*Write a program which computes the sum and difference between two numbers. Two numbers are taken as user input. Plan using 3 functions in this program.
- sum () : it adds two numbers and returns the sum
- diff(): It computes the difference between two numbers and returns the
difference.
- getSum_n_Diff : it takes numbers and returns both sum and differences */

package main

import "fmt"

func sum(a float64, b float64) float64 {
	return a + b
}

func diff(a float64, b float64) float64 {
	return a - b
}
func getSum_n_Diff(a float64, b float64) (float64, float64) {
	return sum(a, b), diff(a, b)
}

func main() {
	var a, b, option float64
	fmt.Print("Enter the Number a: ")
	fmt.Scan(&a)
	fmt.Print("Enter the Number b: ")
	fmt.Scan(&b)
	fmt.Print("Enter the option to perform function \n 1. sum\n 2. difference\n 3. getSum_n_Diff\n")
	fmt.Scan(&option)
	switch option {
	case 1:
		fmt.Printf("the sum is %v and %v is %v", a, b, sum(a, b))
	case 2:
		fmt.Printf("The difference is %v and %v is %v", a, b, diff(a, b))
	case 3:
		// var (
		// 	s, d float64 = getSum_n_Diff(a, b)
		// )
		fmt.Printf("The sum of %v and %v is %v \n The difference of %v and %v is %v", a, b, sum(a, b), a, b, diff(a, b))
	}
}
