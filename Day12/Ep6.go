package main

import (
	"fmt"
	"strings"
)

type Direction struct {
	up, down bool
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	dr := Direction{
		up:   false,
		down: false,
	}

	row := 0
	result := make([]string, numRows)

	for i := 0; i < len(s); i++ {
		result[row] += string(s[i])

		if row == 0 {
			dr.down = true
			dr.up = false
			row++
		} else if row == numRows-1 {
			dr.up = true
			dr.down = false
			row--
		} else if dr.up == true {
			row--
		} else if dr.down == true {
			row++
		}
	}

	return strings.Join(result, "")
}

func main() {
	var s string
	var n int

loop:
	fmt.Printf("Input: ")
	fmt.Scanf("%s %d", &s, &n)

	if s == "" {
		fmt.Println("Please enter a string not nil!")
		goto loop
	} else if n <= 0 {
		fmt.Println("Please enter a number >= 1!")
		goto loop
	}

	fmt.Printf("Output: %s\n", convert(s, n))
}
