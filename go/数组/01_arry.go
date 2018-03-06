package main

import "fmt"

func main() {
	arry := [...]int{1, 2, 3, 45}
	for i := 0; i < len(arry); i++ {

		fmt.Printf("数组下标为:%v 值为:%v  内存地址为:%v 所占长度:%v \n",
			i, arry[i], &arry[i], "未设定")
	}

	fmt.Printf("\n")

	for k, v := range arry {
		fmt.Printf("arry索引下标为:%v 值为%v\n",
			k, v)
	}

}
