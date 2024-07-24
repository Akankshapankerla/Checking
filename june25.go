package main

import (
	"fmt"
)

func sumRange(start, end int, ch chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
}
func mainjun25() {
	ch := make(chan int)

	ranges := [][]int{{1, 10}, {11, 20}, {21, 30}}

	for _, r := range ranges {
		go sumRange(r[0], r[1], ch)
	}

	totalSum := 0
	for i := 0; i < len(ranges); i++ {
		sum := <-ch
		fmt.Printf("Sum of range %v to %v is: %d\n", ranges[i][0], ranges[i][1], sum)
		totalSum += sum
	}
	fmt.Printf("Total sum is:%d\n", totalSum)
}
