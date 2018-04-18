package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Microsecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(3 * time.Microsecond)
	}
	wg.Done()
}

/*
Bar: 0
Foo: 0
Bar: 1
Foo: 1
Bar: 2
Foo: 2
Foo: 3
Bar: 3
Bar: 4
Foo: 4
Bar: 5
Foo: 5
Foo: 6
Bar: 6
Foo: 7
Bar: 7
Bar: 8
Foo: 8
Foo: 9
Bar: 9


*/
