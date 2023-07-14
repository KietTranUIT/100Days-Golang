/*
type Context interface {
	Done() <-chan struct{} // Trả về một channel khi khi context đã bị hủy
	Err() error // Tại sao context lại bị hủy
	Deadline() (deadline time.Time, ok bool)  // Khi context có giá trị deadline ok trả ra là true, ngược lại nếu không có giá trị deadline thì trả ra là false
	Value(key interface{}) interface{}
}
*/

package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ctx,cancel := context.WithTimeout(context.Background(), time.Second * 10)
	time.Sleep(3 * time.Second)
	cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	fmt.Println("Canceled context!")
}