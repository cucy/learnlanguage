package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int { return len(p) }
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func sort_ex() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup) // [Zeno John Al Jenny]
	sort.Sort(studyGroup)
	fmt.Println(studyGroup) // [Al Jenny John Zeno]

}

func main() {
	sort_ex()
}
