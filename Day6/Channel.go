package main

import (
	"fmt"
	"time"
)

// This is a bidirectional channel
func Say(c chan bool) {
	fmt.Println("Say Hello Goroutines...!")
	time.Sleep(2 * time.Second)
	c <- true
}

// This is a unidirectional channel
func SendData(c chan<- string) {
	c<- "Tran Quang Kiet"
}

// Close channel and for loop on channel
func CountDown(c chan int) {
	for i:=0; i<10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan bool)
	go Say(c)
	<- c
	fmt.Println("End Program for bidirectional channel...!")

	c_new := make(chan<- string)
	go SendData(c_new)
	//fmt.Println("Data received from a unidirectional channel: ", <-c_new)
	// Error: invalid operation: cannot receive from send-only channel c_new (variable of type chan<- string)
	

	chanl := make(chan int)
	go CountDown(chanl)
	for {
		i,ok := <- chanl
		if(ok == false) {
			break
		}
		fmt.Println("Received i from channel: ", i)
	}
	// or use for range on channel
	chanl1 := make(chan int)
	go CountDown(chanl1)
	for i:= range chanl1 {
		fmt.Println("Received i from channel: ", i)
	}

	// Buffer in channel
	// Sender block when the buffer is full
	// Receiver block when the buffer is empty
	chanl2 := make(chan int, 2)
	chanl2<-1
	chanl2<-2
	fmt.Println(<-chanl2)
	fmt.Println(<-chanl2)
}