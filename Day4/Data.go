package main

import "fmt"

func main() {
	// Allocate with new
	number := new(int64)
	fmt.Println("Value of number is: ", *number)
	fmt.Println("Address of number is: ", number)
	fmt.Printf("Type of number is: %T\n", number)

	// Allocate with make: only apply for slice, map, channel
	// Create slice with make
	slice := make([]int, 10)
	fmt.Println("Slice is: ", slice)

	// Create map with make
	mp := make(map[string]string)
	fmt.Println("Map is: ", mp)
}