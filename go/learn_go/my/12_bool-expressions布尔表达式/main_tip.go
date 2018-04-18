package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func half_float(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

// 可变最大值
func max_num(numbers ...int) int {
	var largest int
	/*
		// 负数的是时候 -11 永远小于0
		for _, v := range numbers {
			if v > largest  {
				largest = v
			}
		}

	*/

	for i, v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}

// 布尔表达式
func bool_expression() {
	fmt.Println(
		(true && false) || (false && true) || !(false && false)) // true
}

// params and args 参数和可变参数
func params_and_args(numbers ...int) {
	fmt.Println(numbers)
}

func main() {
	h, even := half(5)
	fmt.Println(h, even) // 2 false

	//
	h1, even1 := half_float(5)
	fmt.Println(h1, even1) // 2.5 false

	//	 函数表达式
	half_func_expression := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(half_func_expression(5)) // 2 false

	//	max_num
	m_value := max_num(234, 11, 45, 3232, 12345643)
	fmt.Println(m_value) // 12345643

	//	 布尔表达式
	bool_expression() // true

	//	 参数和可变参数
	params_and_args(1, 2)
	params_and_args(1, 2, 3)
	a_slice := []int{12, 21, 324, 5435}
	params_and_args(a_slice...)
	params_and_args() // []
}
