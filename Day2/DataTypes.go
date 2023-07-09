package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Boolean
	var b1 bool
	size := unsafe.Sizeof(b1)
	var b2 = true
	b3 := false
	fmt.Printf("Type of variable: %T\n", b1)
	fmt.Printf("Size of bool: %d\n", size)
	fmt.Println("b1 = ", b1)
	fmt.Println("b2 = ", b2)
	fmt.Println("b3 = ", b3)

	/*Integers
	        - Signed integers
			- Unsigned integers
	*/

	// Signed integers: int, int8 int16, int32, int64
	var i1 int = -9223372036854775807
	fmt.Printf("Size of int: %d Bytes\n", unsafe.Sizeof(i1))
	fmt.Println("Value of int: ", i1)
	var i2 int8 = -127
	fmt.Printf("Size of int8: %d Bytes\n", unsafe.Sizeof(i2))
	fmt.Println("Value of int8: ", i2)
	var i3 int16 = -32767
	fmt.Printf("Size of int16: %d Bytes\n", unsafe.Sizeof(i3))
	fmt.Println("Value of int16: ", i3)
	var i4 int32 = -2147483647
	fmt.Printf("Size of int32: %d Bytes\n", unsafe.Sizeof(i4))
	fmt.Println("Value of int32: ", i4)
	var i5 int64 = -9223372036854775807
	fmt.Printf("Size of int64: %d Bytes\n", unsafe.Sizeof(i5))
	fmt.Println("Value of int64: ", i5)

	// Unsigned integers: uint, uint8, uint16, uint32, uint64
	var ui1 uint = 18446744073709551615 
	fmt.Printf("Size of uint: %d Bytes\n", unsafe.Sizeof(ui1))
	fmt.Println("Value of uint: ", ui1)
	var ui2 uint8 = 255
	fmt.Printf("Size of uint8: %d Bytes\n", unsafe.Sizeof(ui2))
	fmt.Println("Value of uint8: ", ui2)
	var ui3 uint16 = 65535
	fmt.Printf("Size of uint16: %d Bytes\n", unsafe.Sizeof(ui3))
	fmt.Println("Value of uint16: ", ui3)
	var ui4 uint32 = 4294967295
	fmt.Printf("Size of uint32: %d Bytes\n", unsafe.Sizeof(ui4))
	fmt.Println("Value of uint32: ", ui4)
	var ui5 uint64 = 18446744073709551615
	fmt.Printf("Size of uint64: %d Bytes\n", unsafe.Sizeof(ui5))
	fmt.Println("Value of uint64: ", ui5)

	// Byte
	var bt1 byte = 'A'
	var bt2 []byte = []byte("ABC")
	fmt.Printf("Type of variable: %T\n", bt1)
	fmt.Printf("%c = %d\n", bt1, bt1)
	fmt.Println(bt2)

	// Rune
	var r1 rune = '€'
	fmt.Printf("Type of variable: %T\n", r1)
	fmt.Printf("%c = %U and %c = %d\n", r1, r1, r1, r1)

	// Float: float32 and float64
	var f1 float64 = 1.11111111111111122222
	var f2 float32 = 1.1234567
	var f3 float32
	fmt.Printf("Type of variable: %T\n", f1)
	fmt.Println(f1, f2, f3)

	// Complex: complex64 and complex128
	var c1 complex64 = 2.1234567 + 3.1234567i
	var c2 complex128 = 1.12345678912345678 + 2.12345678912345678i
	var c3 complex64
	fmt.Printf("Type of c1 and c2: %T and %T\n", c1, c2)
	fmt.Println(c1, c2, c3)

	// uintptr
	var number int = 8
	fmt.Println(unsafe.Pointer(&number))
	p := uintptr(unsafe.Pointer(&number))
	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Size of p: %d Bytes\n", unsafe.Sizeof(p))
	fmt.Println(p)
}