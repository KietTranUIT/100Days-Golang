package main

import "fmt"

// All defer are stored in list with LIFO (Last in first out)

func Defer1(i int) {
	fmt.Printf("defer 1/ i = %d\n", i)
}

func Defer2(i int) {
	fmt.Printf("defer 2/ i = %d\n", i)
}

func Count(i int) int {
	defer Defer1(i)
	i = i + 1
	defer Defer2(i)
	i = i + 1
	return i
}

func main() {
	i := Count(0)
	fmt.Println("I = ", i)
}