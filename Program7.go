/*
For example if n is 10, return value will be 23 Because all the natural numbers below 10 that are multiples of
3 or 5, are 3, 5, 6 and 9. And the sum of 3, 5, 6, 9 is 23
*/
package main

import "fmt"

func main() {

	var number, i int

	fmt.Print("Enter the Natural Number = ")
	fmt.Scan(&number)
	total := 0
	for i = 1; i <= number; i++ {
		total = total + i
		fmt.Println(i)
	}

	fmt.Println("The Sum of Natural Numbers from 1 to ", total)
}
