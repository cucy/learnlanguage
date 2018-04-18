package main

import "fmt"

type foo int

func bar() {
	var myAge foo
	myAge = 88
	fmt.Printf("%T %v \n", myAge, myAge)
	//	 main.foo 88

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)
	//	int 29

	// this doesn't work:
	//	 fmt.Println(myAge + yourAge)

	// this conversion works:
	//	 fmt.Println(int(myAge) + yourAge)
}

func main() {
	bar()
}
