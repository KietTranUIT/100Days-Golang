/*Viết một chương trình với các yêu cầu sau: 
- Nhập vào một số la mã bất kì
- In ra dạng số của số la mã vừa nhập vào
*/

package main

import (
	"fmt"
	"strings"
)

var conversion_table map[string]int
var list_key []string

func Initial() {
	conversion_table = make(map[string]int)
	conversion_table["I"] = 1
	conversion_table["V"] = 5
	conversion_table["X"] = 10
	conversion_table["L"] = 50
	conversion_table["C"] = 100
	conversion_table["D"] = 500
	conversion_table["M"] = 1000

	list_key = []string{"I", "V", "X", "L", "C", "D", "M"}
}

func CheckInput(number_roman string) bool {
	list := strings.Join(list_key, " ")
	for _,char:= range number_roman {
		if(strings.Index(list, string(char)) == -1) {
			return false
		}
	}
	return true
}

func NumberRomanToNumber(number_roman string) (sum int) {
	sum = 0
	for i:=0; i<len(number_roman); i++ {
		list := strings.Join(list_key, "")
		if(i == len(number_roman) - 1) {
		    n1 := conversion_table[list_key[strings.Index(list, string(number_roman[i]))]]
			sum += n1
			break
		}
		n1 := conversion_table[list_key[strings.Index(list, string(number_roman[i]))]]
		n2 := conversion_table[list_key[strings.Index(list, string(number_roman[i + 1]))]]
		if(n1 < n2) {
			sum += (n2 - n1)
			i++
		} else {
			sum += n1
		}
	}
	return sum
}

func main() {
	Initial()

	var number_roman string
	loop: fmt.Println("Enter number roman: ")
	fmt.Scanf("%s", &number_roman)
	number_roman = strings.ToUpper(number_roman)
	
	check := CheckInput(number_roman)
	if(check == false) {
		fmt.Println("Please enter a number roman: ")
		goto loop
	}
	
	number := NumberRomanToNumber(number_roman)
	fmt.Printf("Number of %s is: %d\n", number_roman, number)
}