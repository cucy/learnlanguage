package main

import (
	"fmt"
	"sort"
)

func sort_Strings_ex() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s) // [Al Jenny John Zeno]
}

func main() {
	sort_Strings_ex()
}
