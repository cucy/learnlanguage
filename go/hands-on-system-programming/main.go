package main

import (
	"fmt"
	"time"
)

func main() {
	var a int // passing value
	go func(v int) { fmt.Println(v) }(a)
	// passing pointer
	go func(v *int) { fmt.Println(*v) }(&a)
	a = 42
	time.Sleep(time.Nanosecond)
}
