package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(n)
	fmt.Println(n) // [1 2 3 4 5 6]
}
