package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) information() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

func main10() {
	person := person{Name: "Alice", Age: 30}

	person.information()
}
