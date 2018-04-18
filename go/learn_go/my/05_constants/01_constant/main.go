package main

import "fmt"

const p = "death & taxes 哥特摇滚"

// 一次定义多个值

const (
	pi       = 3.14
	language = "Go"
)

func main() {

	const q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("Pi - ", pi)
	fmt.Println("language - ", language)

}

// a CONSTANT is a simple unchanging value
// 常量是一个简单不变的值。
