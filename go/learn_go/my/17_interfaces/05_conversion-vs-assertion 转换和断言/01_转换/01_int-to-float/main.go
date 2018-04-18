package main

import "fmt"

// int 转换成浮点
func int_to_float_conversion() {
	var x = 12
	var y = 12.130123
	fmt.Println(y + float64(x)) // 24.130122999999998
	// conversion: int to float64
}

func main() {
	int_to_float_conversion()
}
