//Write programs using if-else statements, switch cases, and for loops.
//Example: Write a program that checks if a number is even or odd.

package main

import "fmt"

func main3() {
	a := 18
	//using if-else statements
	if a%2 == 0 {
		fmt.Println(a, "is an Even number")
	} else {
		fmt.Println("Not an Even Num")
	}

	//using switch-case
	switch a % 2 {
	case 0:
		fmt.Println(a, "is even")
	default:
		fmt.Println(a, "is odd")
	}

	//using for loop

	var c int
	var d = 27
	c = d % 2

	for i := c; i == 1; i++ {
		fmt.Println("d is odd")
	}
	for i := c; i == 0; i++ {
		fmt.Println("d is even")
	}

}

//using for loop
/*package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)
    //for i:=10;i<=num;i++{
	for i := num; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
*/
