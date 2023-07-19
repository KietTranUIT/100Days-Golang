package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	state := make([][]bool, n)
	start, max := 0, 1

	for i := 0; i < n; i++ {
		state[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		state[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			state[i][i+1] = true
			start = i
			max = 2
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			j := i + k - 1
			if state[i+1][j-1] == true && s[i] == s[j] {
				state[i][j] = true
				start = i
				max = k
			}
		}
	}
	return s[start : start+max]
}

func main() {
	var s string
	fmt.Printf("Input: ")
	fmt.Scanf("%s", &s)
	fmt.Printf("Output: %s\n", longestPalindrome(s))
}
