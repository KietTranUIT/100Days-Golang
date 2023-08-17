package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	count := len(nums)
	j, i := 0, -1
	for j < len(nums) {
		if nums[j] != val {
			i++
			nums[i] = nums[j]
		} else {
		        count--
		}
		j++
	}
	return count
}

func main() {
	var n int
	fmt.Printf("Enter size of nums:")
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	fmt.Printf("Enter values of nums:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	var val int
	fmt.Printf("Enter val: ")
	fmt.Scanf("%d", &val)
	fmt.Println("Output: ", removeElement(nums, val))
	fmt.Println(nums)
}
