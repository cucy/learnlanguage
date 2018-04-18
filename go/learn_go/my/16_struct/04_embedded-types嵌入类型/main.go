package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

type double_zero struct {
	person
	Lisense bool
}

func main() {
	p1 := double_zero{
		person: person{
			Fname: "张",
			Lname: "四",
			Age:   22,
		},
		Lisense: true,
	}

	p2 := double_zero{
		person: person{
			Fname: "王",
			Lname: "四",
			Age:   22,
		},
		Lisense: false,
	}

	fmt.Println(p1.Fname, p1.Lname, p1.Lisense, p1.Age)
	fmt.Println(p2.Fname, p2.Lname, p2.Lisense, p2.Age)
}

/*
张 四 true 22
王 四 false 22
*/
