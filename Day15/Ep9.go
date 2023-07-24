package main

import "fmt"

func isPalindrome(x int) bool {
	reverse := 0
	number := x

	if x < 0 {
		return false
	}

	for number != 0 {
		reverse = reverse*10 + (number % 10)
		number /= 10
	}

	if reverse == x {
		return true
	}

	return false
}

func main() {
	var number int
	fmt.Printf("Input: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Output: %t\n", isPalindrome(number))
}
