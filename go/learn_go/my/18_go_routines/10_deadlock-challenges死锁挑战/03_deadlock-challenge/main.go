package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c) // 0
}

// Why does this only print zero? 为什么这只打印0 ?
// And what can you do to get it to print all 0 - 9 numbers?
