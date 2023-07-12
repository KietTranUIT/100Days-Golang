package main

import "fmt"

type Person struct {
	name string
	age int
	height float64
}

func main() {
	// Initialize Person
	person := Person{"Tran Quang Kiet", 21, 1.72}

	// Access filed of struct
	fmt.Println("Name is: ", person.name)
	fmt.Println("Age is: ", person.age)
	fmt.Println("Height is: ", person.height)

	// Initialize pointer Person
	p := &Person{"Tran Nhat Anh Thu", 11, 1.3}
	fmt.Println("Name is: ", p.name)
	fmt.Println("Age is: ", p.age)
	fmt.Println("Height is: ", p.height)
}