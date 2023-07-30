package main

import (
	"fmt"
	"sort"
)

func Sort(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func UniRow(row []int, result [][]int) bool {
	for _, rows := range result {
		if rows[0] == row[0] && rows[1] == row[1] && rows[2] == row[2] {
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	nums = Sort(nums)
	i := 0
	var result [][]int
	count := 0

	for i < (len(nums) - 2) {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			if -nums[i] == nums[j]+nums[k] {
				newRow := []int{nums[i], nums[j], nums[k]}
				if UniRow(newRow, result) {
					result = append(result, newRow)
					count++
				}
				j++
				k--
			} else if -nums[i] < nums[j]+nums[k] {
				k--
			} else {
				j++
			}
		}
		i++
	}
	return result
}

func main() {
	var n int
loop:
	fmt.Printf("Enter size of nums: ")
	fmt.Scanf("%d", &n)

	if n < 3 || n > 3000 {
		fmt.Printf("Pleasse enter again n!\n")
		goto loop
	}

	nums := make([]int, n)
	fmt.Printf("Input: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	result := threeSum(nums)

	fmt.Println(result)
}
