package main

import "fmt"

/* Name convention
    A name must begin with a letter, and can have any number of additional letters and numbers.
    A function name cannot start with a number.
    A function name cannot contain spaces.
    If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
    If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
    function names are case-sensitive (car, Car and CAR are three different variables).
*/

// User defined a function types in golang
type User func(string) string
	


// Create a simple function no parameters and return value
func printHello() {
	fmt.Println("Hello World!")
}

// Create a function with paremeters
func printNumber(number int) {
	fmt.Println("Number: ", number)
}

// Create a function with parameters and return value
func sumNumber(a, b int) int {
	return a + b
}

// Create a function with return values
func nameNumber(number int) (int, string) {
	number = 1
	return number, "one"
}

// The return value of function can be named
func multiNumber(a, b int) (mul int) {
	mul = a * b
	return mul
}

// Passing address to function
func changeNumber(number *int) {
	*number = 10
}

// Higher-order function
func mulSum(a int) func(int) int {
	return func(b int) int {
		return sumNumber(a, b) * 2
	}
}

func printUser(name string) User {
	return func(id string) string {
		return name + " " + id
	}
}

// Variadic function
func printColor(colors ...string) {
	length := len(colors)
	fmt.Println(length)
	for _,color := range colors {
		fmt.Printf("%s ", color)
	}
	fmt.Println()
}

// function for deferred function calls
func first() {
	fmt.Println("First function")
}

func second() {
	fmt.Println("Second function")
}

func third() {
	fmt.Println("Third function")
}


func main() {
	printHello()

	printNumber(200)

	sum := sumNumber(1, 2)
	fmt.Println("Sum of 1 and 2 is: ", sum)

	n, name_number := nameNumber(10)
	fmt.Printf("Number is %d and Name number is %s\n", n, name_number)
	mul := multiNumber(2, 3)
	fmt.Println("Mul of 2 and 3 is: ", mul)

	number := 1
	fmt.Println("Number before: ", number)
	changeNumber(&number)
	fmt.Println("Number after: ", number)

	// Create a Anonymous function
	a := func(a, b int) int {
		return a + b
	}
	fmt.Println("Return of Anonymous function is: ", a(5, 10))

	// Closure function
	l := 3.14
	b := func() float64 {
		l += l
		return 2 * l
	}
	fmt.Println("Value of l is: ", l)
	fmt.Println("Return of Closure function is: ", b())

	higher := mulSum(2)
	fmt.Println("Value of higher is: ", higher(3))

	fmt.Println(printUser("TranQuangKiet")("21522265"))

	// Test vadiaric function
	printColor("red")
	printColor("red", "yellow", "green", "blue")

	// Deferred function call
	defer third()
	first()
	second()
}