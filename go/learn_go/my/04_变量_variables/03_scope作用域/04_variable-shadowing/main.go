package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // 49  max is now the result, not the function

}

// don't do this; bad coding practice to shadow variables
// 不这样做;对阴影变量进行糟糕的编码实践。
