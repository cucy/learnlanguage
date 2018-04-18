package main

import "fmt"

func main() {
	for i := 10000; i < 100000; i ++ {
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}

}
