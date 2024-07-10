package main

import (
	"fmt"
	"os"
)

func fileread(CSV string) (string, error) {

	file, err := os.Open(CSV)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	content, err := os.ReadFile(CSV)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

func main() {
	CSV := "D:\task\task7.go"
	content, err := readFile(CSV)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:", content)
}
