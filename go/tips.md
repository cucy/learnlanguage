
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
