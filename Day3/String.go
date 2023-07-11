package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Tran Quang Kiet"

	// Some functions for string
	
	// Length of string
	fmt.Printf("Length of string is: %d\n", len(name))

	// check Whether string s contains substr or not
	check1 := strings.Contains(name, "Kiet")
	if(check1 == true) {
		fmt.Println("name have Kiet!")
	} else {
		fmt.Println("name don't contain Kiet!")
	}

	// find index of a character in string
	index := strings.Index(name, "Quang")
	fmt.Println("index of Quang in name is: ",index)

	// replace characters in string
	new_name := strings.Replace(name, " ", "-", -1)
	fmt.Println("name after replace ' ' to '-' is: ",new_name)

	// Lower string
	name_lower := strings.ToLower(name)
	fmt.Println("name to lower: ", name_lower)

	// Upper string
	name_upper := strings.ToUpper(name)
	fmt.Println("name to upper: ", name_upper)

	// split string
	str := strings.Split(name, " ")
	fmt.Println("Array string is: ", str)

	// Has Prefix string
	check2 := strings.HasPrefix(name, "Tran")
	if(check2 == true) {
		fmt.Println("name start with Tran!")
	} else {
		fmt.Println("name no start with Tran!")
	}

	// Has suffix string
	check3 := strings.HasSuffix(name, "Kiet")
	if(check3 == true) {
		fmt.Println("name end with Kiet!")
	} else {
		fmt.Println("name no end with Kiet!")
	}
}