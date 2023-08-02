// Ep18: 4Sum
// Way 1

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := len(nums) - 1
			k := j + 1

			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					row := []int{nums[i], nums[j], nums[k], nums[l]}
					result = append(result, row)
					k++
					l--

					for k < l && nums[k] == nums[k-1] {
						k++
					}

					for l > k && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return result
}

func main() {
	var n int
loop:
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d\n", &n)

	if n < 1 || n > 200 {
		fmt.Println("Please enter again n!")
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

	result := fourSum(nums, target)
	fmt.Println("Output: ", result)
}
