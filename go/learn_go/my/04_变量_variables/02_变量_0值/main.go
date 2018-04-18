package main

import "fmt"

func main() {

	var a_int int
	var b_int64 int64
	var d_int32 rune
	var s_string string
	var f_float64 float64
	var b_bool bool
	var ptr_int *int

	fmt.Printf("a_int %v \n", a_int)
	fmt.Printf("b_int64 %v \n", b_int64)
	fmt.Printf("d_int32 %v \n", d_int32)
	fmt.Printf("s_string %v \n", s_string)
	fmt.Printf("f_float64 %v \n", f_float64)
	fmt.Printf("b_bool %v \n", b_bool)
	fmt.Printf("ptr_int %v \n", ptr_int)

	fmt.Println()

	// go 已定义变量而未初始化的变量都赋予0值,或者

	/*
		a_int 0
		b_int64 0
		d_int32 0
		s_string
		f_float64 0
		b_bool false
		ptr_int <nil>
	*/

}
