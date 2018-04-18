package main

import "fmt"

func no_intreface_error_invalid_code() {
	// 错误
	name := "suse"
	//	str, ok := name.(string) //  invalid type assertion: name.(string) (non-interface type string on left)
	str, ok := 1, 2 //  invalid type assertion: name.(string) (non-interface type string on left)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func main() {
	//no_intreface_error_invalid_code()
}
