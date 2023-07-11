package main

import "fmt"


func main() {
	// 3 Way declare a slice in go
	// 1.
	var slice1 []int = []int{1, 2, 3}
	fmt.Printf("Type of slice1 is: %T\n", slice1)
	fmt.Printf("Len = %d and Cap = %d\n", len(slice1), cap(slice1))
	fmt.Println(slice1)

	// 2.
	// give a array of int
	array := [5]int{1, 2, 3, 4, 5}
	slice2 := array[1:3]
	fmt.Printf("Len = %d and Cap = %d\n", len(slice2), cap(slice2))
	slice2[0] = 15
	fmt.Println("Slice2 is: ", slice2)
	fmt.Println("Array after modify slice2[0] is: ", array)

	// 3.
	// use make()
	slice3 := make([]int, 2, 5)
	fmt.Printf("Len = %d and Cap = %d\n", len(slice3), cap(slice3))
	fmt.Println("Slice3 is: ", slice3)

	// Modify slice

	// Change element of slice
	slice3[0] = 10000
	fmt.Println("Modify Element at index 0 of slice: ", slice3)

	// Add value to slice
	slice3 = append(slice3, 4, 5)
	fmt.Println("Add elements 4 and 5 to slice3: ", slice3)

	// Copy a slice
	slice4 := make([]int, 4, 5)
	copy(slice4, slice3)
	fmt.Println("Copy slice3 to slice4: ", slice4)

	// Delete element at index 1
	index := 1
	slice3 = append(slice3[:index], slice3[index + 1:]...)
	fmt.Println("Slice3 after remove element at index 1", slice3)

	// Slice in slice
	slice5 := [][]string {
		[]string{"1", "0", "0"},
		[]string{"0", "1", "0"},
		[]string{"0", "0", "1"},
	}

	for i,value := range slice5 {
		fmt.Println(i, " ", value)
	}


}