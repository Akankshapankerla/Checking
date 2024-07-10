//Write a program using switch cases to print the name of the day based on an integer input (1 for Monday, 2 for Tuesday, etc.).

package main

import (
	"fmt"
)

func main() {

	var day int
	fmt.Println("Enter a number from 1-7:")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
		break
	case 2:
		fmt.Println("Tuesday")
		break
	case 3:
		fmt.Println("Wednesday")
		break
	case 4:
		fmt.Println("Thursday")
		break
	case 5:
		fmt.Println("Friday")
		break
	case 6:
		fmt.Println("Saturday")
		break
	case 7:
		fmt.Println("Sunday")
		break

	}

}
