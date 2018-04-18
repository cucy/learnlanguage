package main

import "fmt"

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
func main() {
	increment := wrapper()

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

闭包帮助我们限制多个函数使用的变量的范围。
没有闭包，两个或两个以上的函数可以访问同一个变量，
该变量需要是包范围。
*/
