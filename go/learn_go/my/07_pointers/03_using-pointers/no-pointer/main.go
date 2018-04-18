package main

import "fmt"

func zero(z int) {
	z = 0

	fmt.Printf("%p\n", &z) // address in func zero  0xc042052080
	fmt.Println(&z)        // address in func zero   0xc042052080
	fmt.Println(z)         // 0

}
func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main 0xc042052088
	fmt.Println(&x)        // address in main 0xc042052088

	zero(x)
	fmt.Println(x) // 依然是 5
}
