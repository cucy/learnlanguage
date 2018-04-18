package main

import "fmt"

func factorial(x int) int {
	// 阶乘
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {

	v := factorial(3)
	fmt.Println(v)
}
