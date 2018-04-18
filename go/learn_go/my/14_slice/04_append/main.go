package main

import "fmt"

// 无效的追加元素
func append_invalid() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham" // panic: runtime error: index out of range

	fmt.Println(greeting[2])
}
func append_method() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[3]) // Suprabadham

}

func append_beyond_capacity() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])   // gidday
	fmt.Println(len(greeting)) // 7
	fmt.Println(cap(greeting)) // 10
}

func append_slice_to_slice_by_int_type() {
	my_slice := []int{1, 2, 3, 4, 5}
	my_other_slice := []int{6, 7, 8, 9}
	my_slice = append(my_slice, my_other_slice...)
	fmt.Println(my_slice) // [1 2 3 4 5 6 7 8 9]
}

func append_slice_to_slice_by_string_type() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice) // [Monday Tuesday Wednesday Thursday Friday]

}

func main() {
	//append_invalid()

	//append_method()

	//append_beyond_capacity()
	//append_slice_to_slice_by_int_type()
	append_slice_to_slice_by_string_type()
}
