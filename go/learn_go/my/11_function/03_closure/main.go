package main

import "fmt"

func ex1() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)

	}
	// fmt.Println(y) // y的作用域是在块里,外部无法引用
}

var x int

func increment() int {
	x++
	return x
	/*
		closure helps us limit the scope of variables used by multiple functions
		without closure, for two or more funcs to have access to the same variable,
		that variable would need to be package scope

		闭包帮助我们限制多个函数使用的变量的范围。
		没有闭包，两个或两个以上的函数可以访问同一个变量，
		该变量需要是包范围。

	*/
}

func ex2() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("ex2", increment()) // ex2 1
	fmt.Println("ex2", increment()) // ex2 2

	/*
	   closure helps us limit the scope of variables used by multiple functions
	   without closure, for two or more funcs to have access to the same variable,
	   that variable would need to be package scope

	   anonymous function
	   a function without a name

	   func expression
	   assigning a func to a variable

	   闭包帮助我们限制多个函数使用的变量的范围。
	   没有闭包，两个或两个以上的函数可以访问同一个变量，
	   该变量需要是包范围。

	   匿名函数
	   一个没有名字的函数。

	   函数表达式
	   将func赋给一个变量。
	*/

}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(x)           // 2

	ex2()

	increment_wrapper := wrapper()
	fmt.Println(increment_wrapper()) // 1
	fmt.Println(increment_wrapper()) // 2

	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println("A:", incrementA()) // A: 1
	fmt.Println("A:", incrementA()) // A: 2
	fmt.Println("B:", incrementB()) // B: 1
	fmt.Println("B:", incrementB()) // B: 2
	fmt.Println("B:", incrementB()) // B: 3
}
