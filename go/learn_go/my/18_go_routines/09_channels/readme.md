
# channel

```go
make(chan Type)
make(chan Type,capacity)
```

```
channel <- value    // 阻塞发送
<- channel          // 接收并将其丢弃
x := <- channel     // 接收并保存
x,ok := <- channel  // 功能同上,同时检查通道是否已关闭 或者是否为空
```


```go 
// 创建int 类型的channel
	nums := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 往 channel里传数据
			nums <- i
		}
		// 关闭channel
		close(nums)
	}()

	for i := range nums {
		// 取数据
		print(i, " ")
	}
```
