package main

import (
	"fmt"
	"strconv"
)

func Itoa_int_to_string() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y) // I have this many: 12

}

func main() {
	Itoa_int_to_string()
}
