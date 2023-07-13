package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Triangle struct {
	side, height float64
}

func(r Rectangle) Area() float64 {
	return r.width * r.height
}

func(t Triangle) Area() float64 {
	return (t.side * t.height) / 2
}

func CalculateArea(s Shape) {
	fmt.Println("Area is: ", s.Area())
}

func main() {
	var opt int
	fmt.Println("Menu Options: ")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Triangle")

	fmt.Printf("Enter Options: ")
	fmt.Scanf("%d", &opt)

	if(opt == 1) {
		r := Rectangle{3.5, 4.5}
		CalculateArea(r)
	} else {
		t := Triangle{3.5, 4.5}
		CalculateArea(t)
	}
}