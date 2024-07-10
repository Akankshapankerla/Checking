// Create a slice of strings and write a program to add, remove, and display elements in the slice.
package main

import (
	"fmt"
)

func addElement(slice []string, element string) []string {
	return append(slice, element)
}

func removeElement(slice []string, element string) []string {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1])
		}
	}
	return slice
}

func displaySlice(slice []string) {
	fmt.Println("SLice element")
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main() {
	slice := []string{"apple", "orange", "banana"}

	displaySlice(slice)

	slice = addElement(slice, "mango")
	slice = addElement(slice, "Avocado")
	fmt.Println("after adding elements")
	displaySlice(slice)

	slice = removeElement(slice, "banana")
	fmt.Println("After removing elements")
	displaySlice(slice)
}
