package main

import (
	"fmt"
)

func mainp() {
	a := [5]int{24, 4, 6, 8, 9}

	largestNumber := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > largestNumber {
			largestNumber = a[i]
		}
	}
	fmt.Println("largestnumber is:", largestNumber)
}
