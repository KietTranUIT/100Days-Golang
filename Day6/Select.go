package main

import "fmt"

func Count(chanl, quit chan int) {
	number := 0
	for {
		select {
		case x := <-chanl:
			number+=x
		case <-quit:
			fmt.Println("Sum is: ", number)
			fmt.Println("End Goroutine...!")
		}
	}
}

func main() {
	chanl := make(chan int)
	quit := make(chan int)
	go Count(chanl, quit)
	for i:=0; i<10; i++ {
		chanl <- i
	}
	quit <- 0
	fmt.Println("Good bye...!")
}