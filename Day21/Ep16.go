package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i := 0
	var j, k int
	var closest int
	var m int = 20000

	for i < len(nums)-2 {
		j = i + 1
		k = len(nums) - 1

		for j < k {
			sum3 := nums[i] + nums[j] + nums[k]

			if sum3 > target {
				if sum3-target < m {
					m = sum3 - target
					closest = sum3
				}
				k--
			} else if sum3 < target {
				if target-sum3 < m {
					m = target - sum3
					closest = sum3
				}
				j++
			} else {
				return sum3
			}
		}
		i++
	}
	return closest
}

func main() {
	var n int
loop:
	fmt.Printf("Enter size of slice n: ")
	fmt.Scanf("%d", &n)

	if n < 3 || n > 300 {
		fmt.Printf("Please enter again size of slice n!")
		goto loop
	}

	nums := make([]int, n)

	fmt.Printf("Enter nums: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	var target int
	fmt.Printf("Enter target: ")
	fmt.Scanf("%d", &target)

	fmt.Println("Output: ", threeSumClosest(nums, target))
}
