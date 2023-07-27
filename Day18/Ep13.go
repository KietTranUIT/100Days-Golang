package main

import (
	"fmt"
)

func romanToInt(s string) int {
	roman := make(map[string]int)

	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	number := 0

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			number += roman[string(s[i])]
			break
		}

		if roman[string(s[i])] < roman[string(s[i+1])] {
			number += roman[string(s[i+1])] - roman[string(s[i])]
			i++
			continue
		}

		number += roman[string(s[i])]
	}
	return number
}

func main() {
	var s string
	fmt.Printf("Input: ")
	fmt.Scanf("%s", &s)
	fmt.Printf("Output: %d\n", romanToInt(s))
}
