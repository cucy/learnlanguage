package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
	/*
		c
		a
		b

	*/
}

/*
To bind the current value of v to each closure as it is launched,
one must modify the inner loop to create a new variable each iteration.
One way is to pass the variable as an argument to the closure.

In this example, the value of v is passed as an argument to the anonymous function.
That value is then accessible inside the function as the variable u.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines


将v的当前值绑定到每个闭包，
我们必须修改内部循环来创建每个迭代的新变量。
一种方法是将变量作为参数传递给闭包。

在本例中，v的值作为参数传递给匿名函数。
这个值可以在函数中作为变量u访问。
*/
