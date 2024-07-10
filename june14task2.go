package main

import (
	"fmt"
)

type MyError struct {
	When string
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %s, %s", e.When, e.What)
}

func main142() {
	err := &MyError{
		When: "Monday",
		What: "something went wrong",
	}
	fmt.Println(err)
}
