// Ep20: Valid Parentheses
package main

import "fmt"

func pop(stack []byte) []byte {
	return stack[:len(stack)-1]
}

func isValid(s string) bool {
	var stack []byte
	top := -1

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			top++
			continue
		}

		if s[i] == ')' {
			if top == -1 {
				return false
			}
			if stack[top] != '(' {
				return false
			}
		} else if s[i] == ']' {
			if top == -1 {
				return false
			}
			if stack[top] != '[' {
				return false
			}
		} else {
			if top == -1 {
				return false
			}
			if stack[top] != '{' {
				return false
			}
		}

		top--
		stack = pop(stack)
	}

	if top != -1 {
		return false
	}
	return true
}

func main() {
	var s string
	fmt.Printf("Input: ")
	fmt.Scanf("%s", &s)

	fmt.Println("Output: ", isValid(s))
}
