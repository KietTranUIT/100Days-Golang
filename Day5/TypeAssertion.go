package main

import "fmt"

func main() {
	// Type Assertion t := i.(Type data)
	var a interface{} = "Tran Quang Kiet"

	str := a.(string)
	fmt.Printf("Type of str is %T and Value is %s\n", str, str)

	number, ok := a.(int)
	fmt.Printf("Type of number is %T and Value is %s | %s\n", number, number, ok)

	// Type switch
	var b interface{} = 45.67
	switch t := b.(type) {
	case int:
		fmt.Println("Type of b is int: ", t)
	case string:
		fmt.Println("Type of b is string: ", t)
	case bool: 
	    fmt.Println("Type of b is bool: ", t)
	case float64:
		fmt.Println("Type of b is float64: ", t)
	default: 
	    fmt.Println("None type!")
	}
}