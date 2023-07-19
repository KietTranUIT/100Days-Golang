package main

import "fmt"

func main() {
	var number int = 10

	fmt.Printf("Enter number CountDown: ")
	//fmt.Scanf("%d", &number)

	for i := 0; i < number; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	number = number + 100

	fmt.Println(number)
}
