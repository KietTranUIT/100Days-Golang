package main

import "fmt"

func main() {
	s := "Hello kiet"

	str := s[3:7]
	for i,_ := range str {
		fmt.Printf("%d ", i)
	}
}