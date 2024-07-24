//Write a function that performs division of two numbers and returns an error if the divisor is zero.
//Handle the error gracefully in the main function.

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func main() {

	numbers := [][2]float64{
		{10, 2},
		{10, 0},
		{20, 5},
	}

	for _, pair := range numbers {
		a, b := pair[0], pair[1]
		result, err := divide(a, b)
		if err != nil {
			fmt.Printf("Error dividing %.2f by %.2f: %v\n", a, b, err)
		} else {
			fmt.Printf("%.2f divided by %.2f is %.2f\n", a, b, result)
		}
	}

}
