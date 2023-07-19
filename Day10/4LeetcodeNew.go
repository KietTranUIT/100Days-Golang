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
	m := len(nums1)
	n := len(nums2)

	half := (m + n) / 2

	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	if n == 0 {
		median := (0 + (m - 1)) / 2
		if m%2 == 0 {
			return float64(nums1[median]+nums1[median+1]) / 2
		} else {
			return float64(nums1[median])
		}
	}

	left_nums1, right_nums1 := 0, m-1
	mid_nums1, mid_nums2 := 0, 0

	for {
		mid_nums1 = (left_nums1 + right_nums1) / 2
		mid_nums2 = half - mid_nums1 - 2

		if mid_nums2 >= -1 {
			fmt.Println(mid_nums1, mid_nums2)
			if mid_nums1 == -1 {
				goto flag
			}

			if nums1[mid_nums1] > nums2[mid_nums2+1] {
				right_nums1--
				continue
			}

			if mid_nums2 == -1 {
				goto flag
			}

			if nums1[mid_nums1+1] < nums2[mid_nums2] {
				left_nums1++
				continue
			}
		}
	flag:

		if (m+n)%2 == 0 {
			var max, min int
			if mid_nums1 == m-1 && mid_nums2 < 0 {
				min = nums2[mid_nums2+1]
				max = nums1[mid_nums1]
				return float64(min+max) / 2
			}
			if mid_nums1 == m-1 {
				min = nums2[mid_nums2+1]
				max = ChooseMax(nums1[mid_nums1], nums2[mid_nums2])
				return float64(min+max) / 2
			}
			if mid_nums2 < 0 {
				max = nums1[mid_nums1]
				min = ChooseMin(nums1[mid_nums1+1], nums2[mid_nums2+1])
				return float64(min+max) / 2
			}
			if mid_nums1 != m-1 && mid_nums2 >= 0 {
				min = ChooseMin(nums1[mid_nums1+1], nums2[mid_nums2+1])
				max = ChooseMax(nums1[mid_nums1], nums2[mid_nums2])
				return float64(min+max) / 2
			}
		} else {
			if mid_nums1 >= m-1 {
				return float64(nums2[mid_nums2+1])
			}
			return float64(ChooseMin(nums1[mid_nums1+1], nums2[mid_nums2+1]))
		}
	}
}

func main() {
	num1 := InputNumber()
	num2 := InputNumber()

	median := findMedianSortedArrays(num1, num2)
	fmt.Println("Median is: ", median)
}
