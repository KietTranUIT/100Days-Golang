package main

import (
	"fmt"
)

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	number := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	romanNumber := ""

	for num != 0 {
		for i := 12; i >= 0; i-- {
			if number[i] <= num {
				romanNumber += string(roman[i])
				num -= number[i]
				break
			}
		}
	}
	return romanNumber
}

func main() {
	num := 0
	fmt.Printf("Input: ")
	fmt.Scanf("%d", &num)

	fmt.Printf("Output: %s\n", intToRoman(num))
}
