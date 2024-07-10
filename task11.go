//Write a program to find the factorial of a number using recursion.
/*
package main

import (
	"fmt"
)

func fact(num int) int {
	//   if num<=1{
	//	return num
	//    }

	if num == 1 || num == 0 {
		return num
	}
	return num * fact(num-1)
}

func main() {
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(50))
}

*/
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main11() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	result := fact(n)
	fmt.Println("Factorial of num", n, ":", result)
}
