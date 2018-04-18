package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) full_name() string {
	return p.fname + p.lname
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.full_name()) // JamesBond
	fmt.Println(p2.full_name()) // MissMoneypenny
}
