package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	n := 0
	var m sync.Mutex

	// now, both goroutines call m.Lock() before accessing `n`
	// and call m.Unlock once they are done
	go func() {
		m.Lock()
		fmt.Println("Go 1")
		defer m.Unlock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		m.Lock()
		fmt.Println("Go 2")
		n++
		fmt.Println("Go 2: ", n)
		time.Sleep(2 * time.Second)
		m.Unlock()
	}()

	time.Sleep(1 * time.Second)
	m.Lock()
	fmt.Println("Hello")

	time.Sleep(5 * time.Second)

	fmt.Println(n)
}
