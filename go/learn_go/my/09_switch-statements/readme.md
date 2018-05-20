
# switch

```go
swicth option_statement;optionExpression{

case optionList1:
  block1
 
 
...

case optionListN:
  blockN


default:
  blockD

}

```


# ex 

```go

func BoundedUbt(min, max, val int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val

	}

}

```

# ex2

```go
	switch suffix := Suffix(file);suffix { // 原始的非经典用法
	case ".gz":
		return GizpFileList(file)
	case ".tar":
		fallthrough
	case ".tar.gz":
		fallthrough // 交给下一个函数执行
	case ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	
	}
 // ------------
 
 	switch Suffix(file) { // 经典用法
	case ".gz":
		return GizpFileList(file)
	case ".tar", ".tar.gz", ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)

	}

```

# 类型开关

```go
func classfier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case int8, int16, int32, int64, int:
			fmt.Printf("param #%d is an int\n", i)

		case nil:
			fmt.Printf("param #%d is nil \n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		case stu:
			fmt.Printf("param #%d is stu\n", i)

		case *stu:
			fmt.Printf("param #%d is  stu addr\n", i)

		default:
			fmt.Printf("param #%d 未知数据类型\n", i)

		}
	}

}

type stu struct {
	name string
}
```
