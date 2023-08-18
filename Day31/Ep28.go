package main

import "fmt"

func strStr(haystack string, needle string) int {
	j := 0
	distance := len(needle) - 1

	for i, _ := range haystack {
		if haystack[i] == needle[0] {
			j = i + distance
			if j > len(haystack)-1 {
				return -1
			}
			sub := haystack[i : j+1]
			if sub == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	var haystack, needle string
	fmt.Printf("Enter haystack: ")
	fmt.Scanf("%s", &haystack)

	fmt.Printf("Enter needle: ")
	fmt.Scanf("%s", &needle)

	result := strStr(haystack, needle)
	fmt.Println("Output: ", result)
}
