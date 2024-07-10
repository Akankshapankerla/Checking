//Create a function that takes two arguments (integers) and returns their sum and product.

package main

import (
	"fmt"
)

func sumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	c, d := 3, 5

	sum, product := sumAndProduct(c, d)

	fmt.Printf("The sum of %d and %d is: %d\n", c, d, sum)
	fmt.Printf("The product of %d and %d is: %d\n", c, d, product)
}
