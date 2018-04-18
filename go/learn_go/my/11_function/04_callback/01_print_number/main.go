package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// callback: passing a func as an argument
//回调:通过func作为参数。

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs

}

func main() {

	visit([]int{1, 2, 3, 4}, func(n int) { fmt.Println(n) })

	//	过滤
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool { return n > 1 })
	fmt.Println("filter", xs) //  filter [2 3 4]
}
