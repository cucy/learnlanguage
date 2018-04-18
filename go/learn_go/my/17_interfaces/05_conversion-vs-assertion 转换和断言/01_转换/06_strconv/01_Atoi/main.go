package main

import (
	"fmt"
	"strconv"
)

// 转换成int
func atoi_method() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x) // 转换成int
	fmt.Println(y + z)      // 18
}

func main() {
	atoi_method()
}
