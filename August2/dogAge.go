package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserAge() (int, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the input:", err)
		return 0, fmt.Errorf("Error reading the age :%w", err)
	}

	input = strings.TrimSpace(input)

	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid age: %w", err)
	}
	return age, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name:")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the input:", err)
		return
	}

	age, err := getUserAge()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dogyears := age * 7

	fmt.Print("\n", "Hello", " ", input)
	fmt.Println("Your age in dog years is:", dogyears)
}
