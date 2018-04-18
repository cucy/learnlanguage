package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}

/*
YAYAYA Bob: 31
[YAYAYA Bob: 31 YAYAYA John: 42 YAYAYA Michael: 17 YAYAYA Jenny: 26]
[YAYAYA Michael: 17 YAYAYA Jenny: 26 YAYAYA Bob: 31 YAYAYA John: 42]
*/
// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

// String() string
// https://golang.org/doc/effective_go.html#printing
