package main

import "fmt"

func interface_string() {
	var name interface{} = "suse"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str) // string
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func interface_string_not_ok() {
	var name interface{} = 7
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n") // value is not a string
	}

}

func main() {
	interface_string_not_ok()
}
