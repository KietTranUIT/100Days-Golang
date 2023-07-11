package main

import "fmt"

func inputArray(array [5]int) [5]int {
	length := len(array)
	for i:=0; i<length; i++ {
		_, err := fmt.Scanf("%d", &array[i])
		if(err != nil) {
			panic(err)
		}
	}
	return array
}

func printArray(array [5]int) {
	for i:=0; i<len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
}

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Size of array: %d\n", len(array))
	fmt.Println("Array is: ", array)

	var array1 [5]int
	array1 = inputArray(array1)
	printArray(array1)

	var array2 [5]int
	array2 = array1
	printArray(array2)
}