# 字符串连接

```go
first_name := "xxx"
last_name := "yyy"

var stringBuilder bytes.Buffer

stringBuilder.WriterString(first_name)
stringBuilder.WriterString(last_name)

fmt.Pringln(stringBuilder.String)
```
