// Write a function to reverse a string in Go.
package main

import (
	"fmt"
)

func reverse(s string) string {
	b := []rune(s)
	//A rune in Go is an alias for int32 and is used to represent a Unicode code point. This is important for handling multi-byte characters correctly.
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main7() {
	s := "Akanksha"
	reversedStr := reverse(s)
	fmt.Println("reverse of string Akanksha is:", reversedStr)
}

/*
package main

import (
	"fmt"
)

func reverseString(s string) string {
    // Convert the string to a rune slice
    runes := []rune(s)

    // Reverse the rune slice
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    // Convert the rune slice back to a string
    return string(runes)
}

func main() {
    original := "Hello"
    reversed := reverseString(original)
    fmt.Println("Original:", original)
    fmt.Println("Reversed:", reversed)
}



package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Convert the string to a slice of runes to handle multi-byte characters correctly
	runes := []rune(s)
	// Initialize two pointers, one at the beginning and one at the end of the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Swap the elements at the pointers
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert the slice of runes back to a string
	return string(runes)
}

func main() {
	str := "Hello"
	reversedStr := reverseString(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversedStr)
}
*/
