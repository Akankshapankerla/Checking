package main

import (
	"fmt"
)

func main1() {
	//Write a program to find the largest number in an array.
	a := [6]int{4, 8, 6, 10, 12, 345}

	largestNumber := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > largestNumber {
			largestNumber = a[i]
		}
	}

	fmt.Println("largest number in the given array is:", largestNumber)

}
