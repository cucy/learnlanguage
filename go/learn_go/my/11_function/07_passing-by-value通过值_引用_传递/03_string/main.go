package main

import "fmt"

func main() {

	name := "中国"
	fmt.Println(name) // 中国

	change_me(name)

	fmt.Println(name) // 中国
}

func change_me(z string) {

	fmt.Println(z) // 中国
	z = "北京"

	fmt.Println(z) // 北京

	//	因为是传值 所以修改并不会影响函数外的值
}
