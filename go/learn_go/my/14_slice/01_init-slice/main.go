package main

import "fmt"

func init_slice() {
	my_slice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", my_slice) // []int
	fmt.Println(my_slice, "\n")
}
func init_slice_by_make() {
	my_slice := make([]int, 0, 1)
	fmt.Println("\n-----------------")
	fmt.Println(my_slice)      // []
	fmt.Println(len(my_slice)) // 0
	fmt.Println(cap(my_slice)) // 1
	fmt.Println("-----------------")

	for i := 0; i < 10; i++ {
		my_slice = append(my_slice, i)
		fmt.Printf("Len:%d, Capacity:%d, Value:%d %d\n", len(my_slice), cap(my_slice), my_slice[i], my_slice)
	}
	/*
		Len:1, Capacity:1, Value:0 [0]
		Len:2, Capacity:2, Value:1 [0 1]
		Len:3, Capacity:4, Value:2 [0 1 2]
		Len:4, Capacity:4, Value:3 [0 1 2 3]
		Len:5, Capacity:8, Value:4 [0 1 2 3 4]
		Len:6, Capacity:8, Value:5 [0 1 2 3 4 5]
		Len:7, Capacity:8, Value:6 [0 1 2 3 4 5 6]
		Len:8, Capacity:8, Value:7 [0 1 2 3 4 5 6 7]
		Len:9, Capacity:16, Value:8 [0 1 2 3 4 5 6 7 8]
		Len:10, Capacity:16, Value:9 [0 1 2 3 4 5 6 7 8 9]

	*/

}

func range_slice() {
	xs := []int{1, 3, 5, 7, 9, 11}

	for i, value := range xs {
		fmt.Printf("index:%d value:%d\n", i, value)
	}
}

// 字符串切片

func string_slice() {
	greeting := []string{"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!"}
	for i, curren_entry := range greeting {
		fmt.Println(i, curren_entry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}

func main() {

	init_slice()

	range_slice()

	init_slice_by_make()
}
