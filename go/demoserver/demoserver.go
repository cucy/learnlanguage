package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

// 算术
type Arith int

// 参数
type Args struct {
	A, B int
}

// 商,除法
type Quotient struct {
	Quo, Rem int
}

// 函数必须导出(首字符大写)
// 必须有两个导出类型的参数
// 第一个参数是接收参数,
// 第二个参数是返回给客户端的参数
// 第二个参数必须是指针类型的
// 函数还要有一个返回值 error

func (t *Arith) Multiply(args *Args, reply *int) error {
	// Multiply: 乘法
	*reply = args.A * args.B
	fmt.Printf("这个Arith.Multiply方法被调用了, 参数args.A: %d,参数args.B: %d,返回值:%d \n", args.A, args.B, *reply)
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	//	 除法
	if args.B == 0 {
		return errors.New("分母不能为0")
	}
	// 商, 除法
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("这个方法执行了啊---嘿嘿--- Divide quo==", quo)
	return nil
}

/*
Go RPC 的函数只有符合四个条件才能够被远程访问，不然会被忽略
	函数必须是首字母大写（可以导出的）
	必须有两个导出类型的参数
	第一个参数是接受的参数，第二个参数是返回给客户端的参数，而且第二个参数是指针的类型
	函数还要有一个返回值error

func (t *T) MethodName(argType T1, replyType *T2) error


T、T1和T2类型必须能被encoding/gob包编解码。
*/
func main() {
	rpcDemo()
}
func rpcDemo() {
	arith := new(Arith)
	//arith=== 0xc04204e090
	fmt.Println("arith===", arith)

	rpc.Register(arith)
	// HandleHTTP将RPC消息的HTTP处理程序注册到Debug服务器
	// DEFAUTUPCPATH和Debug调试路径上的调试处理程序。
	// 仍然需要调用http.Services（），通常是在GO语句中。
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("err=====", err.Error())
	}
}
