package main

import "fmt"

func main() {
	x := 42 // 只在main函数内可访问

	fmt.Println(x)

	foo()
}

func foo() {
	// 无法通过编译, 引用未定义的变量
	//
	fmt.Println(x)
}
