package main

import (
	"fmt"
)

func main() {
	
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	var c64 complex64 = complex(1.5, 2.5)
	var c128 complex128 = complex(1.5, 2.5)

	var b byte = 255 // Alias for uint8
	var r rune = 'a' // Alias for int32

	fmt.Println("Unsigned integers:")
	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)

	fmt.Println("\nSigned integers:")
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)

	fmt.Println("\nFloating-point numbers:")
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	fmt.Println("\nComplex numbers:")
	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	fmt.Println("\nByte and rune:")
	fmt.Println("byte:", b)
	fmt.Println("rune:", r)
}
