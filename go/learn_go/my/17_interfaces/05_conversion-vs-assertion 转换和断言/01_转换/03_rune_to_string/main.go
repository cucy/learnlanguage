package main

import "fmt"

// int32(rune) to string 转换
func rune_to_string_conversion() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement rune是int32的别名;在这个语句中通常省略。
	var y int32 = 'b'
	fmt.Println(x)         // 97
	fmt.Println(y)         // 98
	fmt.Println(string(x)) // a
	fmt.Println(string(y)) // b
	// conversion: rune to string
}

func main() {
	rune_to_string_conversion()
}
