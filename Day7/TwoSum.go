package main

import (
	"fmt"
)

const NULL = -10000000000

type Number struct {
	index, value int
}

type HashTable struct {
	nums []Number
}

func Hashing(value int, len int) int {
	return value % len
}

func SearchHashTable(h HashTable, value int, i int) int {
	number := value
	if value < 0 {
		number = -number
	}
	key := Hashing(number, len(h.nums) - 1)

	for h.nums[key].index != -1 {
		if(h.nums[key].value == value && h.nums[key].index != i) {
			return h.nums[key].index
		}
		key++
		if(key >= len(h.nums)) {
			key = 0
		}
	}
	return -1
}

func AddHashTable(h HashTable, index int, value int) HashTable {
	number := value
	if(value < 0) {
		number = -number
	}
	key := Hashing(number, len(h.nums) - 1)
	
	for h.nums[key].index != -1 {
		key++
		if(key >= len(h.nums)) {
			key = 0
		}
	}

	h.nums[key] = Number{index: index, value: value}
	return h
}

func InitialHashTable(nums []int) HashTable {
	var h HashTable
	len_nums := len(nums)

	h.nums = make([]Number, len_nums + 1)
	for i,_ := range h.nums {
		h.nums[i].value = NULL
		h.nums[i].index = -1
	}

	for i,num := range nums {
		h = AddHashTable(h, i, num)
	}
	return h
}

func twoSum(nums []int, target int) []int {
    var result []int
	
	if(len(nums) < 2) {
		return result
	}

	h := InitialHashTable(nums)

	for i,num := range nums {
		search := target - num
		j := SearchHashTable(h, search, i)
		if j != -1 {
			result = append(result, i, j)
			return  result
	    }
	}
	return result

}

func InputNums(nums []int) []int {
	num := 0
	len_nums := 0
	loop: fmt.Printf("Enter length of nums: ")
	fmt.Scanf("%d", &len_nums)
	
	if(len_nums < 0) {
		goto loop
	}

	for i:=0; i<len_nums; i++ {
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}
	return nums
}

// Input: nums, target
func main() {
	var nums []int
	var target int
	nums = InputNums(nums)
	fmt.Printf("Enter target: ")
	fmt.Scanf("%d", &target)
	
	result := twoSum(nums, target)
	fmt.Println(result)
}