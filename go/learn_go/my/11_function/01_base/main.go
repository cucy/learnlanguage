package main

import "fmt"

// func param-arg 函数参数
func greet(name string) {
	// greet is declared with a parameter
	// when calling greet, pass in an argument

	fmt.Println(name)
}

// two params 两个参数
func greet_with_2_params(first_name, last_name string) {
	fmt.Println(first_name, last_name)
}

// func return value 函数返回值
func greet_return_value(first_name, last_name string) string {
	return fmt.Sprint(first_name, last_name)
}

// func 命名返回值
func greet_return_naming(first_name, last_name string) (s string) {
	s = fmt.Sprint(first_name, last_name)
	return
}

// func 返回多个值
func greet_return_multipe(first_name, last_name string) (f, l string) {
	f = fmt.Sprint(first_name)
	l = fmt.Sprint(last_name)
	return
}

// func 可变参数
func greet_variadic_params_average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("数据类型:%T \n", sf) // 数据类型:[]float64
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// func 可变参数 <- 以args方式传入
func greet_variadic_params_average_get_args(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64((len(sf)))
}

// func slice param arg 切片参数
func greet_variadic_params_average_slice(sl []float64) float64 {
	total := 0.0
	for _, v := range sl {
		total += v
	}
	return total / float64(len(sl))
}

func main() { // main is the entry point to your program   main是你程序的入口点。

	greet("centos")
	greet("fly")
	greet_with_2_params("中国", "北京")

	fmt.Println(greet_return_value("中国1", "北京1")) // 中国1北京1

	fmt.Println(greet_return_naming("中国1命名返回值", "北京1命名返回值")) // 中国1命名返回值北京1命名返回值

	f_value, l_value := greet_return_multipe("中国1返回多个值", "北京1返回多个值")
	fmt.Printf("f_value:%v, l_value:%v", f_value, l_value) // f_value:中国1返回多个值, l_value:北京1返回多个值

	//	可变参数
	n_args := greet_variadic_params_average(11, 12.2, 99, 190)
	fmt.Println(n_args) // 78.05

	//	参数传入 (args ...)  自动分解
	args := []float64{100, 2345, 11, 1.0}
	args_func_return := greet_variadic_params_average_get_args(args...)
	fmt.Println(args_func_return)

	// func slice param arg 切片参数
	slice_data := []float64{1, 289, 11, 1.0}
	return_slice_data := greet_variadic_params_average_slice(slice_data)
	fmt.Println("return_slice_data: ", return_slice_data)

}
