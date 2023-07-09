package main

import "fmt"

func main() {
	// for statement1; statement2; statement3 {}
	sum1 := 0
	for i := 0; i < 5; i++{
		sum1 += i
	}
	fmt.Printf("Sum (1->5) = %d\n", sum1)

	// While
	sum2 := 0
	for sum2 < 10 {
		sum2 += 1
	}
	fmt.Printf("Sum = %d\n", sum2)

	// Range
	var numbers []int = []int{1, 2, 3, 4, 5}
	for index_numbers, value_numbers := range numbers {
		fmt.Printf("Index %d = %d\n", index_numbers, value_numbers)
	}
 
	// Go to
	a := 0
	loop: a++
	if a < 5 {
	    goto loop
	}
	fmt.Println(a)

    // Infinity loop
	for {

	}
}