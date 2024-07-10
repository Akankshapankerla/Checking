//Create variables of different types (string, int, float64, bool).
//Perform basic operations and print the results.

package main

import (
	"fmt"
)

func main2() {

	var a int = 12
	b := 14
	var x int = 10
	//increment
	x++
	//decrement
	//x--
	//var c float64 = 3.14
	d := "Akanksha"
	s := "pankerla"
	var e bool = true
	var f bool

	fmt.Println("multiplication of a ,b:", a*b)
	fmt.Println("addition of a,b:", a+b)
	fmt.Println(a % b)
	fmt.Println(a / b)
	fmt.Println("increment of x:", x)
	fmt.Println("full name:", d+s) //appending
	fmt.Println("value of e:", e)
	fmt.Println("value of f:", f)

	var str1 string = "Hello"
	var str2 string = "World"
	var integer1 int = 42
	var integer2 int = 100
	var floatNum1 float64 = 3.14
	var floatNum2 float64 = 2.71
	var boolean1 bool = true
	var boolean2 bool = false

	// Performing comparison operations
	strComparison := str1 == str2
	intComparison := integer1 > integer2
	floatComparison := floatNum1 <= floatNum2
	boolComparison := boolean1 == boolean2

	// Printing the results
	fmt.Println("String comparison (str1 == str2):", strComparison)
	fmt.Println("Integer comparison (integer1 > integer2):", intComparison)
	fmt.Println("Float comparison (floatNum1 <= floatNum2):", floatComparison)
	fmt.Println("Boolean comparison (boolean1 == boolean2):", boolComparison)
}

/*
package main

import (
	"fmt"
)

func main4() {
	a := 12
	b := 3
	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println("a>b=", a > b)
	fmt.Println("a<b=", a < b)

	c := 4
	d := 2
	e := 4
	var f bool = true
	var g bool
	if c == d {
		fmt.Println("equal:", c+d)
	} else {
		fmt.Println("not equal:", c-d)
	}
	if c == d && d == e && e == c {
		fmt.Println(f)
	} else if c == d || d == e || e == c {
		fmt.Println("hello")
	} else {
		fmt.Println("conditions:", g)
	}

}

*/
