package main

import "fmt"

func slicing_a_slice() {
	var results []int
	fmt.Println(results) // []

	my_slice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(my_slice)      // [a b c g m z]
	fmt.Println(my_slice[2:4]) // [c g] slicing a slice
	fmt.Println(my_slice[2])   // c index access; accessing by index
	fmt.Println("mystring"[2]) // 115 index access; accessing by index
}

func slicing_a_slice_1() {
	greeting := []string{
		"0 Good morning!",
		"1 Bonjour!",
		"2 dias!",
		"3 Bongiorno!",
		"4 Ohayo!",
		"5 Selamat pagi!",
		"6 Gutten morgen!",
	}

	fmt.Printf("[1:2]:%v\n", greeting[1:2]) // [1:2]:[1 Bonjour!]
	fmt.Printf("[:2]:%v\n", greeting[:2])   // [:2]:[0 Good morning! 1 Bonjour!]
	fmt.Printf("[5:]:%v\n", greeting[5:])   // [5:]:[5 Selamat pagi! 6 Gutten morgen!]
	fmt.Printf("[:]:%v\n", greeting[:])     // [:]:[0 Good morning! 1 Bonjour! 2 dias! 3 Bongiorno! 4 Ohayo! 5 Selamat pagi! 6 Gutten morgen!]

}

func main() {
	//slicing_a_slice()

	slicing_a_slice_1()
}
