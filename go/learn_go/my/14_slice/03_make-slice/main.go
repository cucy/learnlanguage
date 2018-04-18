package main

import "fmt"

func make_slice() {
	customer_number := make([]int, 3)
	// 3 is length & capacity
	// length - number of elements referred to by the slice 长度-切片所引用的元素个数。
	// capacity - number of elements in the underlying array 基础数组中元素的数量。
	customer_number[0] = 7
	customer_number[1] = 10
	customer_number[2] = 15

	fmt.Println(customer_number[0]) // 7
	fmt.Println(customer_number[1]) // 10
	fmt.Println(customer_number[2]) // 15

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2]) // dias!

}

func main() {
	make_slice()
}
