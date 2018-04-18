package main

import (
	"fmt"
	"sort"
)

func sort_ex() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s) // [Zeno John Al Jenny]

	//sort.StringSlice(s).Sort()

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s) // [Al Jenny John Zeno]
}

func main() {
	sort_ex()
}
