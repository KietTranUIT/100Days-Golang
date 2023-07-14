package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println("BUFFER IN GOLANG")

	var buffer bytes.Buffer

	buffer.WriteString("Tran Quang ")
	buffer.WriteString("Kiet")

	fmt.Println("Length of buffer = ", buffer.Len())
	fmt.Println("Capacity of buffer = ", buffer.Cap())
	fmt.Println("Buffer is: ", buffer)
	fmt.Println("String buffer is: ", buffer.String())
}