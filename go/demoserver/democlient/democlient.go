package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {

	serverAddress := "127.0.0.1:1234"
	fmt.Println("severAddress==", serverAddress)
	//DelayHTTP在指定的网络地址连接到HTTP RPC服务器

	// /在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("发生错误了 在这里地方  DialHTTP", err)
	}

	args := ArgsTwo{100, 3}

	var reply int
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Call Multiply  发生错误了哦   arith error:", err)
	}
	fmt.Printf("Arith 乘法: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientTwo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith 除法取整数: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
