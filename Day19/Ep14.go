package main

import (
	"fmt"
)

func LCP(a, b string) string {
	var max int
	if len(a) < len(b) {
		max = len(a)
	} else {
		max = len(b)
	}

	result := ""
	for i := 0; i < max; i++ {
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			break
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	size := len(strs)
	index, i := 0, 1
	var lcp = ""

	for i < size {
		cal := LCP(strs[index], strs[i])

		if i == 1 {
			lcp = cal
			i++
			continue
		}

		if lcp != cal {

			if LCP(lcp, cal) == "" {
				return ""
			} else {
				lcp = LCP(lcp, cal)
			}
		}
		i++
	}
	return lcp
}

func main() {
	var n int
loop:
	fmt.Printf("Enter size of slice string: ")
	fmt.Scanf("%d", &n)

	if n < 1 || n > 200 {
		fmt.Printf("Please enter again n!")
		goto loop
	}

	strs := make([]string, n)

	fmt.Printf("Enter string: \n")
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &strs[i])
	}
	fmt.Printf("Output: %s\n", longestCommonPrefix(strs))
}
