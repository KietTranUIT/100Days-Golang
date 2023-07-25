package main

import (
	"fmt"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		if area := Min(height[i], height[j]) * (j - i); area > max {
			max = area
		}

		if height[j] < height[i] {
			j--
		} else {
			i++
		}
	}

	return max
}

func main() {
	n := 0

loop:
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d\n", &n)

	if n < 2 {
		fmt.Println("Please enter n >= 2")
		goto loop
	}

	height := make([]int, n)

	fmt.Printf("Enter height: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &height[i])
	}

	fmt.Printf("Output: %d\n", maxArea(height))
}
