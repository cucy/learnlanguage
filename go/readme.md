# goland

# sync 
## waitgroup

```go 
// 接收指针类型的WaitGroup
func msg(i int, wg *sync.WaitGroup) {
	defer wg.Done() // 队列 -1
	fmt.Print(i, " ")
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 队列 + 1
		go func(num int) {
			msg(num, &wg) //
		}(i)

	}
	wg.Wait() // 等待队列长度为0 则,往后执行
	fmt.Println("完成")
}

// 1 6 4 5 7 8 3 0 9 2 完成
/*
只能保证所有的协程全部执行,但无法保证执行顺序
*/
```

## 互斥锁

加锁需要耗费时间

```go 

var mu sync.Mutex

func main() {

	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			mu.Lock()         // 上锁
			defer mu.Unlock() // 释放锁🔐
			m[i] = i
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Printf("map: %+v\n", m) // map: map[2:2 33:33 34:34 55:55 65:65 70:70 47:47 78:78 93:93 97:97 ...  ]
}
```


# 指针类型

```go 
func main() {
	msg := "hello"
	defer status(&msg) // 传入地址,类型为string

	fmt.Println(msg) // hello 
	msg = "goodbye"
}

// 接收string类型的指针
func status(msg *string) {
	// * 操作符为取值
	fmt.Println(*msg) // goodbye
}
```

# 时间time

`星期几`

```go 
fmt.Println(time.Now().Weekday())
```

## 计时器

```go 
func main() {
	done := make(chan struct{})
	tick := time.NewTicker(time.Millisecond * 500) // 定时器
	count := 0

	defer tick.Stop() // 最后停止计时

	go func() {
		time.Sleep(time.Second * 2)
		done <- struct{}{} // 2秒后往channel 传数据
	}()

OUTER:

	for {
		select {
		case <-tick.C:
			// 满足500毫秒则进入 此逻辑
			count++
			fmt.Println(count)
		case <-done:
			// 2秒后 channel中有数据, 退出循环
			fmt.Println("quitting")
			break OUTER
		}
	}

	fmt.Println("done")
}

```



# map

`find`

```go 
func main(){
	sizes := map[string]int{
		"hat": 11,
	}

	for _, item := range []string{"鞋子", "帽子", "hat"} {
		find_type, found := sizes[item]

		if !found {
			fmt.Printf("%s 在map中没有找到\n", item)
			continue
		}

		fmt.Printf("%s 在map中, 值为%d\n", item, find_type)
	}
}

/*
1
2
3
4
quitting
done
*/
```

# slice 

## copy 

```go 
func main() {
	list := []string{"牛奶", "🐆", "🥚鸡蛋"}
	fmt.Println(list)

	backup := make([]string, len(list))
	i := copy(backup, list)
	fmt.Printf("拷贝了%d个元素 \n", i) // 拷贝了3个元素
	fmt.Println(backup)          // [牛奶 🐆 🥚鸡蛋]
}
```

## 遍历

```go 
func main() {
	primes := []int{2, 3, 7, 11, 13, 17}
	for _, p := range primes {
		fmt.Println(p)
	}
}
```


# json 格式

## 解析成json

```go 
type Person struct {
	Name  string `json:"stu_name"`
	Phone string `json:"telphone"`
	// 需要大写,否则在json包中不可见
}

func main() {
	p := Person{Name: "zhansan", Phone: "1234"}
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("解析成json格式失败, marshal p: %s", err)
	}

	fmt.Printf("%+v\n", p) // {Name:zhansan Phone:1234}
	fmt.Printf("json output:%v\n", j) // json output:[123 34 115 116 117 95 110 97 109 ... ]
	fmt.Printf("json output:%s", string(j)) // json output:{"stu_name":"zhansan","telphone":"1234"}
}
```


# 协程

通常协程执行的前后顺序无法确定

```go 
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("sending %d\n", i)
		go func(i int) {
			fmt.Print(i, " ")// 0 8 9 2 1 3 6 5 4 7  
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\n done")
}

/*
sending 0
sending 1
sending 2
sending 3
sending 4
sending 5
sending 6
sending 7
sending 8
sending 9
0 8 9 2 1 3 6 5 4 7 
 done
*/
```
