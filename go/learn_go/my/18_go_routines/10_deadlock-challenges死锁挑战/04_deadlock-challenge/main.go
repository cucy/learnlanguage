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

	for {
		fmt.Println(<-c)
	}
}

/*
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()

*/

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// Why are all goroutines asleep?
// How can we fix this?

//这是打印的号码，但我们收到了这个错误:
//致命错误:所有的goroutines都睡着了——死锁!
//死锁在哪里?
//为什么所有的goroutines都睡着了?
//我们如何解决这个问题?
