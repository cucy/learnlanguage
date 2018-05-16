package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("运行编译器:%v \n", runtime.Compiler)
	fmt.Printf("运行架构:%v \n", runtime.GOARCH)
	fmt.Printf("运行go版本:%v \n", runtime.Version())
	fmt.Printf("Number of Goroutines:%v \n", runtime.NumGoroutine())
}
