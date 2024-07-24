package main

import "fmt"

func main() {
	var x int = 5748

	var p *int

	p = &x

	fmt.Println("value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
}
