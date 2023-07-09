package main

import "fmt"

func main() {
	// If else statements
	a := 2
	if a == 0 {
		fmt.Println("a == 0")
	} else if a < 0 {
		fmt.Println("a < 0")
	} else {
		fmt.Println("a > 0")
	}

	// Switch
	b := 3
	switch(b) {
	case 1: fmt.Println(b)
	case 2: fmt.Println(b)
	case 3: fmt.Println(b)
	case 4: fmt.Println(b)
	default: fmt.Println("Nothing")
	}
}