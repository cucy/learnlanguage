package main

import "fmt"

// 字节切片转为 字符串
func rune_to_slice_of_bytes_to_string() {
	a_slice := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(a_slice)         // [104 101 108 108 111]
	fmt.Println(string(a_slice)) // hello
	// conversion: []bytes to string
	// we'll learn about []bytes soon
}

func main() {
	rune_to_slice_of_bytes_to_string()
}
