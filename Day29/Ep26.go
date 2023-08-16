package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	count := len(nums)
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			count--
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return count
}

func main() {
	var n int
loop:
	fmt.Printf("Enter the size of nums: ")
	fmt.Scanf("%d", &n)

	if n < 0 {
		fmt.Printf("Please enter again the size of nums!")
		goto loop
	}

	nums := make([]int, n)
	fmt.Printf("Enter nums:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	fmt.Println(nums)

	result := removeDuplicates(nums)

	fmt.Println("Ouput: ", result)
	fmt.Println(nums)
}
