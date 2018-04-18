package main

import "fmt"

func main() {
	// 秩序问题 必须是先定义后引用
	fmt.Println(x) // 报错
	fmt.Println(y)

	x := 100
}

var y = 100
