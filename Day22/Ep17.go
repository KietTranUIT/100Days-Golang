// 17. Letter Combinations of a Phone Number
package main

import "fmt"

func letterCombinations(digits string) []string {
	var ans []string
	l := len(digits)

	if l == 0 {
		return ans
	}

	mp := make(map[string]string, 8)
	mp["2"] = "abc"
	mp["3"] = "def"
	mp["4"] = "ghi"
	mp["5"] = "jkl"
	mp["6"] = "mno"
	mp["7"] = "pqrs"
	mp["8"] = "tuv"
	mp["9"] = "wxyz"

	var f func(int, string)
	f = func(pos int, str string) {
		if pos == l {
			ans = append(ans, str)
		} else {
			letter := mp[string(digits[pos])]
			for i := 0; i < len(letter); i++ {
				f(pos+1, str+string(letter[i]))
			}
		}
	}
	f(0, "")
	return ans
}

func main() {
	var digits string
loop:
	fmt.Printf("Enter digits string: ")
	fmt.Scanf("%s", &digits)

	if l := len(digits); l < 0 || l > 4 {
		fmt.Printf("Please enter again digits string!")
		goto loop
	}

	for i := 0; i < len(digits); i++ {
		if digits[0] < '2' || digits[0] > '9' {
			fmt.Printf("Please enter again digits string!")
			goto loop
		}
	}
	ans := letterCombinations(digits)
	fmt.Println("Output: ", ans)
}
