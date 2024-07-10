package main

import (
	"fmt"
	"strings"
)

func main5() {
	text := "The world is a source of immense beauty and joy, providing us with a healthy environment in which to live"
	wordFreq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		wordFreq[word]++
	}

	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
