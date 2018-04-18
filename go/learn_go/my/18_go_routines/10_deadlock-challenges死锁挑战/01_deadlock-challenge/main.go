package main

import "fmt"

func main() {

	c := make(chan int)

	c <- 1
	fmt.Println(<-c)
}

/*
fatal error: all goroutines are asleep - deadlock!
*/

// This results in a deadlock.
// Can you determine why? 你能确定为什么吗?
// And what would you do to fix it?
