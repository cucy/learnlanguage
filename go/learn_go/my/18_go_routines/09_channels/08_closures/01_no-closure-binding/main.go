package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting 在退出之前等待所有的goroutines完成。

	for _ = range values {
		<-done
	}

}

/*
Some confusion may arise when using closures with concurrency.

One might mistakenly expect to see a, b, c as the output.
What you'll probably see instead is c, c, c. This is because
each iteration of the loop uses the same instance of the variable v,
so each closure shares that single variable. When the closure runs,
it prints the value of v at the time fmt.Println is executed,
but v may have been modified since the goroutine was launched.
To help detect this and other problems before they happen,
run go vet.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines



在使用闭包并发性时可能会出现一些混淆。
人们可能会错误地期望看到a、b、c作为输出。
你可能会看到c c c，这是因为!
循环的每个迭代都使用变量v的相同实例，
每个闭包共享一个变量。当关闭运行时,
它在fmt时输出v的值。执行Println,
但是自从goroutine发射后，v可能已经被修改了。
为了在这些问题发生之前帮助检测这些问题，
跑去审查。
来源:
https://golang.org/doc/faq closures_and_goroutines
*/
