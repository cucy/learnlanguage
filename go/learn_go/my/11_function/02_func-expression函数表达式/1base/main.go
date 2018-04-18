package main

import "fmt"

func greeting() {
	fmt.Println("Hello world!")
}

func main() {
	greeting()

	greeting1 := func() {
		fmt.Println("Hello world!")
	}

	greeting1()
	fmt.Printf("%T\n", greeting1) // func()

	// // another way func expression 另一种函数表达式 (函数返回函数)
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("函数返回函数 %T\n", greet) // 函数返回函数 func() string

}

// another way func expression 另一种函数表达式 (函数返回函数)
func makeGreeter() func() string {
	return func() string {
		return "函数返回函数 Hello world"
	}
}
