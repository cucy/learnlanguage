package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
https://golang.org/ref/spec#Channel_types

可选的<-操作符指定通道方向，发送或接收。
如果没有方向，信道是双向的。
*/
