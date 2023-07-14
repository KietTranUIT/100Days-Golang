package main

import (
	"fmt"
	"time"
)

var count int = 1000

func CountDown() {
	for i:=count; i>=0; i-- {
		fmt.Printf("Goroutines: %d\n", count)
		count--
	}
}

func main() {
	go CountDown()
	
	time.Sleep(1 * time.Millisecond)
	for i:=count; i>0; i-- {
		fmt.Printf("Main: %d\n", count)
		count--
	}
}

// race condition between main() and CountDown() on count variable
