package main

import "fmt"

// 浮点 转换成int
func floa_to_int_conversion() {
	var x = 12
	var y = 12.130123
	fmt.Println(int(y) + x) // 24
	// conversion: float64 to int
}

func main() {
	floa_to_int_conversion()
}
