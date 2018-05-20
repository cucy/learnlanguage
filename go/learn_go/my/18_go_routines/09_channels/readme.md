
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
