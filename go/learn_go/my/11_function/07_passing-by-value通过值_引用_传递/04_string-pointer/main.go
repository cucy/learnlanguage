package main

import "fmt"

func main() {
	name := "中国"
	fmt.Println(&name) // 0xc0420461c0

	change_me(&name)

	fmt.Println(&name) // 0xc0420461c0
	fmt.Println(name)  // 北京
}

func change_me(z *string) {
	fmt.Println(z)  // 0xc0420461c0
	fmt.Println(*z) // 中国

	*z = "北京"
	fmt.Println(z)  // 0xc0420461c0
	fmt.Println(*z) // 北京
}
