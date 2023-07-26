package main

import (
	"fmt"
)

func InputNumber() []int {
	size := 0
loop:
	fmt.Printf("Enter Size of Number: ")
	fmt.Scanf("%d", &size)

	if size < 0 {
		fmt.Println("Please enter a size of Number positive!")
		goto loop
	}

	nums := make([]int, size)
	fmt.Printf("Enter number: ")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	return nums
}

func ChooseMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ChooseMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, len(nums1) - 1
	half := (len(nums1) + len(nums2)) / 2

	for {
		mid1 := 
	}
}

func main() {
	num1 := InputNumber()
	num2 := InputNumber()

	median := findMedianSortedArrays(num1, num2)
	fmt.Println("Median is: ", median)
}
