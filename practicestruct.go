package main

import (
	"fmt"
)

type Vehicle struct {
	Name   string
	Colour string
	Model  int
}

func main22() {

	car1 := Car{Model: "BMW", Year: 2020}
	car2 := Car{Model: "Range Rover", Year: 2022}

	fmt.Println("Car 1:", car1)
	fmt.Println("Car 2:", car2)

}
