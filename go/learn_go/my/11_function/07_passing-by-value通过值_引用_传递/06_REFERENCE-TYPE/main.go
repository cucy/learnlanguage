package main

import "fmt"

func main() {
	m := make([]string, 1, 25) // 定义一个长度为1, 容量为25的字符串切片
	fmt.Println(len(m))        // 1
	fmt.Println(cap(m))        //25
	fmt.Println(m)             // []

	change_me(m)
	fmt.Println(m)            // [中国]
	fmt.Printf("&m %p\n", &m) // &m 0xc04204c3e0
}

func change_me(z []string) {
	fmt.Printf("before  &z %p\n", &z) // before  &z 0xc0420024a0
	z[0] = "中国"
	z = append(z, "北京")
	fmt.Println(z)                   // [中国 北京]
	fmt.Printf("after  &z %p\n", &z) // after  &z 0xc04204c440
}
