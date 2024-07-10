package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) (string, error) {
	// Attempt to open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Read the file contents
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

func main12() {
	fileName := "D:\task\task1.go"
	content, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:", content)
}
