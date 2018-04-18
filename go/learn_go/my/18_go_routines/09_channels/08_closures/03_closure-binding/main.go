package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

/*
Even easier is just to create a new variable,
using a declaration style that may seem odd but works fine in Go.
更简单的是创建一个新的变量，
使用一种声明风格，这看起来很奇怪，但效果很好。


SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
