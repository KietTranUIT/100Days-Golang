// Ep18: 4Sum
// Way 1

package main

import (
	"fmt"
	"sort"
)

type aux struct {
	sum    int
	n1, n2 int
}

func check(a, b aux) bool {
	if a.n1 == b.n1 || a.n2 == b.n2 || a.n1 == b.n2 || a.n2 == b.n1 {
		return false
	}
	return true
}

func Sort(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

var count int = 0

func checkUni(a, b, c, d int, result [][]int) bool {
	compare := []int{a, b, c, d}
	compare = Sort(compare)

	for i := 0; i < len(result); i++ {
		row := result[i]
		row = Sort(row[:])
		if row[0] == compare[0] && row[1] == compare[1] && row[2] == compare[2] && row[3] == compare[3] {
			return false
		}
	}
	return true
}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	l := n * (n - 1) / 2
	temp := make([]aux, l)
	index := 0
	var result [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp[index].sum = nums[i] + nums[j]
			temp[index].n1 = i
			temp[index].n2 = j
			index++
		}
	}

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].sum < temp[j].sum
	})

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if temp[i].sum+temp[j].sum == target && check(temp[i], temp[j]) {
				if checkUni(nums[temp[i].n1], nums[temp[i].n2], nums[temp[j].n1], nums[temp[j].n2], result) {
					r := []int{nums[temp[i].n1], nums[temp[i].n2], nums[temp[j].n1], nums[temp[j].n2]}
					result = append(result, r)
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
