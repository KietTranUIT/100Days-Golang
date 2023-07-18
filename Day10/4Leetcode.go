package main

import "fmt"

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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := 0, 0
	nums := make([]int, len(nums1)+len(nums2))

	for i := 0; i < len(nums); i++ {
		if m >= len(nums1) {
			for j := n; j < len(nums2); j++ {
				nums[i] = nums2[j]
				i++
			}
			break
		}
		if n >= len(nums2) {
			for j := m; j < len(nums1); j++ {
				nums[i] = nums1[j]
				i++
			}
			break
		}

		if nums1[m] < nums2[n] {
			nums[i] = nums1[m]
			m++
		} else if nums2[n] < nums1[m] {
			nums[i] = nums2[n]
			n++
		} else {
			nums[i] = nums1[m]
			i++
			nums[i] = nums2[n]
			m++
			n++
		}
	}

	low, high := 0, len(nums)-1
	median := (low + high) / 2

	if len(nums)%2 == 0 {
		return float64(nums[median]+nums[median+1]) / 2
	}

	return float64(nums[median])
}

func main() {
	num1 := InputNumber()
	num2 := InputNumber()

	median := findMedianSortedArrays(num1, num2)
	fmt.Println("Median is: ", median)
}
