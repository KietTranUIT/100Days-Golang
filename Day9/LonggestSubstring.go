package main

import (
	"fmt"
	"bufio"
	"os"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Search(begin, end int, s string, char byte) int {
	for i := begin; i < end; i++ {
		if char == s[i] {
			return i
		}
    }
	return -1
}

func lengthOfLongestSubstring(s string) int {
	SubMax, j := 1, 0
	s_len := len(s)

	if s_len == 0 {
		return 0
	}

	for i:=1; i<s_len; i++ {
		check := Search(j, i, s, s[i])

		if check != -1 {
			j = check + 1
			continue
		}

		if len(s[j:i + 1]) > SubMax {
			SubMax = len(s[j:i + 1])
		}
	}

	return SubMax
}

func main() {
	fmt.Printf("Input: ")
	s := Input()

	result := lengthOfLongestSubstring(s)
	fmt.Println("Output: ", result)
}