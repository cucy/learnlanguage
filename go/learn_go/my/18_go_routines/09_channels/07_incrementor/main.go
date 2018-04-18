package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int { // 拉出器
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

Bar: 0
Bar: 1
Bar: 2
Foo: 0
Foo: 1
Bar: 3
Bar: 4
Bar: 5
Bar: 6
Bar: 7
Bar: 8
Bar: 9
Bar: 10
Bar: 11
Bar: 12
Bar: 13
Bar: 14
Bar: 15
Foo: 2
Foo: 3
Foo: 4
Foo: 5
Foo: 6
Foo: 7
Foo: 8
Foo: 9
Foo: 10
Bar: 16
Bar: 17
Bar: 18
Bar: 19
Foo: 11
Foo: 12
Foo: 13
Foo: 14
Foo: 15
Foo: 16
Foo: 17
Foo: 18
Foo: 19
Final Counter: 40
*/
