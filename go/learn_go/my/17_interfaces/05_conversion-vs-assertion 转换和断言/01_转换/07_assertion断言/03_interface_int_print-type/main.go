package main

import "fmt"

func interface_int_print_type() {
	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
}

//mismatched types interface {} and int 不匹配的类型接口{}和int。 (报错)
func mismatched_types_interface() {
	var val interface{} = 7
	//fmt.Println(val + 6)  // mismatched types interface {} and int 不匹配的类型接口{}和int。 (报错)
	fmt.Println(val)
}

// interface int sum
func interface_int_sum() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6) // 13
}

// casting reminder
func casting_reminder() {
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int
}

// 08_interface-cast-error_need-type-assertion 接口转换错误需要类型断言。
func interface_cast_error_need_type_assertion() {
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int

	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
	//fmt.Printf("%T\n", int(val))
	fmt.Printf("%T\n", val.(int)) // int
}

func main() {
	interface_cast_error_need_type_assertion()

}
