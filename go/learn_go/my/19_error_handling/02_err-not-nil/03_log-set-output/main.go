package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ... the Fatal functions call os.Exit(1) after writing the log message ... the Panic functions call panic after writing the log message.
*/

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

/*

/ *
包日志实现一个简单的日志包…写入标准错误并打印每个日志消息的日期和时间…在写入日志消息后，致命函数调用os.Exit(1)。在写入日志消息后，Panic函数调用Panic。
* /

// Println调用输出以打印到标准日志记录器。参数以fmt.Println的方式处理。

*/
