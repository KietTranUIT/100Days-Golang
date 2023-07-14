package main

import (
	"fmt"
	"time"
	"context"
)

func DoSomeThing(ctx context.Context) {
	chanel := make(chan bool)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			chanel <- true
		}
	}()

	check := <- chanel
	if(check == true) {
		fmt.Println("Time out ...")
		return
	}

	fmt.Println("Do some thing ...")
}

func main() {
	ctx,_ := context.WithTimeout(context.Background(), time.Second * 5)

	DoSomeThing(ctx)
}