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
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
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

*/

// remember to close your channel
// if you do not close your channel, you will receive this error
// fatal error: all goroutines are asleep - deadlock!

// ************** IMPORTANT **************
// YOU NEED GO VERSION 1.5.2 OR GREATER
// otherwise you will receive this error
// fatal error: all goroutines are asleep - deadlock!

//记得关闭你的channel
//如果你不关闭你的channel，你将会收到这个错误。
//致命错误:所有的goroutines都睡着了——死锁!

// * * * * * * * * * * * * * *重要* * * * * * * * * * * * * *
//您需要版本1.5.2或更高版本。
//否则你将会收到这个错误。
//致命错误:所有的goroutines都睡着了——死锁!
