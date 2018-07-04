# 编译选项
```go
go build -ldflags "-s -w -H=windowsgui"
分别是：去符号表，去调试信息，去控制台
```


```

export GOOS=linux
export GOARCH=amd64

export GOPHERFACE_APP_ROOT=${GOPATH}/src/github.com/cucy/gopherface

go build -o  ${GOPHERFACE_APP_ROOT}/builds/gopherface-linux64
```

# 汇编分析
```go
go build -o ceshi /Users/zrd/Desktop/go/go_path/src/go_2018/2/hello_world2.go
go tool objdump -s "main\.main" ceshi
```

# 字符串转为整形

```go
x,_ := strconv.Atoi("12")
fmt.Printf("%#v %T",x,x) // 12 int
```
# 取得数值类型最大最小值

```go
	println(math.MinInt8, math.MaxInt8) // 	MaxInt8   = 1<<7 - 1 MinInt8   = -1 << 7
```

# 汇编分析

```
go tool compile -S main.go 
 
# 对象分布分析
go tool compile -m xxx.go
```

# 清除可执行文件测试文件

```
go  clean -x -i
```


# go调试

```
GODEBUG=gctrace=1 go run main.go 
```



# fmt
```

    // Go提供了几种打印格式，用来格式化一般的Go值，例如
    // 下面的%v打印了一个point结构体的对象的值
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // 如果所格式化的值是一个结构体对象，那么`%+v`的格式化输出
    // 将包括结构体的成员名称和值
    fmt.Printf("%+v\n", p)

    // `%#v`格式化输出将输出一个值的Go语法表示方式。
    fmt.Printf("%#v\n", p)

    // 使用`%T`来输出一个值的数据类型
    fmt.Printf("%T\n", p)

    // 格式化布尔型变量
    fmt.Printf("%t\n", true)

    // 有很多的方式可以格式化整型，使用`%d`是一种
    // 标准的以10进制来输出整型的方式
    fmt.Printf("%d\n", 123)

    // 这种方式输出整型的二进制表示方式
    fmt.Printf("%b\n", 14)

    // 这里打印出该整型数值所对应的字符
    fmt.Printf("%c\n", 33)

    // 使用`%x`输出一个值的16进制表示方式
    fmt.Printf("%x\n", 456)

    // 浮点型数值也有几种格式化方法。最基本的一种是`%f`
    fmt.Printf("%f\n", 78.9)

    // `%e`和`%E`使用科学计数法来输出整型
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // 使用`%s`输出基本的字符串
    fmt.Printf("%s\n", "\"string\"")

    // 输出像Go源码中那样带双引号的字符串，需使用`%q`
    fmt.Printf("%q\n", "\"string\"")

    // `%x`以16进制输出字符串，每个字符串的字节用两个字符输出
    fmt.Printf("%x\n", "hex this")

    // 使用`%p`输出一个指针的值
    fmt.Printf("%p\n", &p)

    // 当输出数字的时候，经常需要去控制输出的宽度和精度。
    // 可以使用一个位于%后面的数字来控制输出的宽度，默认
    // 情况下输出是右对齐的，左边加上空格
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // 你也可以指定浮点数的输出宽度，同时你还可以指定浮点数
    // 的输出精度
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // To left-justify, use the `-` flag.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // 你也可以指定输出字符串的宽度来保证它们输出对齐。默认
    // 情况下，输出是右对齐的
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // 为了使用左对齐你可以在宽度之前加上`-`号
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // `Printf`函数的输出是输出到命令行`os.Stdout`的，你
    // 可以用`Sprintf`来将格式化后的字符串赋值给一个变量
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // 你也可以使用`Fprintf`来将格式化后的值输出到`io.Writers`
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
```


## go get

```go
go  get -v -x  -u  github.com/gin-gonic/gin
```

## go   protobuf

https://segmentfault.com/a/1190000010098194


```go
go get -u github.com/golang/protobuf/protoc-gen-go

$ protoc --go_out=. ./proto/person.proto

$ protoc --go_out=plugins=grpc:.  ./proto/person.proto

$ protoc -I datafiles/ datafiles/transaction.proto --go_out=plugins=grpc:datafiles
```


```proto
syntax = "proto3";
package protofiles;


message Person {
    string name = 1;
    int32 id = 2; // Unique ID number for this person.
    string email = 3;
    enum PhoneType {
        MOBILE = 0;
        HOME = 1;
        WORK = 2;
    }
    message PhoneNumber {
        string number = 1;
        PhoneType type = 2;
    }
    repeated PhoneNumber phones = 4;
}

// Our address book file is just one of these.
message AddressBook {
    repeated Person people = 1;
}
```

`main.go`

```go
package main

import (
	pb "zrdfuncer/gin_restfulapi/protofiles"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "zhangsan",
		Email: "zhangsan@qq.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}

	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", p1)
}

```





## go dep

```
go get -u github.com/golang/dep/cmd/dep

dep  init

```
