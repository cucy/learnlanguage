package main

import "fmt"

func main() {

	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address -", &a)
	fmt.Printf("%d \n", &a)

	using_address()
}

const metersToYards float64 = 1.09361

// 使用地址
func using_address() {
	var meters float64
	fmt.Println("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is", yards, "yards.")
}
