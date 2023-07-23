package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	i, j := 0, 0
	var number int = 0

	if s[0] < '0' || s[0] > '9' {
		if s[0] == '-' {
			j = 1
			goto loop
		}

		if s[0] == '+' {
			j = 1
			goto loop
		}

		return 0
	}

loop:
	for {
		if j >= len(s) {
			number, _ = strconv.Atoi(s[i:j])
			break
		}

		if s[j] < '0' || s[j] > '9' {
			number, _ = strconv.Atoi(s[i:j])
			break
		}
		j++
	}

	max := math.MaxInt32
	min := math.MinInt32

	if number < min {
		return min
	}

	if number > max {
		return max
	}

	return number
}

func main() {
	fmt.Printf("Input: ")
	s := Input()
	number := myAtoi(s)
	fmt.Printf("Output: %d\n", number)
}
