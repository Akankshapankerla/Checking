package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main6() {
	go printMessage("Hello from goroutine")
	printMessage("Hello from main")
}
