package main

import (
	"fmt"
	"context"
)

func CalA(ctx context.Context) {
	number := ctx.Value("Number one")

	if(number != nil) {
		fmt.Println("Number one is: ", number.(int))
		ctx1 := context.WithValue(ctx, "Number two", 2013)
		CalB(ctx1)
	}
}

func CalB(ctx context.Context) {
	number := ctx.Value("Number two")
	
	if(number != nil) {
		fmt.Println("Number two is: ", number.(int))
		fmt.Println("Sum of Number one and Number two is: ", ctx.Value("Number one").(int) + ctx.Value("Number two").(int))
	}

}

func main() {
	ctx := context.WithValue(context.Background(), "Number one", 2003)
	CalA(ctx)
}