package main

import (
	"fmt"
	"errors"
)

func HandlePanic(a, b *float64) {
	r := recover()
	if(r != nil) {
		fmt.Println("Recover panic: ", r)
	}
}

func Devide(a, b float64) (d float64, e error) {
	e = errors.New("Can't devide a number by zero")
	defer HandlePanic(&a, &b)
	if(b == 0) {
		panic("Can't devide a number by zero")
	}
	e = nil
	d = a / b
	return d, e
}

func main() {
	d1,e1 := Devide(4, 2)
	if(e1 != nil) {
		fmt.Println(e1)
	} else {
	    fmt.Println(d1)
	}


	d,e := Devide(4, 0)
	if(e != nil) {
		fmt.Println(e)
	} else {
	    fmt.Println(d)
	}


	d2,e2 := Devide(4, 1)
	if(e2 != nil) {
		fmt.Println(e2)
	} else {
	    fmt.Println(d2)
	}
}