package main

import "fmt"

func true_and_false() {
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
}

//  取反
func bool_not() {
	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

}

// or
func bool_or() {
	if true || false {
		fmt.Println("This ran")
	}
}

// and
func bool_and() {

	if true && false {
		fmt.Println("This did not run")
	}

}

func main() {

	true_and_false()
	bool_not()
	bool_or()
	bool_and()

}
