package main

import (
	"fmt"
	"strconv"
)

func parse_Int() {
	//  ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")        // 转布尔
	f, _ := strconv.ParseFloat("3.1415", 64) // 转浮点
	i, _ := strconv.ParseInt("-42", 10, 64)  // 负数
	u, _ := strconv.ParseUint("42", 10, 64)  // 正数
	fmt.Println(b, f, i, u)                  // true 3.1415 -42 42

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z) // true 3.1415E+00 -2a 2a
}

func main() {
	parse_Int()
}
