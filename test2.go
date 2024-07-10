// Write a program using a for loop to print the first 10 numbers of the Fibonacci sequence.
package main

import (
	"fmt"
)

func main() {
	n := 10

	a, b := 0, 1

	fmt.Println("Fibonacci series")

	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}

}
