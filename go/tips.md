
```go
go build -o ceshi /Users/zrd/Desktop/go/go_path/src/go_2018/2/hello_world2.go
go tool objdump -s "main\.main" ceshi

```

# 字符串转为整形

```go
x,_ := strconv.Atoi("12")
fmt.Printf("%#v %T",x,x) // 12 int
```
