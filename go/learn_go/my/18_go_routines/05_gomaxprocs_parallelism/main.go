package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Foo: %v\n", i)
		time.Sleep(3 * time.Microsecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Bar: %v\n", i)
		time.Sleep(3 * time.Microsecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

/*
Foo: 0
Bar: 0
Bar: 1
Foo: 1
Foo: 2
Bar: 2
Bar: 3
Foo: 3
Bar: 4
Foo: 4
Foo: 5
Bar: 5
Foo: 6
Bar: 6
Foo: 7
Bar: 7
Bar: 8
Foo: 8
Bar: 9
Foo: 9

*/
