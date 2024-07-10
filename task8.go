package main

import (
	"fmt"
)

func reverse1(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main8() {

	str := "Geeks"
	strReverse := reverse1(str)
	fmt.Println(str)
	fmt.Println(strReverse)
}

/*
func reverseString(s string) string {
	character := []rune(s)
	for i, j := 0, len(character)-1; i < j; i, j = i+1, j-1 {
		character[i], character[j] = character[j], character[i]
	}
	return string(character)
}

func main8() {
	str := "Hello world!"
	reversedStr := reverseString(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversedStr)
}*/

//

// function, which takes a string as
// argument and return the reverse of string.
