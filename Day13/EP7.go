package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var number int
	min := math.MinInt32
	max := math.MaxInt32

	for x != 0 {
		modulo := x % 10

		number = (number * 10) + modulo
		x /= 10
	}

	if number < min || number > max {
		return 0
	}

	return number
}

func main() {
	var number int
	fmt.Printf("Input: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Output: %d\n", reverse(number))

}
