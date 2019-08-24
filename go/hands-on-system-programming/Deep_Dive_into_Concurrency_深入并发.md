
# 了解Goroutines

Go是一种以并发性为中心的语言，其中两个主要特性通道和Goroutine是内置包的一部分。现在我们将看到它们是如何工作的，以及它们的基本功能是什么，从Goroutines开始，这使得可以同时执行应用程序的部分。


# 比较线程和goroutine

Goroutines是用于并发的原语之一，但是它们与线程有什么不同呢？让我们来看看他们每个人。


# Threads

Current OSes are built for modern architectures that have processors with more than one core per CPU, or use technologies, such as hyper-threading, that allow a single core to support more than one thread. Threads are parts of processes that can be managed by the OS scheduler, which can assign them to a specific core/CPU. Like processes, threads carry information about the application execution, but the size of such information is smaller than processes. This includes the current instruction in the program, a stack of the current execution, and the variables needed for it.

The OS is already responsible for the context switch between processes; it saves the older process information and loads the new process information. This is called a process context switch and it's a very costly operation, even more than process execution. 

In order to jump from one thread to another, the same operation can be done between threads. This is called a thread context switch and it is also a heavy operation—even if it is not as hefty as the process switch—because a thread carries less information than a process. 

```
螺纹
当前的操作系统是为现代体系结构而构建的，这些现代体系结构的处理器每个CPU都有一个以上的核心，或者使用允许单个核心支持多个线程的技术，如超线程。线程是可以由OS调度程序管理的进程的一部分，它可以将线程分配给特定的核心/CPU。与进程一样，线程携带有关应用程序执行的信息，但这些信息的大小小于进程。这包括程序中的当前指令、当前执行的堆栈以及所需的变量。
操作系统已经负责进程之间的上下文切换；它保存旧的进程信息并加载新的进程信息。这被称为进程上下文切换，它是一个非常昂贵的操作，甚至比进程执行还要昂贵。
为了从一个线程跳到另一个线程，可以在线程之间执行相同的操作。这被称为线程上下文切换，而且它也是一个繁重的操作，即使它不如进程切换那么重，因为线程的信息比进程少。
```

# Goroutines

Threads have a minimum size in memory; usually, it is in the order of MBs (2 MB for Linux). The minimum size sets some limitations on the application creation of a new thread—if each thread is at least some MBs, 1,000 threads will occupy at least a few GBs of memory. The way that Go tackles these issues is through the use of a construct similar to threads, but this is handled by the language runtime instead of the OS. The size of a goroutine in memory is three orders of magnitude (2 KB per goroutine), meaning that the minimum memory usage of 1,000 goroutines is comparable to the memory usage of a single thread.
  
This is obtained by defining what data the goroutines are retaining internally, using a data structure called g that describes the goroutine information, such as stack and status. This is an unexported data type in the runtime package and it can be found in the Go source code. Go keeps a track of OSes using another data structure from the same package called m. The logical processors that are acquired in order to execute a goroutine are stored in p structures. This can be verified in the Go runtime package documentation:

```
戈罗提内斯
线程在内存中的大小最小；通常按mbs（Linux为2 MB）的顺序排列。最小大小对应用程序创建新线程设置了一些限制。如果每个线程至少有一些MB，1000个线程将占用至少几GB的内存。解决这些问题的方法是使用类似于线程的构造，但这是由语言运行时而不是操作系统来处理的。内存中Goroutine的大小是三个数量级（每个Goroutine 2 KB），这意味着1000个Goroutine的最小内存使用量与单个线程的内存使用量相当。
这是通过定义Gorooutine在内部保留的数据来实现的，使用一个名为G的数据结构来描述Gorooutine信息，例如堆栈和状态。这是运行时包中的未分析数据类型，可以在go源代码中找到。Go使用同一个名为m的包中的另一个数据结构跟踪OSE。为执行Goroutine而获取的逻辑处理器存储在p结构中。这可以在Go运行时包文档中验证：
```

These three entities interact as follows—for each goroutine, a new g gets created, g is queued into p, and each p tries to acquire m to execute the code from g. There are some operations that will block the execution, such as these:

- Built-in synchronization (channels and the sync package)
- System calls that are blocking, such as file operations
- Network operations

When these kinds of operations happen, the runtime detaches p from m and uses (or creates, if it does not already exist) another dedicated m for the blocking operation. The thread becomes idle after executing such operations.

```
这三个实体按如下方式交互：为每个goroutine创建一个新的g，g排队进入p，每个p尝试获取m以从g执行代码。有一些操作会阻止执行，例如：
-内置同步（通道和同步包）
-正在阻塞的系统调用，例如文件操作
-网络操作
当这些类型的操作发生时，运行时将p从m中分离出来，并使用（或创建，如果它不存在的话）另一个专用m来执行阻塞操作。执行此类操作后，线程将变为空闲。
```

# New goroutine

Goroutines are one of the best examples of how Go hides complexity behind a simple interface. When writing an application in order to launch a goroutine, all that is needed is to execute a function preceded by the go keyword:

``` 
Goroutines是如何将复杂性隐藏在简单接口后面的最佳示例之一。当编写应用程序以启动Goroutine时，只需要执行一个前面有go关键字的函数：

func main() {
	 go fmt.Println("Hello, playground")
}
```

If we run the application of the previous example, we will see that it does not produce any output. Why? In Go, the application terminates when the main goroutine does, and it looks like this is the case. What happens is that the Go statements create the goroutine with the respective runtime.g, but this has to be picked up by the Go scheduler, and this does not happen because the program terminates right after the goroutine has been instantiated.

```
如果我们运行前一个示例的应用程序，我们将看到它不会产生任何输出。为什么？在Go中，当主Goroutine执行此操作时，应用程序将终止，情况看起来是这样的。实际情况是，go语句使用各自的runtime.g创建goroutine，但这必须由go调度程序执行，而这不会发生，因为程序在goroutine实例化后立即终止。
```

Using the time.Sleep function to let the main goroutine wait (even a nanosecond!) is enough to let the scheduler pick up the goroutine and execute its code. This is shown in the following code:

```
使用time.sleep函数让主Goroutine等待（甚至一纳秒！）足够让调度程序获取Goroutine并执行其代码。如下代码所示：
func main() {
	go fmt.Println("Hello, playground")
	time.Sleep(time.Nanosecond)
}
```


We already saw that Go methods also count as functions, which is why they can be executed concurrently the with go statement, as they were normal functions:

```
我们已经看到go方法也算作函数，这就是为什么它们可以与with go语句同时执行，因为它们是正常函数：

type a struct{}

func (a) Method() { fmt.Println("Hello, playground") }

func main() {
    go a{}.Method()
    time.Sleep(time.Nanosecond)
}
```

Closures are anonymous functions, so they can be used as well, which is actually a very common practice:

```
闭包是匿名函数，因此也可以使用它们，这实际上是非常常见的做法：

func main() {
    go func() {
        fmt.Println("Hello, playground")
    }()
    time.Sleep(time.Nanosecond)
}
```


# Multiple goroutines

Organizing code in multiple goroutines can be helpful to split the work between processors and has many other advantages,
 as we will see in the next chapters. Since they are so lightweight, we can create a number of goroutines very easily using loops:

```
在多个Goroutine中组织代码有助于在处理器之间拆分工作，并具有许多其他优势，如我们将在下一章中看到的。由于它们非常轻，因此我们可以很容易地使用循环创建多个Goroutine：


func main() {
    for i := 0; i < 10; i++ {
        go fmt.Println(i)
    }
    time.Sleep(time.Nanosecond)
}
```

# Argument evaluation
  
If we change this example slightly by using a closure without arguments, we will see a very different result:

```
如果我们通过使用一个没有参数的闭包稍微改变这个例子，我们将看到一个非常不同的结果：

func main() {
    for i := 0; i < 10; i++ {
         go func() { fmt.Println(i) }()
    }
    time.Sleep(time.Nanosecond)
}
```

If we run this program, we can see that the Go compiler issues a warning in the loop: loop variable i captured by func literal.

The variable in the loop gets referenced in the function we defined—the creation loop of the goroutines is quicker than goroutines executing, and the result is that the loop finishes before a single goroutine is started, resulting in the print of the value of the loop variable after the last iteration.

In order to avoid the error of the captured loop variable, it's better to pass the same variable as an argument to the closure. The arguments of the goroutine function are evaluated upon creation, meaning that changes to that variable will not be reflected inside the goroutine, unless you are passing a reference to a value such as a pointer, map, slice, channel, or function. We can see this difference by running the following example:


```
如果我们运行这个程序，我们可以看到go编译器在循环中发出警告：循环变量i被func-literal捕获。
循环中的变量在我们定义的函数中被引用。Gorooutines的创建循环比Gorooutines的执行速度快，结果是循环在单个Gorooutine启动之前完成，从而在最后一个iter之后打印循环变量的值。。
为了避免捕获循环变量的错误，最好将同一变量作为参数传递给闭包。Goroutine函数的参数在创建时进行评估，这意味着对该变量的更改不会反映在Goroutine中，除非传递对某个值（如指针、映射、切片、通道或函数）的引用。通过运行以下示例，我们可以看到这种差异：

func main() {
    var a int
    // passing value
    go func(v int) { fmt.Println(v) }(a)
  
    // passing pointer
    go func(v *int) { fmt.Println(*v) }(&a)
  
    a = 42
    time.Sleep(time.Nanosecond)
}
```


# Synchronization 同步

Goroutines allow code to be executed concurrently, but the synchronization between values is not ensured out of the box. We can check out what happens when trying to use a variable concurrently with the following example:

```
Goroutines允许同时执行代码，但不能确保值之间的同步。我们可以通过下面的示例来检查在同时使用变量时会发生什么情况：

func main() {
	var i int
	go func(i *int) {
		for j := 0; j < 20; j++ {
			time.Sleep(time.Millisecond)
			fmt.Println("in: ",*i, j)
		}
	}(&i)
	for i = 0; i < 20; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
}


in:  0 0
0
1
in:  2 1
in:  2 2
2
3
in:  3 3
4
in:  4 4
in:  5 5
...
```

We have an integer variable that changes in the main routine—doing a millisecond pause between each operation—and after the change, the value is printed.

In another goroutine, there is a similar loop (using another variable) and another print statement that compares the two values. Considering that the pauses are the same, we would expect to see the same values, but this is not the case. We see that sometimes, the two goroutines are out of sync.

The changes are not reflected immediately because the memory is not synchronized instantaneously. We will learn how to ensure data synchronization in the next chapter.

``` 
我们有一个整型变量，它在主程序中改变，在每次操作之间做一个毫秒的停顿，在改变之后，值被打印出来。
在另一个goroutine中，有一个类似的循环（使用另一个变量）和另一个print语句来比较这两个值。考虑到停顿是相同的，我们希望看到相同的值，但事实并非如此。我们看到，有时两个Goroutine是不同步的。
更改不会立即反映出来，因为内存没有立即同步。我们将在下一章中学习如何确保数据同步。
```

# Exploring channels

Channels are a concept that is unique to Go and a few other programming languages. Channels are very powerful tools that allow a simple method for synchronizing different goroutines, which is one of the ways we can solve the issue raised by the previous example.

```
Channels是Go和其他一些编程语言所独有的概念。通道是非常强大的工具，它允许使用一种简单的方法来同步不同的goroutine，这是我们解决前一个示例提出的问题的方法之一。
```


# Properties and operations

A channel is a built-in type in Go that is typed as arrays, slices, and maps. It is presented in the form of chan type and initialized by the make function.

```
属性和操作
通道是内置的go类型，它被类型化为数组、切片和映射。它以chan类型的形式呈现，并由make函数初始化。
```

# Capacity and size

As well as the type that is traveling through the channel, there is another property that the channel has: its capacity. This represents the number of items a channel can hold before any new attempt to send an item is made, resulting in a blocking operation. The capacity of the channel is decided at its creation and its default value is 0:

```
除了通过通道的类型外，通道还具有另一个属性：容量。这表示在进行任何新的发送项目尝试之前，通道可以保留的项目数，从而导致阻塞操作。通道的容量在创建时确定，其默认值为0:
```

// channel with implicit zero capacity
var a = make(chan int)

// channel with explicit zero capacity
var a = make(chan int, 0)

// channel with explicit capacity
var a = make(chan int, 10)

```go
//具有隐式零容量的通道
var a=品牌（chan int）

//具有显式零容量的通道
var a=制造（chan int，0）

//具有显式容量的通道
var a=品牌（chan int，10）
```


The capacity of the channel cannot be changed after its creation and can be read at any time using the built-in cap function:

```
通道的容量在创建后无法更改，可以使用内置的cap功能随时读取：

func main() {
    var (
        a = make(chan int, 0)
        b = make(chan int, 5)
    )

    fmt.Println("a is", cap(a))
    fmt.Println("b is", cap(b))
}
```


The len function, when used on a channel, tells us the number of elements that are held by the channel:

```
当在通道上使用len函数时，它告诉我们通道所包含的元素数：

func main() {
	var (
		a = make(chan int, 5)
	)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}
}

a is 1 / 5
a is 2 / 5
a is 3 / 5
a is 4 / 5
a is 5 / 5

```



# Blocking operations

If a channel is full or its capacity is 0, then the operation will block. If we take the last example, which fills the channel and tries to execute another send operation, our application gets stuck:

```
如果通道已满或其容量为0，则操作将被阻止。如果我们以最后一个例子为例，它填充通道并尝试执行另一个发送操作，那么我们的应用程序会卡住：

func main() {
	var a = make(chan int, 5)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}
	a <- 0 // 阻塞,死锁
}
```

When all goroutines are locked (in this specific case, we only have the main goroutine), the Go runtime raises a deadlock—a fatal error that terminates the execution of the application:

```
当所有Goroutine都被锁定时（在这种特定情况下，我们只有主Goroutine），Go运行时会引发死锁，这是一个致命错误，它会终止应用程序的执行：
```

This is can happen with both receive or send operations, and it's the symptom of an error in the application design. Let's take the following example:

```
这在接收或发送操作中都会发生，这是应用程序设计中出现错误的症状。让我们举个例子：

func main() {
	var a = make(chan int)
	a <- 10
	fmt.Println("a is", <-a, "/", cap(a))
}

// fatal error: all goroutines are asleep - deadlock!

```


In the previous example, there is the a <- 10 send operation and the matching <-a receive operation, but nevertheless, it results in a deadlock. However, the channel we created has no capacity, so the first send operation will block. We can intervene here in two ways:

By increasing the capacity: This is a pretty easy solution that involves initializing the channel with make(chan int, 1). It works best only if the number of receivers is known a priori; if it is higher than the capacity, then the problem appears again.
By making the operations concurrent: This is a far better approach because it uses the channels for what they made for—concurrency.

```
在前面的示例中，有一个<-10发送操作和一个匹配的<-a接收操作，但是它会导致死锁。但是，我们创建的通道没有容量，因此第一个发送操作将阻塞。我们可以通过两种方式进行干预：
通过增加容量：这是一个非常简单的解决方案，涉及使用make初始化通道（chan int，1）。只有事先知道接收器的数量，它才能发挥最佳效果；如果它高于容量，则问题再次出现。
通过使操作并发：这是一种更好的方法，因为它使用通道来实现并发。

func main() {
    var a = make(chan int)
    go func() {
        a <- 10
    }()
    fmt.Println(<-a)
}

```

Now, we can see that there are no deadlocks here and the program prints the values correctly. Using the capacity approach will also make it work, but it will be tailored to the fact that we are sending a single message, while the other method will allow us to send any number of messages through the channel and receive them accordingly from the other side:

```
现在，我们可以看到这里没有死锁，程序正确地打印了值。使用容量方法也可以使其工作，但它将根据我们发送单个消息的事实进行调整，而另一种方法将允许我们通过通道发送任意数量的消息，并相应地从另一方接收这些消息：

func main() {
    const max = 10
    var a = make(chan int)

    go func() {
        for i := 0; i < max; i++ {
            a <- i
        }
    }()
    for i := 0; i < max; i++ {
        fmt.Println(<-a)
    }
}
```

We now have a constant to store the number of operations executed, but there is a better and more idiomatic way to let a receiver know when there are no more messages. We will cover this in the next chapter about synchronization.

```
我们现在有一个常量来存储执行的操作数，但是有一种更好、更惯用的方法来让接收者知道何时没有更多的消息。我们将在下一章讨论同步。
```


# Closing channels

The best way of handling the end of a synchronization between a sender and a receiver is the close operation. This function is normally executed by the sender because the receiver can verify whether the channel is still open each time it gets a value using a second variable:

``` 
处理发送方和接收方之间同步结束的最佳方法是关闭操作。此函数通常由发送方执行，因为接收方可以使用第二个变量验证每次获取值时通道是否仍然打开：

value,ok:= <-ch
```

The second receiver is a Boolean that will be true if the channel is still open, and false otherwise. When a receive operation is done on a close channel, the second received variable will have the false value, and the first one will have the 0 value of the channel type, such as these:

0 for numbers
false for Booleans
"" for strings
nil for slices, maps, or pointers


```
第二个接收器是一个布尔值，如果通道仍然打开，则为真，否则为假。当在关闭通道上执行接收操作时，第二个接收变量将具有假值，第一个接收变量将具有通道类型的0值，例如：
0代表数字
布尔值为假
对于字符串
切片、映射或指针为零


```

The example of sending multiple messages can be rewritten using the close function, without having prior knowledge of how many messages will be sent:

```
发送多条消息的示例可以使用close函数重写，而不必事先知道将发送多少条消息：

func main() {
	const max = 10
	var a = make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			a <- i
		}
		close(a)
	}()

	for {
		v, ok := <-a
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
```


There is a more synthetic and elegant way to receive a message from a channel until it's closed: by using the same keyword that we used to iterate maps, arrays, and slices. This is done through range:

```
有一种更为综合和优雅的方式来接收来自通道的消息，直到它关闭：使用我们用于迭代映射、数组和切片的相同关键字。这是在整个范围内完成的：

for v := range a {
    fmt.Println(v)
}
```


# One-way channels

Another possibility when handling channel variables is specifying whether they are only for sending or only for receiving data. This is indicated by the <- arrow, which will precede chan if it's just for receiving, or follow it if it's just for sending:

```
处理通道变量时的另一种可能性是指定它们是仅用于发送还是仅用于接收数据。这是由<-箭头指示的，如果只是接收，它将位于chan之前，如果只是发送，则跟随它：

func main() {
	var a = make(chan int)
	s, r := (chan<- int )(a), (<-chan int )(a)
	fmt.Printf("%T  %T", s, r)
}

// chan<- int  <-chan int
```

Channels are already pointers, so casting one of them to its send-only or receive-only version will return the same channel, but will reduce the number of operations that can be performed on it. The types of channels are as follows:

Send only channels, chan<-, which will allow you to send items, close the channel, and prevent you from sending data with a compile error.
Receive only channel, <-chan, that will allow you to receive data, and any send or close operations will be compiling errors.
When a function argument is a send/receive channel, the conversion is implicit and it is a good practice to adopt because it prevents mistakes such as closing the channel from the receiver. We can take the other example and make use of the one-way channels with some refactoring.

```
通道已经是指针，因此将其中一个通道强制转换为其只发送或只接收版本将返回相同的通道，但将减少可以对其执行的操作数。信道类型如下：
只发送频道，chan<-，这将允许您发送项目，关闭频道，并阻止您发送带有编译错误的数据。
只接收通道<-chan，允许您接收数据，任何发送或关闭操作都将编译错误。
当函数参数是发送/接收通道时，转换是隐式的，采用转换是一个很好的实践，因为它可以防止错误，例如从接收器关闭通道。我们可以举另一个例子，利用单向通道进行重构。
```

We can also create a function for sending values that uses a send-only channel:

```
我们还可以创建一个函数来发送使用只发送通道的值：

func send(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}
```

```
Do the same thing for receiving using a receive-only channel:

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
```

And then, use them with the same channel that will be automatically converted in the one-way version:

然后，将它们与将在单向版本中自动转换的相同通道一起使用：

```
func main() {
    var a = make(chan int)

    go send(a, 10)
  
    receive(a)
}
```


# Waiting receiver


Most of the examples we saw in the previous section had the sending operations done in a goroutine, and had the receiving operations done in the main goroutine. It could be the case that all operations are handled by goroutines, so do we synchronize the main one with the others?

A typical technique is the use of another channel used for the sole purpose of signaling that a goroutine has finished its job. The receiving goroutine knows that there are no more messages to get with the closure of the communication channel and it closes another channel that is shared with the main goroutine after finishing its operation. The main function can wait for the closure of the channel before exiting.

The typical channel that is used for this scope does not carry any additional information except for whether it is open or closed, so it is usually a chan struct{} channel. This is because an empty data structure has no size in memory. We can see this pattern in action by making some changes to the previous example, starting with the receiver function:

```
我们在上一节中看到的大多数示例都是在goroutine中执行发送操作，而在主goroutine中执行接收操作。可能是所有操作都由goroutines处理，那么我们是否将主操作与其他操作同步？

一种典型的技术是使用另一个通道，唯一的目的是通知Gorotine已经完成了它的工作。接收Goroutine知道通信通道关闭后没有更多的消息可获取，并且在完成操作后关闭与主Goroutine共享的另一个通道。主功能可以在退出前等待通道关闭。

用于此作用域的典型通道不包含任何附加信息，除了它是打开的还是关闭的，因此它通常是chan结构通道。这是因为空数据结构在内存中没有大小。我们可以通过对前一个示例进行一些更改，从receiver函数开始，看到这个模式正在运行：

func receive(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		println(v)
	}
	close(done)
}

```

The receiver function gets an extra argument—the channel. This is used to signal that the sender is done and the main function will use that channel to wait for the receiver to end its task

```
receiver函数在通道中获得一个额外的参数。这用于表示发送方已完成，主功能将使用该通道等待接收方结束其任务。

func send(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println(v)
	}
	close(done)
}

func main() {
	a := make(chan int)
	go send(a, 10)
	done := make(chan struct{})
	go receive(a, done)
	<-done
}
```


# Special values

There are a couple of situations in which channels behave differently. We will now see what happens when a channel is set to its zero value—nil—or when it is already closed.
  
```

特殊值
在一些情况下，通道的行为会有所不同。现在我们将看到当一个通道设置为零值nil或当它已经关闭时会发生什么.
```  

If we create a very simple app that tries to send to an empty channel, we get a deadlock:

```
func main() {
	var a chan int
	a <- 1
}
```

If we do the same for a receiving operation, we get the same result of a deadlock:

```
如果对接收操作执行相同的操作，则会得到相同的死锁结果：


func main() {
    var a chan int
    <-a
}

```

The last thing left to check is how the close function behaves with a nil channel. It panics with the close of nil channel explicit value:

```
最后要检查的是close函数在nil通道中的行为。当接近零通道显式值时会恐慌：

func main() {
    var a chan int
    close(a)
}
```

