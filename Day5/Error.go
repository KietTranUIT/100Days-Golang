package main

import (
	"fmt"
	"errors"
	"math"
)


// Go Errors use New()
func Sqrt(number float64) (float64, error) {
	if(number < 0) {
		return 0, errors.New("Math sqrt: number is negative!")
	}
	return math.Sqrt(number), nil
}

// Go Errors use fmt.Errorf()
func devide(a, b float64) (float64, error) {
	if(b == 0) {
		return 0, fmt.Errorf("Error: %f / %f no calculate", a, b)
	}
	return a / b, nil
}

// Custom Errors in Golang
type devideByZero struct {
	a, b float64
}

func(e *devideByZero) Error() string {
	return fmt.Sprintf("Error! No calculate devide by zero: %f / %f", e.a, e.b)
}

func DevideNew(a, b float64) (float64, error){
	if(b == 0) {
		return 0, &devideByZero{a,b}
	}
	return a / b, nil
}

func main() {
	a, err1 := Sqrt(float64(4))

	fmt.Printf("a = %g\n", a)
	if(err1 != nil) {
		fmt.Println(err1)
	}

	b, err2 := Sqrt(float64(-2))

	fmt.Printf("b = %g\n", b)
	if(err2 != nil) {
		fmt.Println(err2)
	}

	c, err3 := devide(4, 0)
	if(err3 != nil) {
		fmt.Println(err3)
	} else {
		fmt.Println("c = ", c)
	}

	d, err4 := DevideNew(4, 0)
	if(err4 != nil) {
		fmt.Println(err4)
	} else {
		fmt.Println("d = ", d)
	}
}