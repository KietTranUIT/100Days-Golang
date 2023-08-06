package main

import "fmt"

func generate(result []string, s string, open, close, n int) []string {
	// Base condition
	if open == n && close == n {
		result = append(result, s)
		return result
	}
	// If the number of _open parentheses is less than the given n
	if open < n {
		result = generate(result, s+"(", open+1, close, n)
	}
	// If we need more close parentheses to balance
	if close < open {
		result = generate(result, s+")", open, close+1, n)
	}
	return result
}

func generateParenthesis(n int) []string {
	// Resultant list
	var result []string
	// Recursively generate parentheses
	result = generate(result, "", 0, 0, n)
	return result
}

func main() {
	var n int
loop:
	fmt.Printf("Input: ")
	fmt.Scanf("%d", &n)

	if n < 1 || n > 8 {
		fmt.Printf("Please enter again n!")
		goto loop
	}

	str := generateParenthesis(n)
	fmt.Println("Ouput: ", str)
}
