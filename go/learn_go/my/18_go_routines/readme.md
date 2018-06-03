# waitgroup

```go 
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 每创建一个协程, 就把任务队列中的任务数加1
		go hello(&wg)
	}

	wg.Wait() // .Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成, 则队列数减1
	fmt.Println("hello")
}
```

