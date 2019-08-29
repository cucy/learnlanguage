
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
# nil channels

We have previously discussed how channels belong to the pointer types in Go, so their default value is nil. But what happens when you send or receive from a nil channel?

```
我们之前已经讨论了通道如何属于go中的指针类型，因此它们的默认值是nil。但是当你从一个零通道发送或接收时会发生什么呢？
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

To recap, we have seen that a nil channel's send and receive are blocking operations, and that close causes a panic.

```
综上所述，我们已经看到一个零通道的发送和接收正在阻塞操作，而关闭操作会引起恐慌。
```



# Closed channels

We already know that receiving from a closed channel returns a zero value for the channel type, and a second Boolean is false. But what happens if we try to send something after closing the channel? Let's find out with the help of the following code:

```
我们已经知道从一个封闭的通道接收会为通道类型返回一个零值，第二个布尔值是假的。但是，如果我们在关闭频道后尝试发送内容，会发生什么？让我们通过以下代码来了解：

func main() {
	a := make(chan int)
	close(a)
	a <- 1  // panic: send on closed channel
}
```

If we try to send data after closing, it will return a very specific panic: send on closed channel. A similar thing will happen when we try to close a channel that has already been closed:

```
如果我们试图在关闭后发送数据，它将返回一个非常具体的恐慌：在关闭的通道上发送。
当我们试图关闭已关闭的频道时，也会发生类似的情况：

func main() {
	a := make(chan int)
	close(a)
	close(a)  // panic: close of closed channel
}
```
This example will cause a panic with a specific value—close of closed channel.

``` 
此示例将导致恐慌，其特定值为Close of Closed Channel。
```

# Managing multiple operations

There are many situations in which more than one goroutines are executing their code and communicating through channels. A typical scenario is to wait for one of the channels' send or receive operations to be executed.

When you operate with many channels, Go makes it possible to use a special keyword that executes something similar to switch but for channel operations. This is done with the select statement, followed by a series of case statements and an optional default case.

We can see a quick example of where we are receiving a value from a channel in a goroutine, and sending a value to another channel in a different goroutine. In these, the main goroutine we are using is a select statement to interact with the two channels, receive from the first, and send to the second:

```
在许多情况下，不止一个goroutine正在执行其代码并通过通道进行通信。典型的场景是等待某个通道的发送或接收操作被执行。
当您使用多个通道操作时，go可以使用一个执行类似于switch但用于通道操作的特殊关键字。这是通过select语句完成的，后面是一系列case语句和可选的默认case。
我们可以看到一个快速的例子，我们从Goroutine中的一个通道接收到一个值，然后将一个值发送到另一个Goroutine中的另一个通道。其中，我们使用的主要goroutine是一个select语句，用于与两个通道交互、从第一个通道接收和发送到第二个通道：

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	go func() { <-ch1 }()
	go func() { ch2 <- a }()
	select {
	case ch1 <- b:
		fmt.Println("ch1 got a", b)
	case v := <-ch2:
		fmt.Println("ch2 got a", v)
	}
}
```

When running this program in the playground, we can see that the receive operation from the second channel is always the first to finish.
If we switch the execution order of the goroutines, we get the opposite results. 
The operation executed last is the one picked up first.
This happens because the playground is a web service that runs and executes Go code in a safe environment and does some optimizations to make this operation deterministic.

```
当在操场上运行这个程序时，我们可以看到第二个通道的接收操作总是第一个完成的。
如果我们切换goroutine的执行顺序，就会得到相反的结果。最后执行的操作是先执行的操作。
这是因为Playground是一个在安全环境中运行和执行Go代码的Web服务，并进行了一些优化以使此操作具有确定性。
```

# Default clause

If we add a default case to the previous example, the result of the application execution will be very different, particularly if we change select:

```
如果我们在前面的示例中添加一个缺省情况，应用程序执行的结果将非常不同，特别是如果我们更改选择：

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	go func() { <-ch1 }()
	go func() { ch2 <- a }()
	select {
	case ch1 <- b:
		fmt.Println("ch1 got a", b)
	case v := <-ch2:
		fmt.Println("ch2 got a", v)
	default:
		fmt.Println("too slow")
	}
}
```

The select statement will always choose the default statement. This happens because the goroutines are not picked up by the scheduler yet, when the select statement is executed. If we add a very small pause (using time.Sleep) before the select switch, we will have the scheduler pick at least one goroutine and we will then have one of the two operations executed:

```
select语句将始终选择默认语句。这是因为在执行select语句时，调度程序尚未拾取goroutine。如果在选择开关之前添加一个非常小的暂停（使用time.sleep），调度程序将至少选择一个goroutine，然后执行两个操作之一：

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	for i := 0; i < 10; i++ {
		go func() { <-ch1 }()
		go func() { ch2 <- a }()
		time.Sleep(time.Nanosecond)

		select {
		case ch1 <- b:
			fmt.Println("ch1 got a", b)
		case v := <-ch2:
			fmt.Println("ch2 got a", v)
		default:
			fmt.Println("too slow")
		}
	}
}

/*************/
ch2 got a 2
ch2 got a 2
ch2 got a 2
ch2 got a 2
ch2 got a 2
ch2 got a 2
ch1 got a 10
ch1 got a 10
ch2 got a 2
ch1 got a 10

```

In this case, we will have a mixed set of operations executed, depending on which one gets picked up by the Go scheduler.

```
在这种情况下，我们将执行一组混合的操作，这取决于Go调度程序选择哪一个操作。
```

# Timers and tickers

The time package offers a couple of tools that make it possible to orchestrate goroutines and channels—timers and tickers.

```
时间包提供了两种工具，可以协调goroutine和channels计时器和ticker。


```

# Timers

The utility that can replace the default clause in a select statement is the time.Timer type. This contains a receive-only channel that will return a time.Time value after the duration specified during its construction, using time.NewTimer:

```
可以替换select语句中默认子句的实用工具是time.timer类型。
这包含一个只接收的通道，该通道将返回一个时间。时间值在其构造过程中指定的持续时间之后，使用time.newtimer:


func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	for i := 0; i < 10; i++ {
		go func() { <-ch1 }()
		go func() { ch2 <- a }()

		t := time.NewTimer(time.Nanosecond)

		select {
		case ch1 <- b:
			fmt.Println("ch1 got a", b)
		case v := <-ch2:
			fmt.Println("ch2 got a", v)
		case <-t.C:
			fmt.Println("too slow")
		}
	}
}
```

A timer exposes a read-only channel, so it's not possible to close it. When created with time.NewTimer, it waits for the specified duration before firing a value in the channel.

The Timer.Stop method will try to avoid sending data through the channel and return whether it succeeded or not. If false is returned after trying to stop the timer, we still need to receive the value from the channel before being able to use the channel again.

Timer.Reset restarts the timer with the given duration, and returns a Boolean as it happens with Stop. This value is either true or false:

true when the timer is active
false when the timer was fired or stopped

```

计时器公开只读通道，因此无法关闭它。当使用time.newtimer创建时,它会等待指定的持续时间，然后在通道中触发一个值。
stop方法将尝试避免通过通道发送数据，并返回是否成功。如果在尝试停止计时器后返回false，则在再次使用通道之前，我们仍然需要从通道接收值。
reset以给定的持续时间重新启动计时器，并返回与stop相同的布尔值。该值为真或假：
当计时器激活时为真
当计时器被触发或停止时为假

```

We will test these functionalities with a practical example:

```
我们将用一个实际例子来测试这些功能：

t := time.NewTimer(time.Millisecond)
time.Sleep(time.Millisecond / 2)
if !t.Stop() {
    panic("it should not fire")
}
select {
case <-t.C:
    panic("not fired")
default:
    fmt.Println("not fired")
}

We are creating a new timer of 1ms. Here, we wait 0.5ms and then stop it successfully:

我们正在创建一个1毫秒的新计时器。在这里,我们等待0.5毫秒,然后成功停止它：
```


```


if t.Reset(time.Millisecond) {
    panic("timer should not be active")
}
time.Sleep(time.Millisecond)
if t.Stop() {
    panic("it should fire")
}
select {
case <-t.C:
    fmt.Println("fired")
default:
    panic("not fired")
}

```

Then, we are resetting the timer back to 1ms and waiting for it to fire, to see whether Stop returns false and the channel gets drained.

``` 
然后，我们将计时器重置回1毫秒，并等待它启动，以查看Stop是否返回false，以及通道是否被耗尽。



func main() {
	t := time.NewTimer(time.Millisecond)
	time.Sleep(time.Millisecond / 2)
	if !t.Stop() {
		panic("it should not fire")
	}
	select {
	case <-t.C:
		panic("not fired")
	default:
		fmt.Println("not fired")
	}
	if t.Reset(time.Millisecond) {
		panic("timer should not be active")
	}
	time.Sleep(time.Millisecond)
	if t.Stop() {
		panic("it should fire")
	}
	select {
	case <-t.C:
		fmt.Println("fired")
	default:
		panic("not fired")
	}
}



```

Then, we are resetting the timer back to 1ms and waiting for it to fire, to see whether Stop returns false and the channel gets drained.

```
然后,我们将计时器重置回1毫秒,并等待其启动,以查看停止是否返回错误,以及通道是否被排空。


```

# AfterFunc

A very useful utility that uses time.Timer is the time.AfterFunc function, which returns a timer that will execute the passed function in its own goroutine when the timer fires:

```
使用time.timer的一个非常有用的实用程序是time.afterfunc函数，它返回一个计时器，当计时器触发时，该计时器将在自己的goroutine中执行传递的函数：


func main() {
	time.AfterFunc(time.Millisecond, func() {
		fmt.Println("Hello 1!")
	})
	t := time.AfterFunc(time.Millisecond*5, func() {
		fmt.Println("Hello 2!")
	})
	if !t.Stop() {
		panic("should not fire")
	}
	time.Sleep(time.Millisecond * 10)
}

```

In the previous example, we define two timers for two different closures, and we stop one of them and let the other trigger.

在前一个示例中，我们为两个不同的闭包定义了两个计时器，并停止其中一个计时器，让另一个触发器触发。


# Tickers

time.Ticker is similar to time.Timer, but its channel delivers more elements at regular intervals equal to the duration. They are specified when creating it with time.NewTicker.  This makes it possible to stop the ticker from firing with the Ticker.Stop method:

时间。断续器类似于时间计时器，但它的通道以与持续时间相等的定期间隔提供更多的元素。使用time.newticker创建时会指定它们。这样就可以阻止ticker使用ticker.stop方法启动：

```
func main() {
	tick := time.NewTicker(time.Millisecond)
	stop := time.NewTimer(time.Millisecond * 10)

	for {
		select {
		case a := <-tick.C:
			fmt.Println(a)
		case <-stop.C:
			tick.Stop()
		case <-time.After(time.Millisecond):
			return
		}
	}
}


```

In this example, we are also using time.After—a function that returns the channel from an anonymous time.Timer. This can be used when there's no need to stop the timer. There is another function, time.Tick, that returns the channel of an anonymous time.Ticker. Both functions will return a channel that the application has no control over and this channel will eventually be picked up by the garbage collector.

This concludes the overview of channels, from their properties and basic usage to some more advanced concurrency examples. We also checked some special cases and how to synchronize multiple channels.

在本例中，我们还使用time.after—一个从匿名time.timer返回通道的函数。这可以在不需要停止计时器时使用。还有一个函数time.tick，它返回匿名time.ticker的通道。这两个函数都将返回应用程序无法控制的通道，垃圾收集器最终将接收该通道。
这总结了通道的概述，从它们的属性和基本用法到一些更高级的并发性示例。我们还检查了一些特殊情况以及如何同步多个通道。

# Combining channels and goroutines

Now that we know the fundamental tools and properties of Go concurrency, we can use them to build better tools for our applications. We will see some examples that make use of channels and goroutines to solve real-world problems.

组合通道和Goroutines
既然我们了解了Go并发的基本工具和属性，那么就可以使用它们为我们的应用程序构建更好的工具。我们将看到一些利用渠道和Goroutine解决现实问题的例子。

# Rate limiter

A typical scenario is having a web API that has a certain limit to the number of calls that can be done in a certain period of time. This type of API will just prevent the usage for a while if this threshold is crossed, making it unusable for the time being. When creating a client for the API, we need to be aware of this and make sure our application does not overuse it.

That's a very good scenario where we can use time.Ticker to define an interval between calls. In this example, we will create a client for Google Maps' geocoding service that has a limit of 100,000 requests per 24 hours. Let's start by defining the client:

一个典型的场景是有一个Web API，它对在特定时间段内可以完成的调用数量有一定的限制。如果跨过此阈值，这种类型的API将暂时阻止使用，使其暂时不可用。在为API创建客户机时，我们需要注意这一点，并确保我们的应用程序不会过度使用它。
这是一个很好的场景，我们可以使用time.ticker来定义通话间隔。在这个例子中，我们将为谷歌地图的地理编码服务创建一个客户端，它的请求限制为每24小时100000个。让我们从定义客户机开始：

```
type Client struct {
	client *http.Client
	tick *time.Ticker
} 
```


The client is made by an HTTP client that will call maps, a ticker that will help prevent passing the rate limit, and needs an API key for authentication with the service. We can define a custom Transport struct for our use case that will inject the key in the request as follows:

客户机由一个HTTP客户机生成，该客户机将调用maps，这是一个有助于防止超过速率限制的标记器，并且需要一个API密钥来对服务进行身份验证。我们可以为我们的用例定义一个自定义的传输结构，它将在请求中注入密钥，如下所示：

```
// API传输
type apiTransport struct {
	http.RoundTripper
	key string
}
// RoundTrip 往返
func (a apiTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	q := r.URL.Query()
	q.Set("key", a.key)
	r.URL.RawQuery = q.Encode()
	return a.RoundTripper.RoundTrip(r)
}

```

This is a very good example of how Go interfaces allow the extension of their own behavior. We are defining a type that implements the http.RoundTripper interface, and also an attribute that is an instance of the same interface. The implementation injects the API key to the request before executing the underlying transport. This type allows us to define a helper function that creates a new client, where we are using the new transport that we defined together with the default one:
这是一个非常好的例子，说明了Go接口如何允许扩展自己的行为。我们正在定义一个实现http.roundtripper接口的类型，以及一个属于同一接口的实例的属性。在执行底层传输之前，实现将API密钥注入请求。此类型允许我们定义一个创建新客户机的助手函数，在该函数中，我们使用与默认传输一起定义的新传输：

```
func NewClient(tick time.Duration, key string) *Client {
	return &Client{
		client: &http.Client{
			Transport: apiTransport{http.DefaultTransport, key},
		},
		tick: time.NewTicker(tick),
	}
}


```

The maps geocoding API returns a series of addresses that are composed of various parts. This is available at 

maps geocoding api返回一系列由不同部分组成的地址。这是在 https://developers.google.com/maps/documentation/geocoding/intro#GeocodingResponses

```

type Result struct {
	AddressComponents []struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		// more fields
	} `json:"geometry"`
	PlaceID string `json:"place_id"`
	// more fields
}


```

We can use the structure to execute a reverse geocoding operation—getting a location from the coordinates by using the respective endpoint. We wait for the ticket before executing the HTTP request, remembering to defer the closure of the body:

我们可以使用该结构执行反向地理编码操作，通过使用各自的端点从坐标中获取位置。在执行HTTP请求之前，我们等待通知单，记住延迟关闭主体：

```
const url = "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v"
<-c.tick.C
resp, err := c.client.Get(fmt.Sprintf(url, lat, lng))
if err != nil {
    return nil, err
}
defer resp.Body.Close()
```

Then, we can decode the result in a data structure that uses the Result type we already defined and checks for the status string:
然后，我们可以在使用我们已经定义的结果类型的数据结构中解码结果，并检查状态字符串：

```
var v struct {
    Results []Result `json:"results"`
    Status string `json:"status"`
}
// get the result
if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
    return nil, err
}
switch v.Status {
case "OK":
    return v.Results, nil
case "ZERO_RESULTS":
    return nil, nil
default:
    return nil, fmt.Errorf("status: %q", v.Status)
}
}
```

Finally, we can use the client to geocode a series of coordinates, expecting the requests to be at least 860ms from each other:
最后，我们可以使用客户机对一系列坐标进行地理编码，期望相互之间的请求至少为860ms：

```
c := NewClient(24*time.Hour/100000, os.Getenv("MAPS_APIKEY"))
start := time.Now()
for _, l := range [][2]float64{
    {40.4216448, -3.6904040},
    {40.4163111, -3.7047328},
    {40.4123388, -3.7096724},
    {40.4145150, -3.7064412},
} {
    locs, err := c.ReverseGeocode(l[0], l[1])
    e := time.Since(start)
    if err != nil {
        log.Println(e, l, err)
        continue
    }
    // just print the first location
    if len(locs) != 0 {
        locs = locs[:1]
    }
    log.Println(e, l, locs)
}
```


# Workers

The previous example is a Google Maps client that uses a time.Ticker channel to limit the rate of the requests. The rate limit makes sense for an API key. Let's imagine that we have more API keys from different accounts, so we could potentially execute more requests.

A very typical concurrent approach is the workers pool. Here, you have a series of clients that can be picked up to process an input and different parts of the application can ask to use such clients, returning the clients back when they are done.

We can create more than one client that shares the same channels for both requests and responses, with requests being the coordinates and the results being the response from the service. Since the channel for responses is unique, we can define a custom type that holds all the information needed for that channel:


```
前一个例子是一个GoogleMaps客户端，它使用一个time.ticker通道来限制请求的速率。速率限制对于API密钥是有意义的。假设我们有更多来自不同帐户的API密钥，这样我们就可以潜在地执行更多的请求。
一个非常典型的并发方法是工人池。在这里，您有一系列的客户机可以用来处理输入，应用程序的不同部分可以要求使用这些客户机，完成后将客户机返回。
我们可以为请求和响应创建多个共享相同通道的客户机，请求是坐标，结果是来自服务的响应。由于响应通道是唯一的，因此我们可以定义一个自定义类型，该类型保存该通道所需的所有信息：


```

```
type result struct {
    Loc [2]float64
    Result []maps.Result
    Error error
}
```

The next step is creating the channels—we are going to read a comma-separated list of values from an environment variable here. We will create a channel for requests, and one for responses. Both channels have a capacity equal to the number of workers, in this case, but this would work even if the channels were unbuffered. Since we are just using channels, we will need another channel, done, which signals whether a worker has finished working on their last job:

```
下一步是创建通道，我们将在这里从环境变量中读取以逗号分隔的值列表。我们将为请求创建一个通道，并为响应创建一个通道。在这种情况下，两个通道的容量都等于工人的数量，但即使通道没有缓冲，这也会起作用。因为我们只是在使用通道，所以我们需要另一个通道done，它指示工人是否完成了上一个工作：

keys := strings.Split(os.Getenv("MAPS_APIKEYS"), ",")
requests := make(chan [2]float64, len(keys))
results := make(chan result, len(keys))
done := make(chan struct{})
```

Now, we will create a goroutine for each of the keys, in which we define a client that feeds on the requests channel, executes the request, and sends the result to the dedicated channel. When the requests channel is closed, the goroutine will exit the range and send a message to the done channel, which is shown in the following code:

现在，我们将为每个键创建一个goroutine，在其中定义一个客户机，该客户机提供请求通道、执行请求并将结果发送到专用通道。
当请求通道关闭时，goroutine将退出该范围并向完成通道发送一条消息，如下代码所示：

```
for i := range keys {
    go func(id int) {
        log.Printf("Starting worker %d with API key %q", id, keys[id])
        client := maps.NewClient(maps.DailyCap, keys[id])
        for j := range requests {
            var r = result{Loc: j}
            log.Printf("w[%d] working on %v", id, j)
            r.Result, r.Error = client.ReverseGeocode(j[0], j[1])
            results <- r
        }
        done <- struct{}{}
    }(i)
}
```

The locations can be sent to the request channel sequentially in another goroutine:
这些位置可以在另一个goroutine中按顺序发送到请求通道：

```
go func() {
    for _, l := range [][2]float64{
        {40.4216448, -3.6904040},
        {40.4163111, -3.7047328},
        {40.4123388, -3.7096724},
        {40.4145150, -3.7064412},
    } {
        requests <- l
    }
    close(requests)
}()
```

The channel is used to count how many workers are done, and once every one of them is done, it will close the result channel. This will allow us to just loop over it to get the result:
通道用于计算完成了多少个工作人员，一旦每个工作人员都完成了，它将关闭结果通道。这将允许我们循环它以获得结果：

```
for r := range results {
    log.Printf("received %v", r)
}
```

Using a channel is just one of the ways to wait for all the goroutines to finish, and we will see more idiomatic ways of doing it in the next chapter with the sync package.

使用通道只是等待所有Goroutine完成的方法之一，我们将在下一章的同步包中看到更多的惯用方法。

# Pool of workers


A channel can be used as a pool of resources that allows us to request them on demand. In the following example, we will create a small application that will look up which addresses are valid in a network, using a third-party client from the github.com/tatsushid/go-fastping package.

The pool will have two methods, one for getting a new client and another to return the client back to the pool. The Get method will try to get an existing client from the channel or return a new one if this is not available. The Put method will try to put the client back in the channel, or discard it otherwise:

一个通道可以作为一个资源池，允许我们根据需要请求它们。在下面的示例中，我们将使用github.com/tatsushid/go-fastping包中的第三方客户端创建一个小应用程序，该应用程序将查找网络中哪些地址有效。
池将有两种方法，一种用于获取新客户机，另一种用于将客户机返回池。get方法将尝试从通道中获取现有的客户机，如果不可用，则返回新的客户机。Put方法将尝试将客户机放回通道中，否则放弃它：


```
const wait = time.Millisecond * 250

type pingPool chan *fastping.Pinger

func (p pingPool) Get() *fastping.Pinger {
    select {
    case v := <-p:
        return v
    case <-time.After(wait):
        return fastping.NewPinger()
    }
}

func (p pingPool) Put(v *fastping.Pinger) {
    select {
    case p <- v:
    case <-time.After(wait):
    }
    return
}
```

The client will need to specify which network needs to be scanned, so it requires a list of available networks starting with the net.Interfaces function, ranging through the interfaces and their addresses:

客户端需要指定需要扫描的网络，因此它需要从net.interfaces函数开始的可用网络列表，范围从接口及其地址开始：

```
ifaces, err := net.Interfaces()
if err != nil {
    return nil, err
}
for _, iface := range ifaces {
    // ...
    addrs, err := iface.Addrs()
    // ...
    for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
            ip = v.IP
        case *net.IPAddr:
            ip = v.IP
        }
        // ...
        if ip = ip.To4(); ip != nil {
            result = append(result, ip)
        }
    }
}
```

We can accept a command-line argument to select between interfaces, and we can show a list of interfaces to the user to select when the argument is either not present or wrong:

我们可以接受命令行参数来在接口之间进行选择，当参数不存在或错误时，我们可以向用户显示要选择的接口列表：

```
if len(os.Args) != 2 {
    help(ifaces)
}
i, err := strconv.Atoi(os.Args[1])
if err != nil {
    log.Fatalln(err)
}
if i < 0 || i > len(ifaces) {
    help(ifaces)
}
```

The help function is just a print of the interfaces IP:
帮助功能只是接口IP的打印：

```
func help(ifaces []net.IP) {
    log.Println("please specify a valid network interface number")
    for i, f := range ifaces {
        mask, _ := f.DefaultMask().Size()
        fmt.Printf("%d - %s/%v\n", i, f, mask)
    }
    os.Exit(0)
}
```

The next step is obtain the range of IPs that need to be checked:
下一步是获取需要检查的IP范围：

```
m := ifaces[i].DefaultMask()
ip := ifaces[i].Mask(m)
log.Printf("Lookup in %s", ip)
```

Now that we have the IP, we can create a function to obtain other IPs in the same network. IPs in Go are a byte slice, so we will replace the least significant bits in order to obtain the final address. Since the IP is a slice, its value will be overwritten by each operation (slices are pointers). We are going to update a copy of the original IP—because slices are pointers to the same array—in order to avoid overwrites:
既然我们有了IP，就可以创建一个函数来获取同一网络中的其他IP。go中的ips是一个字节片，因此我们将替换最低有效位以获得最终地址。由于IP是一个切片，其值将被每个操作覆盖（切片是指针）。我们将更新原始IP的副本，因为切片是指向同一数组的指针，以避免覆盖：

```
func makeIP(ip net.IP, i int) net.IP {
    addr := make(net.IP, len(ip))
    copy(addr, ip)
    b := new(big.Int)
    b.SetInt64(int64(i))
    v := b.Bytes()
    copy(addr[len(addr)-len(v):], v)
    return addr
}
```

Then, we will need one channel for results and another for keeping a track of the goroutines; and for each IP, we need to check whether we can launch a goroutine for each address. We will use a pool of 10 clients and inside each goroutine—we will ask for each client, then return them to the pool. All valid IPs will be sent through the result channel:

然后，我们需要一个通道来获取结果，另一个通道来跟踪Goroutine；对于每个IP，我们需要检查是否可以为每个地址启动Goroutine。我们将使用一个由10个客户组成的池，在每个Goroutine中，我们将请求每个客户，然后将其返回到池中。所有有效IP将通过结果通道发送：

```
done := make(chan struct{})
address := make(chan net.IP)
ones, bits := m.Size()
pool := make(pingPool, 10)
for i := 0; i < 1<<(uint(bits-ones)); i++ {
    go func(i int) {
        p := pool.Get()
        defer func() {
            pool.Put(p)
            done <- struct{}{}
        }()
        p.AddIPAddr(&net.IPAddr{IP: makeIP(ip, i)})
        p.OnRecv = func(a *net.IPAddr, _ time.Duration) { address <- a.IP }
        p.Run()
    }(i)
}
```

Each time a routine finishes, we send a value in the done channel so we can keep count of the done signals received before exiting the application. This will be the result loop:
每次一个例程完成时，我们都会在done通道中发送一个值，这样我们就可以在退出应用程序之前保持接收到的done信号的计数。这将是结果循环：

```
i = 0
for {
    select {
    case ip := <-address:
        log.Printf("Found %s", ip)
    case <-done:
        if i >= bits-ones {
            return
        }
        i++
    }
}
```

The loop will continue until the count from the channel reaches the number of goroutines.
This concludes the more convoluted examples of the usage of channels and goroutines together.
循环将继续，直到来自通道的计数达到goroutine的数目。这就总结了使用通道和Goroutine的更复杂的例子。

# Semaphores 信号量

Semaphores are tools used to solve concurrency issues. They have a certain number of available quotas that is used to limit the access to resources; also, various threads can request one or more quotas from it, and then release them when they are done. If the number of quotas available is one, it means that the semaphore supports only one access at time, with a behavior similar to mutexes. If the quota is more than one, we are referring to the most common type—the weighted semaphore.

In Go, a semaphore can be implemented using a channel with a capacity equal to the quotas, where you send a message to the channel to acquire a quota, and receive one from it to release:

信号量是用来解决并发性问题的工具。它们具有一定数量的可用配额，用于限制对资源的访问；此外，各种线程可以从中请求一个或多个配额，然后在完成后释放它们。
如果可用配额数为1，则意味着信号量一次只支持一个访问，其行为类似于互斥锁。如果配额不止一个，我们指的是最常见的类型加权信号量。
在go中，可以使用容量等于配额的通道来实现信号量，在该通道中，您向该通道发送消息以获取配额，并从该通道接收一个消息以释放：

```
type sem chan struct{}

func (s sem) Acquire() {
    s <- struct{}{}
}

func (s sem) Relase() {
    <-s
}
```

The preceding code shows us how to implement a semaphore using a channel in a few lines. Here's an example of how to use it:
前面的代码向我们展示了如何在几行中使用通道来实现信号量。下面是一个如何使用它的示例：

```
func main() {
    s := make(sem, 5)
    for i := 0; i < 10; i++ {
        go func(i int) {
            s.Acquire()
            fmt.Println(i, "start")
            time.Sleep(time.Second)
            fmt.Println(i, "end")
            s.Relase()
        }(i)
    }
    time.Sleep(time.Second * 3)
}

```

We can see from the previous example how the program serves some requests on the first round of acquisition, and the others on the second round, not allowing more than five executions at the same time.
我们可以从前面的示例中看到，程序如何在第一轮采集中处理某些请求，而在第二轮采集中处理其他请求，同时不允许执行超过五次。

# Summary

In this chapter, we talked about the two main actors in Go concurrency—goroutines and channels. We started by explaining what a thread is, what the differences are between threads and goroutines, and why they are so convenient. Threads are heavy and require a CPU core, while goroutines are lightweight and not bound to a core. We saw how easily a new goroutine can be started by executing a function preceded by the go keyword, and how it is possible to start a series of different goroutines at once. We saw how the arguments of the concurrent functions are evaluated when the goroutine is created and not when it actually starts. We also saw that it is very difficult to keep different goroutines in sync without any additional tools.

Then, we introduced channels that are used to share information between different goroutines and solve the synchronization problem that we mentioned previously. We saw that goroutines have a maximum capacity and a size—how many elements it is holding at present. Size cannot overcome capacity, and when an extra element is sent to a full channel, the operation blocks it until an element is removed from the channel. Receiving from a channel that is empty is also a blocking operation.

We saw how to close channels with the close function, how this operation should be done in the same goroutine that sends data, and how operations behave in special cases such as nil or a closed channel. We introduced the select statement to choose between concurrent channel operations and control the application flow. Then, we introduced the tools related to concurrency from the time package—tickers and timers.

Finally, we showed some real-world examples, including a rate-limited Google Maps client and a tool to simultaneously ping all the addresses of a network.

In the next chapter, we will look at some synchronization primitives that will allow a better handling of goroutines and memory, using more clear and simple code.

总结
在本章中，我们讨论了Go并发Goroutines和Channels中的两个主要参与者。我们首先解释线程是什么，线程和Goroutine之间的区别是什么，以及为什么它们如此方便。线程很重，需要一个CPU核心，而Goroutines是轻量级的，不绑定到核心。我们看到了执行go关键字前面的函数可以多么容易地启动一个新的goroutine，以及如何能够同时启动一系列不同的goroutine。我们看到了在创建goroutine时如何计算并发函数的参数，而不是在实际启动时如何计算。我们还发现，在没有任何附加工具的情况下，很难保持不同的goroutine同步。
然后，我们介绍了用于在不同goroutine之间共享信息的通道，并解决了前面提到的同步问题。我们看到Goroutine有一个最大的容量和大小，它目前可以容纳多少个元素。大小无法克服容量，当一个额外的元素被发送到一个完整的通道时，操作将阻止它，直到从通道中删除一个元素。从空通道接收也是一种阻塞操作。
我们看到了如何使用close函数关闭通道，如何在发送数据的同一个goroutine中执行此操作，以及在诸如nil或封闭通道等特殊情况下操作的行为。我们引入了select语句来在并发通道操作和控制应用程序流之间进行选择。然后，我们从时间包标记器和计时器中介绍了与并发性相关的工具。
最后，我们展示了一些真实的例子，包括一个速率受限的GoogleMaps客户端和一个同步ping网络所有地址的工具。
在下一章中，我们将研究一些同步原语，这些原语将允许使用更清晰和简单的代码更好地处理goroutine和内存。

# Questions
What is a thread and who is responsible for it?
Why are goroutines different from threads?
When are arguments evaluated when launching a goroutine?
What's the difference between buffered and unbuffered channels?
Why are one-way channels useful?
What happens when operations are done on nil or closed channels?
What are timers and tickers used for?

什么是线程，谁负责它？
为什么Goroutines不同于线程？
启动goroutine时何时评估参数？
缓冲通道和非缓冲通道有什么区别？
为什么单向通道有用？
在零通道或封闭通道上进行操作时会发生什么？
什么是计时器和自动售票机？


# Synchronization with sync and atomic

This chapter will continue the journey into Go concurrency, introducing the sync and atomic packages, which are a couple of other tools designed for orchestrating synchronization between goroutines. This will make it possible to write elegant and simple code that allows concurrent usage of resources and manages a goroutine's lifetime. sync contains high-level synchronization primitives, while atomic contains low-level ones.

The following topics will be covered in this chapter:

Lockers
Wait groups
Other sync components
The atomic package
  
同步和原子同步
本章将继续到Go并发的旅程，介绍同步和原子包，这是为协调Goroutines之间的同步而设计的其他两个工具。这将使编写优雅和简单的代码成为可能，这些代码允许并发使用资源并管理Goroutine的生命周期。同步包含高级同步原语，而原子包含低级同步原语。
本章将讨论以下主题：
储物柜
等待组
其他同步组件
原子包


# Synchronization primitives
We saw how channels are focused on communication between goroutines, and now we will focus on the tools offered by the sync package, which includes the basic primitives for synchronization between goroutines. The first thing we will see is how to implement concurrent access to the same resource with lockers.

同步原语
我们看到了通道如何专注于Goroutine之间的通信，现在我们将重点关注同步包提供的工具，其中包括Goroutine之间同步的基本原语。我们将看到的第一件事是如何用锁实现对同一个资源的并发访问。

# Concurrent access and lockers 并行访问和储物柜

Go offers a generic interface for objects that can be locked and unlocked. Locking an object means taking control over it while unlocking releases it for others to use. This interface exposes a method for each operation. The following is an example of this in code:

Go为可以锁定和解锁的对象提供了一个通用接口。锁定一个对象意味着控制它，同时解锁释放它供其他人使用。此接口为每个操作公开一个方法。以下是代码中的一个示例：

```
type Locker interface {
    Lock()
    Unlock()
}
```

# Mutex
The most simple implementation of locker is sync.Mutex. Since its method has a pointer receiver, it should not be copied or passed around by value. The Lock() method takes control of the mutex if possible, or blocks the goroutine until the mutex becomes available. The Unlock() method releases the mutex and it returns a runtime error if called on a non-locked one.

Here is a simple example in which we launch a bunch of goroutines using the lock to see which is executed first:

互斥
locker最简单的实现是sync.mutex。因为它的方法有一个指针接收器，所以不应该通过值来复制或传递它。lock（）方法在可能的情况下控制互斥体，或者在互斥体可用之前阻止Goroutine。unlock（）方法释放互斥体，如果对未锁定的互斥体调用，它将返回运行时错误。
下面是一个简单的示例，其中我们使用锁启动一组goroutine，以查看首先执行的是哪个：

```
func main() {
	var m sync.Mutex
	done := make(chan struct{}, 10)
	for i := 0; i < cap(done); i++ {
		go func(i int, l sync.Locker) {
			l.Lock()
			defer l.Unlock()
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
			done <- struct{}{}
		}(i, &m)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}
```


We are using a channel to signal the main goroutine when a job is done, and exit the application. Let's create an external counter and increment it concurrently using goroutines.

Operations executed on different goroutines are not thread-safe, as we can see from the following example:

当一个任务完成时，我们使用一个通道向主Goroutine发出信号，然后退出应用程序。让我们创建一个外部计数器，并使用goroutine同时递增它。
在不同的goroutine上执行的操作不是线程安全的，如下面的示例所示：

```
func main() {
	done := make(chan struct{}, 10000)
	var a = 0
	for i := 0; i < cap(done); i++ {
		go func(i int) {
			if i%2 == 0 {
				a++
			} else {
				a--
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
	fmt.Println(a)
}

```

We would expect to have 5000 plus one, and 5000 minus one, with a 0 printed in the final instruction. However, what we get are different values each time we run the application. This happens because these kind of operations are not thread-safe, so two or more of them could happen at the same time, with the last one shadowing the others. This kind of phenomena is known as a race condition; that is, when more than one operation is trying to write the same result.

This means that without any synchronization, the result is not predictable; if we check the previous example and use a lock to avoid the race condition, we will have zero as the value for the integer—the result that we were expecting:

我们希望有5000加1，5000减1，最后一条指令中打印0。但是，每次运行应用程序时，我们得到的是不同的值。这是因为这些类型的操作不是线程安全的，所以它们中的两个或多个操作可能同时发生，最后一个操作会隐藏其他操作。这种现象被称为竞争条件；也就是说，当多个操作试图写入相同的结果时。
这意味着，如果不进行任何同步，结果是不可预测的；如果我们检查前面的示例并使用一个锁来避免竞争条件，我们将使用零作为整数的值，即我们所期望的结果：


```
func main() {
	done := make(chan struct{}, 10000)
	m := sync.Mutex{}

	var a = 0
	for i := 0; i < cap(done); i++ {
		go func(l sync.Locker,i int) {
			l.Lock()
			defer l.Unlock()
			if i%2 == 0 {
				a++
			} else {
				a--
			}
			done <- struct{}{}
		}(&m, i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
	fmt.Println(a)
}
```

A very common practice is embedding a mutex in a data structure to symbolize that the container is the one you want to lock. The counter variable from before can be represented as follows:

一个非常常见的实践是在数据结构中嵌入互斥体，以表示容器是要锁定的容器。前面的计数器变量可以表示为：

```
type counter struct {
    m     sync.Mutex
    value int
}
```

The operations that the counter performs can be methods that already take care of locking before the main operation, along with unlocking it afterward, as shown in the following code block:
计数器执行的操作可以是在主操作之前已经处理好锁定的方法，以及随后解锁的方法，如下面的代码块所示：

```
func (c *counter) Incr(){
    c.m.Lock()
    c.value++
    c.m.Unlock()
}

func (c *counter) Decr(){
    c.m.Lock()
    c.value--
    c.m.Unlock()
}

func (c *counter) Value() int {
    c.m.Lock()
    a := c.value
    c.m.Unlock()
    return a
}
```

This will simplify the goroutine loop, resulting in a much clearer code:
这将简化Goroutine循环，从而得到更清晰的代码。

```
var a = counter{}
for i := 0; i < cap(done); i++ {
    go func(i int) {
        if i%2 == 0 {
            a.Incr()
        } else {
            a.Decr()
        }
        done <- struct{}{}
    }(i)
}
// ...
fmt.Println(a.Value())
```

# RWMutex

The problem with race conditions is caused by concurred writing, not by reading the operation. The other data structure that implements the locker interface, sync.RWMutex, is made to support both these operations, having write locks that are unique and mutually exclusive with read locks. This means that the mutex can be locked either by a single write lock, or by one or more read locks. When a reader locks the mutex, other readers trying to lock it will not be blocked. They are often referred to as shared-exclusive locks. This allows read operations to happen all at the same time, without there being a waiting time.

The write lock operations are done using the Lock and Unlock methods of the locker interface. The reading operations are executed using two other methods: RLock and RUnlock. There is another method, RLocker, which returns a locker for reading operations.

We can make a quick example of their usage by creating a concurrent list of strings:

比赛条件的问题是由一致的书写引起的，而不是由读取操作引起的。实现locker接口的另一个数据结构sync.rwmutex支持这两种操作，具有唯一的写锁，与读锁互斥。这意味着互斥锁可以通过一个写锁或一个或多个读锁来锁定。当读卡器锁定互斥体时，其他试图锁定互斥体的读卡器将不会被阻止。它们通常被称为共享独占锁。这允许读取操作同时发生，而不需要等待时间。
写锁操作是使用locker接口的lock和unlock方法完成的。读取操作使用另外两种方法执行：rlock和runlock。还有另一个方法rlocker，它返回一个用于读取操作的锁。
我们可以通过创建一个并发的字符串列表来快速演示它们的用法：

```
type list struct {
    m sync.RWMutex
    value []string
}
```

We can iterate the slice to find the selected value and use a read lock to delay the writing while we are reading:

我们可以迭代切片以找到选定的值，并在读取时使用读取锁延迟写入：

```
func (l *list) contains(v string) bool {
    for _, s := range l.value {
        if s == v {
            return true
        }
    }
    return false
}

func (l *list) Contains(v string) bool {
    l.m.RLock()
    found := l.contains(v)
    l.m.RUnlock()
    return found
}
```

We can use the write lock when adding new elements:

我们可以在添加新元素时使用写锁：

```
func (l *list) Add(v string) bool {
    l.m.Lock()
    defer l.m.Unlock()
    if l.contains(v) {
        return false
    }
    l.value = append(l.value, v)
    return true
}
```

Then we can try to use several goroutines to execute the same operation on the list:
然后我们可以尝试使用多个goroutine在列表上执行相同的操作：

```
var src = []string{
    "Ryu", "Ken", "E. Honda", "Guile",
    "Chun-Li", "Blanka", "Zangief", "Dhalsim",
}
var l list
for i := 0; i < 10; i++ {
    go func(i int) {
        for _, s := range src {
            go func(s string) {
                if !l.Contains(s) {
                    if l.Add(s) {
                        fmt.Println(i, "add", s)
                    } else {
                        fmt.Println(i, "too slow", s)
                    }
                }
            }(s)
        }
    }(i)
}
time.Sleep(500 * time.Millisecond)
```

We are checking whether the name is contained in the lock first, then we try to add the element. This causes more than one routine to attempt to add a new element, but since writing locks are exclusive, only one will succeed.
我们正在检查名称是否包含在锁中，然后尝试添加元素。这会导致多个例程尝试添加一个新元素，但由于写入锁是独占的，因此只有一个会成功。


# Write starvation        写“饥饿”

When designing an application, this kind of mutex is not always the obvious choice, because in a scenario where there is a greater number of read locks and a few write ones, the mutex will be accepting incoming more read locks after the first, letting the write operation wait for a moment where there are no read locks active. This is a phenomenon referred to as write starvation.

To check this out, we can define a type that has both a write and a read operation, which take some time, as shown in the following code:

在设计应用程序时，这种互斥体并不总是一个明显的选择，因为在读锁和写锁数量较多的情况下，互斥体将在第一个之后接受传入的更多读锁，让写操作在没有激活的读锁。这是一种被称为“写饥饿”的现象。
为了检查这个问题，我们可以定义一个既有写操作又有读操作的类型，这需要一些时间，如下面的代码所示：

```
type counter struct {
    m sync.RWMutex
    value int
}

func (c *counter) Write(i int) {
    c.m.Lock()
    time.Sleep(time.Millisecond * 100)
    c.value = i
    c.m.Unlock()
}

func (c *counter) Value() int {
    c.m.RLock()
    time.Sleep(time.Millisecond * 100)
    a := c.value
    c.m.RUnlock()
    return a
}


```

We can try to execute both write and read operations with the same cadence in separate goroutines, using a duration that is lower than the execution time of the methods (50 ms versus 100 ms). We will also check out how much time they spend in a locked state:

我们可以尝试在不同的goroutine中以相同的节奏执行写操作和读操作，使用的持续时间低于方法的执行时间（50 ms与100 ms）。我们还将检查他们在锁定状态下花费的时间：
                                                 
```
var c counter
t1 := time.NewTicker(time.Millisecond * 50)
time.AfterFunc(time.Second*2, t1.Stop)
for {
    select {
    case <-t1.C:
        go func() {
            t := time.Now()
            c.Value()
            fmt.Println("val", time.Since(t))
        }()
        go func() {
            t := time.Now()
            c.Write(0)
            fmt.Println("inc", time.Since(t))
        }()
    case <-time.After(time.Millisecond * 200):
        return
    }
}
```

If we execute the application, we see that for each write operation, more than one read is executed, and each next call is spending more time than the previous, waiting for the lock. This is not true for the read operation, which can happen at the same time, so as soon as a reader manages to lock the resource, all the other waiting readers will do the same. Replacing RWMutex with Mutex will make both operations have the same priority, as in the previous example.

如果我们执行应用程序，我们会看到对于每个写操作，都会执行多个读操作，并且每个下一个调用比前一个调用花费更多的时间，等待锁。读操作可能同时发生，这是不正确的，因此一旦读卡器成功锁定了资源，所有其他等待的读卡器都将执行相同的操作。将rwmutex替换为mutex将使两个操作具有与前一个示例相同的优先级。

# Locking gotchas 锁定gotchas

Some care must be taken when locking and unlocking mutexes in order to avoid unexpected behavior and deadlocks in the application. Take the following snippet:

在锁定和解锁互斥锁时必须注意一些，以避免应用程序中出现意外行为和死锁。取以下代码段：

```
for condition {
    mu.Lock()
    defer mu.Unlock()
    action()
}


```
This code seems okay at first sight, but it will inevitably block the goroutine. This is because the defer statement is not executed at the end of each loop iteration, but when the function returns. So the first attempt will lock without releasing and the second attempt will remain stuck.

A little refactor can help fix this, as shown in the following snippet:

这段代码乍一看似乎还可以，但它不可避免地会阻止Goroutine。这是因为defer语句不是在每个循环迭代结束时执行的，而是在函数返回时执行的。因此，第一次尝试将锁定而不释放，第二次尝试将保持卡住。
一点重构可以帮助解决这个问题，如下面的代码片段所示：


```
for condition {
    func() {
        mu.Lock()
        defer mu.Unlock()
        action()
    }()
}

```

We can use a closure to be sure that the deferred Unlock gets executed, even if action panics.

If the kind of operations that are executed on the mutex will not cause panic, it can be a good idea to ditch the defer and just use it after executing the action, as follows:

我们可以使用一个闭包来确保延迟的解锁得到执行，即使操作非常紧急。
如果在互斥体上执行的操作类型不会导致恐慌，那么最好放弃延迟并在执行操作后使用它，如下所示：

```
for condition {
    mu.Lock()
    action()
    mu.Unlock()
}
```

defer has a cost, so it is better to avoid it when it is not necessary, such as when doing a simple variable read or assignment.
延迟有一定的成本，所以最好在不必要时避免延迟，例如在执行简单变量读取或赋值时。

# Synchronizing goroutines 同步Goroutines

Until now, in order to wait for goroutines to finish, we used a channel of empty structures and sent a value through the channel as the last operation, as follows:

到目前为止，为了等待Goroutines完成，我们使用了一个包含空结构的通道，并通过通道发送一个值作为最后一个操作，如下所示：

```
ch := make(chan struct{})
for i := 0; i < n; n++ {
    go func() {
        // do something
        ch <- struct{}{}
    }()
}
for i := 0; i < n; n++ {
    <-ch
}
```

This strategy works, but it's not the preferred way to accomplish the task. It's not correct semantically, because we are using a channel, which is a tool for communication, to send empty data. This use case is about synchronization rather than communication. That's why there is the sync.WaitGroup data structure, which covers such cases. It has a main status, called counter, which represents the number of elements waiting:
这个策略是可行的，但它不是完成任务的首选方式。这在语义上是不正确的，因为我们使用一个通道来发送空数据，这是一种通信工具。
这个用例是关于同步而不是通信的。这就是为什么有sync.waitgroup数据结构，它涵盖了这种情况。它有一个称为counter的主状态，它表示等待的元素数：


```
type WaitGroup struct {
    noCopy noCopy
    state1 [3]uint32
}
```

The noCopy field prevents the structure from being copied by value with panic. The state is an array made by three int32, but only the first and last entries are used; the remaining one is used for compiler optimizations.

The WaitGroup offers three methods to accomplish the same result:

Add: This changes the value of the counter using the given value, which could also be negative. If the counter goes under zero, the application panics.
Done: This is a shorthand for Add with -1 as the argument. It is usually called when a goroutine finishes its job to decrement the counter by 1.
Wait: This operation blocks the current goroutine until the counter reaches zero.
Using the wait group results in a much cleaner and more readable code, as we can see in the following example:

```
nocopy字段防止按值复制结构时出现恐慌。状态是由三个int32组成的数组，但只使用第一个和最后一个条目；其余条目用于编译器优化。
WaitGroup提供了三种方法来实现相同的结果：
加：这将使用给定值更改计数器的值，该值也可能是负数。如果计数器低于零，应用程序将恐慌。
完成：这是用-1作为参数的add的简写。当一个goroutine完成它的工作时，它通常被称为将计数器递减1。
等待：此操作将阻止当前goroutine，直到计数器达到零。
使用wait组可以得到更清晰、更可读的代码，如下面的示例所示：
 
func main() {
    wg := sync.WaitGroup{}
    wg.Add(10)
    for i := 1; i <= 10; i++ {
        go func(a int) {
            for i := 1; i <= 10; i++ {
                fmt.Printf("%dx%d=%d\n", a, i, a*i)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

In the preceding example, we have a 10% chance of finishing each iteration of the for loop, so we are adding one to the group before starting the goroutine.

A very common error is to add the value inside the goroutine, which usually results in a premature exit without any goroutines executed. This happens because the application creates the goroutines and executes the Wait function before the routines start and add their own delta, as in the following example:

在前面的示例中，我们有10%的机会完成for循环的每个迭代，因此我们在开始Goroutine之前向组中添加一个。
一个非常常见的错误是在goroutine中添加值，这通常会导致过早退出而不执行任何goroutine。这是因为应用程序在例程启动和添加自己的delta之前创建goroutine并执行wait函数，如下例所示：

```
func main() {
    wg := sync.WaitGroup{}
    for i := 1; i < 10; i++ {
        go func(a int) {
            wg.Add(1)
            for i := 1; i <= 10; i++ {
                fmt.Printf("%dx%d=%d\n", a, i, a*i)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

This application will not print anything because it arrives at the Wait statement before any goroutine is started and the Add method is called.
此应用程序将不打印任何内容，因为它在启动任何goroutine并调用add方法之前到达wait语句。


# Singleton in Go 单例模式

The singleton pattern is a commonly used strategy for software development. This involves  restricting the number of instances of a certain type to one, using the same instance across the whole application. A very simple implementation of the concept could be the following code:
单例模式是软件开发中常用的策略。这涉及到在整个应用程序中使用相同的实例，将特定类型的实例数限制为一个。概念的一个非常简单的实现可以是以下代码：

```
type obj struct {}

var instance *obj

func Get() *obj{
    if instance == nil {
        instance = &obj{}
    }
    return instance
}
```

This is perfectly fine in a consecutive scenario but in a concurrent one, like in many Go applications, this is not thread-safe and could generate race conditions.

The previous example could be made thread-safe by adding a lock that would avoid any race condition, as follows:

这在连续的场景中是非常好的，但是在并发的场景中，就像在许多GO应用程序中一样，这不是线程安全的，可能会生成争用条件。
前面的示例可以通过添加一个可以避免任何竞争条件的锁来保证线程安全，如下所示：

```
type obj struct {}

var (
    instance *obj
    lock     sync.Mutex
)

func Get() *obj{
    lock.Lock()
    defer lock.Unlock()
    if instance == nil {
        instance = &obj{}
    }
    return instance
}
```

This is safe, but slower, because Mutex will be synchronizing each time the instance is requested.

The best solution to implement this pattern, as shown in the following example, is to use the sync.Once struct that takes care of executing a function once using a combination of Mutex and atomic readings (which we will see in the second part of the chapter):

```
这是安全的，但速度较慢，因为每次请求实例时，互斥体都将进行同步。
如下面的示例所示，实现此模式的最佳解决方案是使用sync.once结构，该结构负责执行一个函数，使用互斥和原子读取的组合（我们将在本章的第二部分中看到）：

type obj struct {}

var (
    instance *obj
    once     sync.Once
)

func Get() *obj{
    once.Do(func(){
        instance = &obj{}
    })
    return instance
}
```

The resulting code is idiomatic and clear, and has better performance compared to the mutex solution. Since the operation will be executed just the once, we can also get rid of the nil check we were doing on the instance in the previous examples.

生成的代码是惯用的、清晰的，与互斥解决方案相比，它具有更好的性能。由于该操作将只执行一次，因此我们还可以去掉前面示例中对实例所做的nil检查。



# Once and Reset

The sync.Once function is made for executing another function once and no more. There is a very useful third-party library, which allows us to reset the state of the singleton using the Reset method.

The package source code can be found at: github.com/matryer/resync.

Typical uses include some initialization that needs to be done again on a particular error, such as obtaining an API key or dialing again if the connection disrupts.

一次复位
sync.once函数用于执行另一个函数一次或多次。有一个非常有用的第三方库，它允许我们使用reset方法重置singleton的状态。
包源代码位于：github.com/matryer/resync。
典型的使用包括一些在出现特定错误时需要重新进行的初始化，例如，获取API密钥或在连接中断时再次拨号。

# Resource recycling
We have already seen how to implement resource recycling, with a buffered channel with a pool of workers, in the previous chapter. There will be two methods as follows:

A Get method that tries to receive a message from the channel or return a new instance.
A Put method that tries to return an instance back to a channel or discard it.
This is a simple implementation of a pool with channels:

资源回收
在上一章中，我们已经看到了如何通过一个缓冲通道和一个工人池来实现资源回收。有以下两种方法：
尝试从通道接收消息或返回新实例的get方法。
一种Put方法，尝试将实例返回到通道或丢弃它。
这是一个带有通道的池的简单实现：

```
type A struct{}

type Pool chan *A

func (p Pool) Get() *A {
    select {
    case a := <-p:
        return a
    default:
        return new(A)
    }
}

func (p Pool) Put(a *A) {
    select {
    case p <- a:
    default:
    }
}
```

We can improve this using the sync.Pool structure, which implements a thread-safe set of objects that can be saved or retrieved. The only thing that needs to be defined is the behavior of the pool when creating a new object:

我们可以使用sync.pool结构来改进这一点，它实现了一组线程安全的对象，这些对象可以保存或检索。唯一需要定义的是创建新对象时池的行为：

```
type Pool struct {
    // New optionally specifies a function to generate
    // a value when Get would otherwise return nil.
    // It may not be changed concurrently with calls to Get.
    New func() interface{}
    // contains filtered or unexported fields
}
```


The pool offers two methods: Get and Put. These methods return an object from the pool (or create a new one) and place the object back in the pool. Since the Get method returns an interface{}, the value needs to be cast to the specific type in order to be used correctly. We talked extensively about buffer recycling and in the following example, we will try to implement one using sync.Pool.

We will need to define the pool and functions to obtain and release new buffers. Our buffers will have an initial capacity of 4 KB, and the Put function will ensure that the buffer is reset before putting it back in the pool, as shown in the following code example:

```
游泳池提供两种方法：接球和接球。这些方法从池中返回一个对象（或创建一个新的对象），并将该对象放回池中。由于get方法返回接口，因此需要将该值强制转换为特定类型才能正确使用。我们广泛讨论了缓冲区回收，在下面的示例中，我们将尝试使用sync.pool实现一个缓冲区回收。
我们需要定义池和函数来获取和释放新的缓冲区。我们的缓冲区的初始容量为4kb，put函数将确保在将缓冲区放回池之前重置缓冲区，如下面的代码示例所示：

var pool = sync.Pool{
    New: func() interface{} {
        return bytes.NewBuffer(make([]byte, 0, 4096))
    },
}

func Get() *bytes.Buffer {
    return pool.Get().(*bytes.Buffer)
}

func Put(b *bytes.Buffer) {
    b.Reset()
    pool.Put(b)
}
```


Now we will create a series of goroutines, which will use the WaitGroup to signal when they're done, and will do the following: 

Wait a certain amount of time (1-5 seconds).
Acquire a buffer.
Write information on the buffer.
Copy the content to the standard output.
Release the buffer.
We will use a sleep time equal to 1 second, plus another second every 4 iterations of the loop, up to 5:


```
现在，我们将创建一系列goroutine，它将使用waitgroup在完成时发出信号，并将执行以下操作：
等待一定时间（1-5秒）。
获取缓冲区。
在缓冲区上写入信息。
将内容复制到标准输出。
释放缓冲器。
我们将使用等于1秒的睡眠时间，再加上每4次循环的另一秒，最多5次：

start := time.Now()
wg := sync.WaitGroup{}
wg.Add(20)
for i := 0; i < 20; i++ {
    go func(v int) {
        time.Sleep(time.Second * time.Duration(1+v/4))
        b := Get()
        defer func() {
            Put(b)
            wg.Done()
        }()
        fmt.Fprintf(b, "Goroutine %2d using %p, after %.0fs\n", v, b, time.Since(start).Seconds())
        fmt.Printf("%s", b.Bytes())
    }(i)
}
wg.Wait()
```

The information in print also contains the buffer memory address. This will help us to confirm that the buffers are always the same and no new ones are created.

打印信息还包含缓冲区内存地址。这将帮助我们确认缓冲区总是相同的，并且没有创建新的缓冲区。

# Slices recycling issues 切片回收问题

With data structure with an underlying byte slice, such as bytes.Buffer, we should be careful when using them combined with sync.Pool or a similar mechanism of recycling. Let's change the previous example and collect the buffer's bytes instead of printing them to standard output. The following is an example code for this:
对于具有底层字节片（如bytes.buffer）的数据结构，在将它们与sync.pool或类似的回收机制结合使用时，应该小心。让我们更改前面的示例，并收集缓冲区的字节，而不是将它们打印到标准输出。下面是一个示例代码：

```
var (
    list = make([][]byte, 20)
    m sync.Mutex
)
for i := 0; i < 20; i++ {
    go func(v int) {
        time.Sleep(time.Second * time.Duration(1+v/4))
        b := Get()
        defer func() {
            Put(b)
            wg.Done()
        }()
        fmt.Fprintf(b, "Goroutine %2d using %p, after %.0fs\n", v, b, time.Since(start).Seconds())
        m.Lock()
        list[v] = b.Bytes()
        m.Unlock()
    }(i)
}
wg.Wait()

```


So, what happens when we print the list of byte slices? We can see this in the following example:

```go
那么，当我们打印字节片列表时会发生什么？我们可以在下面的示例中看到这一点：

for i := range list {
    fmt.Printf("%d - %s", i, list[i])
}

```

We get an unexpected result as the buffers have been overwritten. That's because the buffers are reusing the same underlying slice and overriding the content with every new usage.A solution to this problem is usually to execute a copy of the bytes, instead of just assigning them:

由于缓冲区被覆盖，我们得到了一个意外的结果。这是因为缓冲区正在重用同一个基础切片，并用每种新用法覆盖内容。解决此问题的方法通常是执行字节的副本，而不是只分配它们：

```
m.Lock()
list[v] = make([]byte, b.Len())
copy(list[v], b.Bytes())
m.Unlock()
```

# Conditions 条件

In concurrent programming, a condition variable is a synchronization mechanism that contains threads waiting for the same condition to verify. In Go, this means that there are some goroutines waiting for something to occur. We already did an implementation of this using channels with a single goroutine waiting, as shown in the following example:

``` 
在并发编程中，条件变量是一种同步机制，其中包含等待验证相同条件的线程。在Go中，这意味着有一些Goroutine在等待发生什么。我们已经使用具有单个goroutine等待的通道实现了这一点，如下面的示例所示：

ch := make(chan struct{})
go func() {
    // do something
    ch <- struct{}{}
}()
go func() {
    // wait for condition
    <-ch
    // do something else
}
```

This approach is limited to a single goroutine, but it can be improved to support more listeners switching from message-sending to closing down the channel:

```
此方法仅限于单个goroutine，但可以改进为支持更多的侦听器从消息发送切换到关闭通道：

go func() {
    // do something
    close(ch)
}()
for i := 0; i < n; i++ {
    go func() {
        // wait for condition
        <-ch
        // do something else
    }()
}
```

Closing the channel works for more than one listener, but it does not allow them to use the channel any further after it closes.

The sync.Cond type is a tool that makes it possible to handle all this behavior in a better way. It uses a locker in its implementation and exposes three methods:

Broadcast: This wakes all goroutines waiting for the condition.
Signal: This wakes a single goroutine waiting for the condition, if there is at least one.
Wait: This unlocks the locker, suspends execution of the goroutine, and later resumes the execution and locks it again, waiting for a Broadcast or Signal.
It is not required, but the Broadcast and Signal operations can be done while holding the locker, locking it before and releasing it after. The Wait method requires holding the locker before calling and unlocking it after the condition has been used.

Let's create a concurrent application which uses sync.Cond to orchestrate more goroutines. We will have a prompt from the command line, and each record will be written to a series of files. We will have a main structure that holds all the data:

```
关闭通道对多个侦听器有效，但它不允许在关闭通道后再进一步使用通道。
cond类型是一种工具，可以更好地处理所有这些行为。它在实现中使用了一个锁，并公开了三种方法：
广播：这会唤醒所有等待状态的Goroutine。
信号：如果至少有一个Goroutine，它会唤醒等待状态的单个Goroutine。
等待：这会解锁储物柜，暂停Goroutine的执行，稍后恢复执行并再次锁定，等待广播或信号。
不需要，但广播和信号操作可以在按住储物柜、锁定储物柜前和释放储物柜后进行。wait方法要求在调用前保存储物柜，并在使用条件后将其解锁。
让我们创建一个使用sync.cond协调更多goroutine的并发应用程序。我们将从命令行得到一个提示，并且每个记录将被写入一系列文件中。我们将拥有一个保存所有数据的主要结构：


type record struct {
    sync.Mutex
    buf string
    cond *sync.Cond
    writers []io.Writer
}
```

The condition we will be monitoring is a change in the buf field. In the Run method, the record structure will start several goroutines, one for each writer. Each goroutine will be waiting for the condition to trigger and will write in its file:

```
我们要监控的情况是buf字段发生了变化。在run方法中，记录结构将启动几个goroutine，每个编写器一个。每个goroutine都将等待触发条件，并将在其文件中写入：

func (r *record) Run() {
    for i := range r.writers {
        go func(i int) {
            for {
                r.Lock()
                r.cond.Wait()
                fmt.Fprintf(r.writers[i], "%s\n", r.buf)
                r.Unlock()
            }
        }(i)
    }
}
```

We can see that we lock the condition before using Wait, and we unlock it after using the value that our condition refers to. The main function will create a record and a series of files, according to the command-line arguments provided:

```
我们可以看到，我们在使用wait之前锁定了条件，在使用条件引用的值之后解锁了条件。根据提供的命令行参数，主函数将创建一个记录和一系列文件：

// let's make sure we have at least a file argument
if len(os.Args) < 2 {
    log.Fatal("Please specify at least a file")
}
r := record{
    writers: make([]io.Writer, len(os.Args)-1),
}
r.cond = sync.NewCond(&r)
for i, v := range os.Args[1:] {
    f, err := os.Create(v)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    r.writers[i] = f
}
r.Run()
```

We will then use bufio.Scanner to read lines and broadcast the change of the buf field. We will also accept a special value, \q, as a quit command:


```
然后我们将使用bufio.scanner读取行并广播buf字段的更改。我们还将接受一个特殊值\q，作为退出命令：


scanner := bufio.NewScanner(os.Stdin)
for {
    fmt.Printf(":> ")
    if !scanner.Scan() {
        break
    }
    r.Lock()
    r.buf = scanner.Text()
    r.Unlock()
    switch {
    case r.buf == `\q`:
        return
    default:
        r.cond.Broadcast()
    }
}
```

We can see that the change of buf is done while holding the lock and this is followed by the call to Broadcast, which wakes up all the goroutines waiting for the condition.

我们可以看到，buf的更改是在保持锁的同时完成的，然后是广播调用，它会唤醒等待条件的所有goroutine。

# Synchronized maps 同步地图

Built-in maps in Go are not thread-safe, and, therefore, trying to write from different goroutines can cause a runtime error: concurrent map writes. We can verify this using a simple program that tries to make changes concurrently:

Go中的内置映射不是线程安全的，因此，尝试从不同的goroutine写入可能会导致运行时错误：并发映射写入。我们可以使用一个尝试同时进行更改的简单程序来验证这一点：

```
func main() {
    var m = map[int]int{}
    wg := sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(i int) {
            m[i%5]++
            fmt.Println(m)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

Reading while writing is also a runtime error, concurrent map iteration and map write, which we can see by running the following example:

写时读也是一个运行时错误、并发映射迭代和映射写入，我们可以通过运行以下示例看到这一点：

```
func main() {
    var m = map[int]int{}
    var done = make(chan struct{})
    go func() {
        for i := 0; i < 100; i++ {
            time.Sleep(time.Nanosecond)
            m[i]++
        }
        close(done)
    }()
    for {
        time.Sleep(time.Nanosecond)
        fmt.Println(len(m), m)
        select {
        case <-done:
            return
        default:
        }
    }
}


```

Sometimes, trying to iterate a map (as the Print statement does) can cause panic such as index out of range, because the internal slices may have been allocated somewhere else.

A very easy strategy to make a map concurrent is to couple it with sync.Mutex or sync.RWMutex. This makes it possible to lock the map when executing the operations:

```
有时，尝试迭代映射（如print语句所做的）会导致恐慌，例如索引超出范围，因为内部切片可能已分配到其他地方。
使映射并发的一个非常简单的策略是将它与sync.mutex或sync.rwmutex结合起来。这样可以在执行操作时锁定映射：

type m struct {
    sync.Mutex
    m map[int]int
}

```

We use the map for getting or setting the value, such as the following, for instance:


```
我们使用映射来获取或设置值，例如：

func (m *m) Get(key int) int {
    m.Lock()
    a := m.m[key]
    m.Unlock()
    return a
}

func (m *m) Put(key, value int) {
    m.Lock()
    m.m[key] = value
    m.Unlock()
}
```

We can also pass a function that takes a key-value pair and executes it for each tuple, while locking the map:

```
我们还可以传递一个接受键值对的函数，并为每个元组执行它，同时锁定映射：

func (m *m) Range(f func(k, v int)) {
    m.Lock()
    for k, v := range m.m {
        f(k, v)
    }
    m.Unlock()
}
```


Go 1.9 introduced a structure called sync.Map that does exactly this. It is a very generic map[interface{}]interface{}, which makes it possible to execute thread-safe operations using the following methods:

Load: Gets a value from the map for the given key.
Store: Sets a value in the map for the given key.
Delete: Removes the entry for the given key from the map.
LoadOrStore: Returns the value for the key, if present, or the stored value.
Range: Calls a function that returns a Boolean for each key-value pair in the map. The iteration stops if false is returned.
We can see how this works in the following snippet, in which we try to attempt several writes at the same time:

```
Go1.9引入了一个名为sync.map的结构，它可以做到这一点。它是一个非常通用的映射[接口]接口，这使得使用以下方法执行线程安全操作成为可能：
加载：从映射中获取给定键的值。
存储：在映射中为给定键设置一个值。
删除：从映射中删除给定键的条目。
loadorstore：返回键的值（如果存在）或存储的值。
range：调用一个函数，该函数为映射中的每个键值对返回一个布尔值。如果返回false，则迭代将停止。
我们可以在下面的代码片段中看到这是如何工作的，其中我们尝试同时进行多个写入操作：


func main() {
    var m = sync.Map{}
    var wg = sync.WaitGroup{}
    wg.Add(1000)
    for i := 0; i < 1000; i++ {
        go func(i int) {
            m.LoadOrStore(i, i)
            wg.Done()
        }(i)
    }
    wg.Wait()
    i := 0
    m.Range(func(k, v interface{}) bool {
        i++
        return true
    })
   fmt.Println(i)
}
```

This application, unlike the version with a regular Map, does not crash and executes all the operations.
这个应用程序与带有常规映射的版本不同，它不会崩溃并执行所有操作。


# Semaphores 信号量

In the previous chapter, we saw how it is possible to use channels to create weighted semaphores. There is a better implementation in the experimental sync package. This can be found at: golang.org/x/sync/semaphore.

This implementation makes it possible to create a new semaphore, specifying the weight with semaphore.NewWeighted.

Quotas can be acquired using the Acquire method, specifying how many quotas you want to acquire. These can be released using the Release method, as shown in the following example:

```
在上一章中，我们看到了如何使用通道来创建加权信号量。在实验同步包中有更好的实现。这可以在：golang.org/x/sync/semaphore找到。
这个实现可以创建一个新的信号量，用semaphore.newweighted指定权重。
可以使用Acquire方法获取配额，指定要获取的配额数。这些可以使用释放方法释放，如下面的示例所示：

func main() {
    s := semaphore.NewWeighted(int64(10))
    ctx := context.Background()
    for i := 0; i < 20; i++ {
        if err := s.Acquire(ctx, 1); err != nil {
            log.Fatal(err)
        }
        go func(i int) {
            fmt.Println(i)
            s.Release(1)
        }(i)
    }
    time.Sleep(time.Second)
}

```

Acquiring quotas requires another argument besides the number, which is context.Context. This is another concurrency tool available in Go and we are going to see how to use this in the next chapter.

获取配额需要数字之外的另一个参数，即context.context。这是Go中可用的另一个并发工具，我们将在下一章中了解如何使用它。

# Atomic operations

The sync package delivers synchronization primitives, and, under the hood, it is using thread-safe operations on integers and pointers. We can find these functionalities in another package called sync/atomic, which can be used to create tools specific to the user use case, with better performance and less memory usage.
  
原子操作
同步包提供了同步原语，并且在后台，它对整数和指针使用线程安全操作。
我们可以在另一个名为sync/atomic的包中找到这些功能，该包可用于创建特定于用户用例的工具，具有更好的性能和更少的内存使用。

# Integer operations

There is a series of functions for pointers to the different types of integers:

int32   
int64   
uint32  
uint64  
uintptr 
This includes a specific type of integer that represents a pointer, uintptr. The operation available for these types are as follows:

Load: Retrieves the integer value from the pointer
Store: Stores the integer value in the pointer
Add: Adds the specified delta to the pointer value
Swap: Stores a new value in the pointer and returns the old one
CompareAndSwap: Swaps the new value for the old one only if this is the same as the specified one

```
整数运算
对于指向不同类型整数的指针，有一系列函数：
英特32
英特64
UIT32
UIT64
uintptr
这包括表示指针uintptr的特定整数类型。可用于这些类型的操作如下：
加载：从指针中检索整数值
存储：将整数值存储在指针中
添加：将指定的增量添加到指针值
交换：在指针中存储新值并返回旧值
CompareAndSwap：仅当新值与指定值相同时，才交换旧值。
```

# clicker
This function can be very helpful for defining thread-safe components really easily. A very obvious example could be a simple integer counter that uses Add to change the counter, Load to retrieve the current value, and Store to reset it:
```
这个函数对于很容易地定义线程安全组件非常有用。一个非常明显的例子是一个简单的整数计数器，它使用add来更改计数器，加载以检索当前值，并存储以重置它：

type clicker int32

func (c *clicker) Click() int32 {
    return atomic.AddInt32((*int32)(c), 1)
}

func (c *clicker) Reset() {
    atomic.StoreInt32((*int32)(c), 0)
}

func (c *clicker) Value() int32 {
    return atomic.LoadInt32((*int32)(c))
}
```

We can see it in action in a simple program, which tries to read, write, and reset the counter concurrently. 

We define the clicker and WaitGroup and add the correct number of elements to the wait group as follows:

```
我们可以在一个简单的程序中看到它的运行，这个程序试图同时读取、写入和重置计数器。
我们定义了clicker和wait group，并向wait组添加正确数量的元素，如下所示：

c := clicker(0)
wg := sync.WaitGroup{}
// 2*iteration + reset at 5
wg.Add(21)
```

We can launch a bunch of goroutines doing different actions, such as: 10 reads, 10 adds, and a reset:

```
我们可以启动一组Goroutine执行不同的操作，例如：10次读取、10次添加和重置：

for i := 0; i < 10; i++ {
    go func() {
        c.Click()
        fmt.Println("click")
        wg.Done()
    }()
    go func() {
        fmt.Println("load", c.Value())
        wg.Done()
    }()
    if i == 0 || i%5 != 0 {
        continue
    }
    go func() {
        c.Reset()
        fmt.Println("reset")
        wg.Done()
    }()
}
wg.Wait()
```

We will see the clicker acting as it is supposed to, executing concurrent sums without race conditions.

我们将看到点击器按预期的方式工作，在没有竞争条件的情况下执行并发和。

# Thread-safe floats


The atomic package offers only primitives for integers, but since float32 and float64 are stored in the same data structure that int32 and int64 use, we use them to create an atomic float value.

The trick is to use the math.Floatbits functions to get the representation of a float as an unsigned integer and the math.Floatfrombits functions to transform an unsigned integer to a float. Let's see how this works with a float64:

```
原子包只为整数提供原语，但是由于float32和float64存储在int32和int64使用的相同数据结构中，因此我们使用它们来创建原子浮点值。
诀窍是使用math.floatBits函数获取浮点的无符号整数表示，并使用math.floatFromBits函数将无符号整数转换为浮点。让我们看看这是如何处理float64的：

type f64 uint64

func uf(u uint64) (f float64) { return math.Float64frombits(u) }
func fu(f float64) (u uint64) { return math.Float64bits(f) }

func newF64(f float64) *f64 {
    v := f64(fu(f))
    return &v
}

func (f *f64) Load() float64 {
  return uf(atomic.LoadUint64((*uint64)(f)))
}

func (f *f64) Store(s float64) {
  atomic.StoreUint64((*uint64)(f), fu(s))
}
```

Creating the Add function is a little bit more complicated. We need to get the value with Load, then compare and swap. Since this operation could fail because the load is an atomic operation and compare and swap (CAS) is another, we keep trying it until it succeeds in a loop:

```
创建add函数有点复杂。我们需要得到带负载的值，然后比较和交换。由于此操作可能会失败，因为加载是一个原子操作，而比较和交换（CA）是另一个操作，因此我们一直尝试此操作，直到它在循环中成功：
func (f *f64) Add(s float64) float64 {
    for {
        old := f.Load()
        new := old + s
        if f.CompareAndSwap(old, new) {
            return new
        }
    }
}

func (f *f64) CompareAndSwap(old, new float64) bool {
    return atomic.CompareAndSwapUint64((*uint64)(f), fu(old), fu(new))
}
```

# Thread-safe Boolean

We can also use int32 to represent a Boolean value. We can use the integer 0 as false, and 1 as true, creating a thread-safe Boolean condition:

```
我们也可以使用int32来表示一个布尔值。我们可以使用0作为假整数，1作为真整数，创建一个线程安全的布尔条件：

type cond int32

func (c *cond) Set(v bool) {
    a := int32(0)
    if v {
        a++
    }
    atomic.StoreInt32((*int32)(c), a)
}

func (c *cond) Value() bool {
    return atomic.LoadInt32((*int32)(c)) != 0
}
```

This will allow us to use the cond type as a thread-safe Boolean value.
这将允许我们使用cond类型作为线程安全布尔值。


# Pointer operations

Pointer variables in Go are stored in intptr variables, integers large enough to hold a memory address. The atomic package makes it possible to execute the same operations for other integers types. There is a package that allows unsafe pointer operations, which offers the unsafe.Pointer type that is used in atomic operations.

In the following example, we define two integer variables and their relative integer pointers. Then we execute a swap of the first pointer with the second:

```
go中的指针变量存储在intptr变量中，该变量是一个足够大以保存内存地址的整数。原子包使对其他整数类型执行相同的操作成为可能。存在允许不安全指针操作的包，该包提供原子操作中使用的不安全指针类型。
在下面的示例中，我们定义了两个整数变量及其相对整数指针。然后，我们用第二个指针交换第一个指针：

v1, v2 := 10, 100
p1, p2 := &v1, &v2
log.Printf("P1: %v, P2: %v", *p1, *p2)
atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&p1)), unsafe.Pointer(p2))
log.Printf("P1: %v, P2: %v", *p1, *p2)
v1 = -10
log.Printf("P1: %v, P2: %v", *p1, *p2)
v2 = 3
log.Printf("P1: %v, P2: %v", *p1, *p2)
```

After the swap, both pointers are now referring to the second variable; any change to the first value does not influence the pointers. Changing the second variable changes the value referred to by the pointers.
交换之后，两个指针现在都引用第二个变量；对第一个值的任何更改都不会影响指针。更改第二个变量会更改指针引用的值。

# Value

The simplest tool we can use is atomic.Value. This holds interface{} and makes it possible to read and write it with thread safety. It exposes two methods, Store and Load, which make it possible to set or retrieve the value. As it happens, for other thread-safe tools, sync.Value must not be copied after its first use.

We can try to have many goroutines to set and read the same value. Each load operation gets the latest stored value and there are no errors being raised by concurrency:

```
我们可以使用的最简单的工具是atomic.value。这样就可以保存接口，并且可以使用线程安全性读写它。它公开了存储和加载两种方法，这使得设置或检索值成为可能。事实上，对于其他线程安全工具，在第一次使用sync.value之后，不能复制它。
我们可以尝试使用许多goroutine来设置和读取相同的值。每个加载操作都会获取最新的存储值，并且不会出现并发引发的错误：

func main() {
    var (
        v atomic.Value
        wg sync.WaitGroup
    )
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("load", v.Load())
            wg.Done()
        }(i)
        go func(i int) {
            v.Store(i)
            fmt.Println("store", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

This is a very generic container; it can be used for any type of variable and the variable type should change from one to another. If the concrete type changes, it will make the method panic; the same thing applies to a nil empty interface.

这是一个非常通用的容器；它可以用于任何类型的变量，并且变量类型应该从一种类型更改为另一种类型。如果具体类型发生更改，则会使方法死机；同样的事情也适用于空接口为零。

# Under the hood

The sync.Value type stores its data in a non-exported interface, as shown by the source code:

sync.value类型将其数据存储在未导出的接口中，如源代码所示：

```
type Value struct {
    v interface{}
}
```

It uses a type of unsafe package to convert that structure into another one, which has the same data structure as an interface:

```
它使用一种不安全的包类型将该结构转换为另一种结构，该结构与接口具有相同的数据结构：
type ifaceWords struct {
    typ unsafe.Pointer
    data unsafe.Pointer
}
```

Two types with the same exact memory layout can be converted in this way, skipping the Go's type safety. This makes it possible to use atomic operations with the pointers and execute thread-safe Store and Load operations.

To get the lock for writing values, atomic.Value uses a compare and swap operation with the unsafe.Pointer(^uintptr(0)) value (which is 0xffffffff) in the type; it changes the value and replaces the type with the correct one. 

In the same way, the load operation loops until the type is different to 0xffffffff, before trying to read the value.

Using this expedient, atomic.Value is capable of storing and loading any value using other atomic operations.

两个具有相同精确内存布局的类型可以通过这种方式进行转换，从而跳过Go的类型安全性。这使得对指针使用原子操作和执行线程安全存储和加载操作成为可能。
若要获取写入值的锁，atomic.value将使用类型中不安全的.pointer（^uintptr（0））值（即0xffffffff）进行比较和交换操作；它将更改该值并用正确的值替换该类型。
同样，在尝试读取值之前，加载操作循环直到类型与0xffffffff不同。
使用这种权宜之计，atomic.value能够使用其他原子操作存储和加载任何值。

# Summary

In this chapter, we saw the tools that are available in the Go standard package for synchronization. They are located in two packages: sync, which provides high-level tools such as mutexes, and sync/atomic, which executes low-level operations.

First, we saw how to synchronize data using lockers. We saw how to use sync.Mutex to lock a resource regardless of the operation type, and sync.RWMutex to allow for concurrent readings and blocking writes. We should be careful using the second one because writes could be delayed by consecutive readings.

Next, we saw how to keep track of running operations in order to wait for the end of a series of goroutines, using sync.WaitGroup. This acts as a thread-safe counter for current goroutines and makes it possible to put the current goroutine to sleep until it reaches zero, using the Wait method.

Furthermore, we checked the sync.Once structure used to execute a functionality once, which allows the implementation of a thread-safe singleton, for instance. Then we used sync.Pool to reuse instances instead of creating new ones when possible. The only thing that a pool needs is the function that returns the new instance.

The sync.Condition struct represents a specific condition and uses a locker to change it, allowing a goroutine to wait for the change. This can be delivered to a single goroutine using Signal, or to all goroutines using Broadcast. The package also offers a thread-safe version of sync.Map.

Finally, we checked out the functionalities of atomic, which are mostly integer thread-safe operations: loading, saving, adding, swapping, and CAS. We saw also atomic.Value, which that makes it possible to change the value of an interface concurrently and does not allow it to change type after the first change.

The next chapter will be about the latest element introduced in Go concurrency: Context, which is an interface that handles deadlines, cancellations, and much more.

在本章中，我们看到了Go标准包中用于同步的工具。它们位于两个包中：提供互斥锁等高级工具的Sync和执行低级操作的Sync/Atomic。
首先，我们了解了如何使用储物柜同步数据。我们了解了如何使用sync.mutex锁定资源，而不管操作类型如何，以及使用sync.rwmutex允许并发读取和阻塞写入。我们应该小心使用第二个，因为连续读取可能会延迟写入。
接下来，我们了解了如何使用sync.waitgroup跟踪正在运行的操作，以便等待一系列goroutine的结束。这可以作为当前goroutine的线程安全计数器，并使用wait方法使当前goroutine可以休眠到零。
此外，我们还检查了用于执行一次功能的sync.once结构，例如，它允许实现线程安全的singleton。然后我们使用sync.pool重用实例，而不是在可能的时候创建新的实例。池唯一需要的是返回新实例的函数。
sync.condition结构表示一个特定的条件，并使用一个locker来更改它，允许goroutine等待更改。这可以通过信号发送到单个Goroutine，也可以通过广播发送到所有Goroutine。该包还提供了一个线程安全版本的sync.map。
最后，我们检查了原子的功能，这些功能主要是整数线程安全操作：加载、保存、添加、交换和CA。我们还看到了atomic.value，这使得可以同时更改接口的值，并且不允许在第一次更改之后更改类型。
下一章将介绍Go并发中引入的最新元素：Context，它是一个处理最后期限、取消等的接口。

# Questions
What's a race condition?
What happens when you try to execute read and write operations concurrently with a map?
What's the difference between Mutex and RWMutex?
Why are wait groups useful?
What's the main use of Once?
How can you use a Pool?
What's the advantage of using atomic operations?

问题
什么是比赛条件？
当您尝试与映射同时执行读写操作时会发生什么？
mutex和rwmutex有什么区别？
为什么等待组有用？
一次性的主要用途是什么？
你怎么用游泳池？
使用原子操作有什么好处？




# Coordination Using Context

This chapter is about the relatively new context package and its usage in concurrent programming. It is a very powerful tool by defining a unique interface that's used in many different places in the standard library, as well as in many third-party packages.

The following topics will be covered in this chapter:

Understanding what context is
Studying its usage in the standard library
Creating a package that uses context

```
本章介绍了相对较新的上下文包及其在并发编程中的用法。它是一个非常强大的工具，通过定义一个独特的接口，可以在标准库的许多不同地方以及许多第三方软件包中使用。
本章将讨论以下主题：
理解上下文是什么
研究其在标准图书馆中的应用
创建使用上下文的包
```

# Understanding context

Context is a relatively new component that entered the standard library in version 1.7. It is an interface for synchronization between goroutines that was used internally by the Go team and ended up being a core part of the language.

```
理解上下文
Context是一个在1.7版中进入标准库的相对较新的组件。它是Goroutines之间的同步接口，Goroutines由Go团队内部使用，最终成为语言的核心部分。
```

# The interface


The main entity in the package is Context itself, which is an interface. It has only four methods:

```
包中的主要实体是上下文本身，它是一个接口。它只有四种方法：

type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

Let's learn about these four methods here:

Deadline: Returns the time when the context should be cancelled, together with a Boolean that is false when there is no deadline

Done: Returns a receive-only channel of empty structs, which signals when the context should be cancelled

Err: Returns nil while the done channel is open; otherwise it returns the cause of the context cancellation

Value: Returns a value associated with a key for the current context, or nil if there's no value for the key

Context has many methods compared to the other interfaces of the standard library, which usually have one or two methods. Three of them are closely related:

Deadline is the time for cancelling
Done signals when the context is done
Err returns the cause of the cancellation
The last method, Value, returns the value associated with a certain key. The rest of the package is a series of functions that allow you to create different types of contexts. Let's go through the various functions that comprise the package and look at various tools for creating and decorating contexts.

```
让我们在这里了解这四种方法：
截止时间：返回应取消上下文的时间，以及在没有截止时间时为false的布尔值。
完成：返回空结构的只接收通道，该通道指示何时应取消上下文。
错误：当done通道打开时返回nil；否则返回上下文取消的原因。
value：返回与当前上下文的键关联的值，如果键没有值，则返回nil。
与标准库的其他接口相比，Context有许多方法，后者通常有一个或两个方法。其中三个密切相关：
最后期限是取消的时间
完成上下文完成时发出信号
err返回取消原因
最后一个方法value返回与某个键关联的值。包的其余部分是一系列函数，允许您创建不同类型的上下文。让我们来介绍组成包的各种功能，并查看用于创建和装饰上下文的各种工具。
```

# Default contexts

The TODO and Background functions return context.Context without the need for any input argument. The value that's returned is an empty context, though, their distinction is just semantic.

```
默认上下文
todo和background函数返回context.context，不需要任何输入参数。返回的值是一个空上下文，但是它们的区别只是语义上的。
```

# Background

Background is an empty context that doesn't get cancelled, hasn't got a deadline, and doesn't hold any values. It is mostly used by the main function as the root context or for testing purposes. The following is some example code for this context:

```
背景
background是一个空上下文，它不会被取消，也没有最后期限，并且不包含任何值。它主要由主函数用作根上下文或用于测试目的。下面是此上下文的一些示例代码：

func main() {
    ctx := context.Background()
    done := ctx.Done()
    for i :=0; ;i++{
        select {
        case <-done:
            return
        case <-time.After(time.Second):
            fmt.Println("tick", i)
        }
    }
}
```

We can see that, in the context of the example, the loop goes on infinitely because the context is never completed. 
我们可以看到，在示例的上下文中，循环无限地进行，因为上下文永远不会完成。

# TODO

TODO is another empty context that should be used when the scope of the context isn't clear or if the type of context isn't available yet. It is used in the exact same way as Background. As a matter of fact, under the hood, they are the same thing; the distinction is only semantical. If we look at the source code, they have the exact same definition:
```
TODO是另一个空上下文，应在上下文范围不明确或上下文类型尚不可用时使用。它的使用方式与背景完全相同。事实上，在引擎盖下，它们是相同的东西；区别只是语义上的。如果我们看一下源代码，它们有完全相同的定义：

var (
    background = new(emptyCtx)
    todo = new(emptyCtx)
)
```

These basic contexts can be extended using the other functions of the package. They will act as decorators and add more capabilities to them.
这些基本上下文可以使用包的其他函数进行扩展。他们将扮演装饰师的角色，并为他们增加更多的功能。

# Cancellation, timeout, and deadline

The context we looked at is never cancelled, but the package offers different options for adding this functionality.

``` 
取消、超时和截止时间
我们看到的上下文永远不会被取消，但是包提供了添加此功能的不同选项。
```

# Cancellation

The context.WithCancel decorator function gets a context and returns another context and a function called cancel. The returned context will be a copy of the context that has a different done channel (the channel that marks that the current context is done) that gets closed when the parent context does or when the cancel function is called – whatever happens first.

In the following example, we can see that we wait a few seconds before calling the cancel function, and the program terminates correctly. The value of Err is the context.Canceled variable:

```
取消
context.withcancel修饰函数获取一个上下文，并返回另一个上下文和一个名为cancel的函数。返回的上下文将是具有不同完成通道（标记当前上下文已完成的通道）的上下文的副本，该通道在父上下文完成或调用cancel函数时关闭（无论先发生什么）。
在下面的示例中，我们可以看到在调用cancel函数之前，我们等待了几秒钟，并且程序正确终止。err的值是上下文。已取消的变量：


func main() {
    ctx, cancel := context.WithCancel(context.Background())
    time.AfterFunc(time.Second*5, cancel)
    done := ctx.Done()
    for i := 0; ; i++ {
        select {
        case <-done:
            fmt.Println("exit", ctx.Err())
            return
        case <-time.After(time.Second):
            fmt.Println("tick", i)
        }
    }
}
```

#Deadline

context.WithDeadline is another decorator, which specifies a time deadline as time.Time, and applies it to another context. If there is already a deadline and it is earlier than the one provided, the specified one gets ignored. If the done channel is still open when the deadline is met, it gets closed automatically.

In the following example, we set the deadline to be 5 seconds from now and call cancel 10 seconds after. The deadline arrives before the cancellation and Err returns a context.DeadlineExceeded error:

```
最后期限
Context.WithDeadline是另一个修饰程序，它将时间期限指定为Time.Time，并将其应用于另一个上下文。如果已经有一个截止日期，并且它早于提供的截止日期，则指定的截止日期将被忽略。如果在截止日期满足时“完成”通道仍处于打开状态，则它将自动关闭。
在下面的示例中，我们将截止时间设置为5秒之后，然后在10秒之后调用Cancel。截止日期早于取消日期，err返回context.deadlineexceeded错误：

func main() {
    ctx, cancel := context.WithDeadline(context.Background(), 
         time.Now().Add(5*time.Second))
    time.AfterFunc(time.Second*10, cancel)
    done := ctx.Done()
    for i := 0; ; i++ {
        select {
        case <-done:
            fmt.Println("exit", ctx.Err())
            return
        case <-time.After(time.Second):
            fmt.Println("tick", i)
        }
    }
}
```

We can see that the preceding example behaves exactly as expected. It will print the tick statement each second a few times until the the deadline is met and the error is returned.

我们可以看到前面的示例的行为与预期完全一样。它将每隔几秒钟打印一次勾号语句，直到达到最后期限并返回错误。

# Timeout   

The last cancel-related decorator is context.WithTimeout, which allows you to specify a time.Duration together with the context and closes the done channel automatically when the timeout is passed.If there a deadline active, the new value applies only if it's earlier than the parent. We can look at a pretty identical example, beside the context definition, and get the same result that we got for the deadline example:

```
超时
最后一个与取消相关的装饰器是Context.WithTimeout，它允许您指定一个时间。Duration与上下文一起，并在超时时间过后自动关闭Done通道。如果存在活动的截止时间，则新值仅在早于父级时适用。我们可以在上下文定义旁边查看一个非常相同的示例，并得到与最后期限示例相同的结果：


func main() {
    ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
    time.AfterFunc(time.Second*10, cancel)
    done := ctx.Done()
    for i := 0; ; i++ {
        select {
        case <-done:
            fmt.Println("exit", ctx.Err())
            return
        case <-time.After(time.Second):
            fmt.Println("tick", i)
        }
    }
}
```

Keys and values 

The context.WithValue function creates a copy of the parent context that has the given key associated with the specified value. Its scope holds values that are relative to a single request while it gets processed and should not be used for other scopes, such as optional function parameters.

The key should be something that can be compared, and it's a good idea to avoid string values because two different packages using context could overwrite each other's values. The suggestion is to use user-defined concrete types such as struct{}. 

Here, we can see an example where we take a base context and we add a different value for each goroutine, using an empty struct as a key:

```
键和值
context.withValue函数创建具有与指定值关联的给定键的父上下文的副本。它的作用域在被处理时保留与单个请求相关的值，不应用于其他作用域，例如可选函数参数。
密钥应该是可以比较的，最好避免使用字符串值，因为使用上下文的两个不同包可能会覆盖彼此的值。建议使用用户定义的具体类型，如struct。
在这里，我们可以看到一个示例，其中我们使用一个空结构作为键为每个goroutine添加一个不同的值：

type key struct{}

func main() {
    ctx, canc := context.WithCancel(context.Background())
    wg := sync.WaitGroup{}
    wg.Add(5)
    for i := 0; i < 5; i++ {
        go func(ctx context.Context) {
            v := ctx.Value(key{})
            fmt.Println("key", v)
            wg.Done()
            <-ctx.Done()
            fmt.Println(ctx.Err(), v)
        }(context.WithValue(ctx, key{}, i))
    }
    wg.Wait()
    canc()
    time.Sleep(time.Second)
}
```


We can also see that cancelling the parent cancels the other contexts. Another valid key type could be exported pointer values, which won't be the same, even if the underlying data is:

```
我们还可以看到取消父级会取消其他上下文。另一个有效的键类型可以是导出的指针值，即使基础数据是：

type key *int

func main() {
	k := new(key)
	ctx, canc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			v := ctx.Value(k)
			fmt.Println("key", v, ctx.Value(new(key)))
			wg.Done()
			<-ctx.Done()
			fmt.Println(ctx.Err(), v)
		}(context.WithValue(ctx, k, i))
	}
	wg.Wait()
	canc()
	time.Sleep(time.Second)
}
```

We can see that defining a key pointer with the same underlying value doesn't return the expected value.
我们可以看到，定义具有相同基础值的键指针不会返回预期值。


# Context in the standard library

Now that we've covered the contents of the package, we will look at how to use them with the standard package or in an application. Context is used in a few functions and methods of standard packages, mostly network packages. Let's go over them now:

http.Server uses it with the Shutdown method so that it has full control over timeout or to cancel an operation.
http.Request allows you to set a context using the WithContext method. It also allows you to get the current context using Context.
In the net package, Listen, Dial, and Lookup have a version that uses Context to control deadlines and timeouts.
In the database/sql package, context is used to stop or timeout many different operations.


```
标准库中的上下文
既然我们已经介绍了包的内容，那么我们将研究如何在标准包或应用程序中使用它们。Context用于标准包的一些函数和方法，主要是网络包。现在我们来看看：
服务器将它与shutdown方法一起使用，这样它就可以完全控制超时或取消操作。
request允许您使用withContext方法设置上下文。它还允许您使用上下文获取当前上下文。
在网络包中，listen、dial和lookup的版本使用上下文来控制最后期限和超时。
在database/sql包中，上下文用于停止或超时许多不同的操作。
```

# HTTP requests

Before the introduction of the official package, each HTTP-related framework was using its own version of context to store data relative to HTTP requests. This resulted in fragmentation, and the reuse of handlers and middleware wasn't possible without rewriting the middleware or any specific binding code.

```
在引入正式包之前，每个与HTTP相关的框架都使用自己版本的上下文来存储与HTTP请求相关的数据。这导致了碎片化，如果不重写中间件或任何特定的绑定代码，就不可能重用处理程序和中间件。
```


# Passing scoped values

The introduction of context.Context in http.Request tries to address this issue by defining a single interface that can be assigned, recovered, and used in various handlers.

The downside is that a context isn't assigned automatically to a request, and context values cannot be recycled. There should be no really good reason to do that since the context should store data that's specific to a certain package or scope, and the packages themselves should be the only ones that are able to interact with them.

A good pattern is the usage of a unique unexported key type combined with auxiliary functions to get or set a certain value:

```
传递作用域值
在http.request中引入context.context试图通过定义一个可以在各种处理程序中分配、恢复和使用的单个接口来解决此问题。
缺点是上下文不会自动分配给请求，并且上下文值不能被回收。没有真正好的理由这样做，因为上下文应该存储特定于某个包或范围的数据，并且包本身应该是唯一能够与它们交互的。
一个好的模式是使用一个独特的未排序的键类型和辅助函数来获取或设置一个特定的值：

type keyType struct{}

var key = &keyType{}

func WithKey(ctx context.Context, value string) context.Context {
    return context.WithValue(ctx, key, value)
}

func GetKey(ctx context.Context) (string, bool) {
    v := ctx.Value(key)
    if v == nil {
        return "", false
    }
    return v.(string), true
}
```

A context request is the only case in the standard library where it is stored in the data structure with the WithContext method and it's accessed using the Context method. This has been done in order to not break the existing code, and maintain the promise of compatibility of Go 1.

在标准库中，上下文请求是唯一一种情况，在这种情况下，使用withContext方法将其存储在数据结构中，并使用上下文方法访问它。这样做是为了不破坏现有的代码，并保证Go1的兼容性。

# Request cancellation 请求取消

A good usage of context is for cancellation and timeout when you're executing an HTTP request using http.Client, which handles the interruption automatically from the context. The following example does exactly that:

当您使用http.client执行HTTP请求时，使用context的一个好用法是取消和超时，它会自动处理来自上下文的中断。下面的示例正好做到了这一点：

```
func main() {
	const addr = "localhost:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
	})

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			panic(err)
		}
	}()

	req, _ := http.NewRequest(http.MethodGet, "http://"+addr, nil)

	ctx, canc := context.WithTimeout(context.Background(), time.Second*2)
	defer canc()

	time.Sleep(time.Second)
	if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
		log.Fatal("超时",err)
	}
}
// 2019/08/28 09:28:28 超时Get http://localhost:8080: context deadline exceeded


```

The context cancellation method can also be used to interrupt the current HTTP request that's passed to a client. In a scenario where we are calling different endpoints and returning the first result that's received, it would be a good idea to cancel the other requests.

Let's create an application that runs a query on different search engines and returns the results from the quickest one, cancelling the others. We can create a web server that has a unique endpoint that answers back in 0 to 10 seconds:

```
上下文取消方法还可以用于中断传递给客户机的当前HTTP请求。在一个调用不同端点并返回收到的第一个结果的场景中，最好取消其他请求。
让我们创建一个应用程序，它在不同的搜索引擎上运行查询，并从最快的搜索引擎返回结果，取消其他搜索引擎。我们可以创建一个具有唯一终结点的Web服务器，该终结点将在0到10秒内回复：

const addr = "localhost:8080"
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    d := time.Second * time.Duration(rand.Intn(10))
    log.Println("wait", d)
    time.Sleep(d)
})
go func() {
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalln(err)
    }
}()
```

We can use a cancellable context for the requests, combined with a wait group to synchronize it with the end of the request. Each goroutine will create a request and try to send the result using a channel. Since we are only interested in the first one, we will use  sync.Once to limit it:

```
我们可以为请求使用一个可取消的上下文，并结合一个等待组，使其与请求的结尾同步。每个goroutine将创建一个请求，并尝试使用通道发送结果。因为我们只对第一个感兴趣，所以我们将使用同步。一旦限制它：

ctx, canc := context.WithCancel(context.Background())
ch, o, wg := make(chan int), sync.Once{}, sync.WaitGroup{}
wg.Add(10)
for i := 0; i < 10; i++ {
    go func(i int) {
        defer wg.Done()
        req, _ := http.NewRequest(http.MethodGet, "http://"+addr, nil)
        if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
            log.Println(i, err)
            return
        }
        o.Do(func() { ch <- i })
    }(i)
}
log.Println("received", <-ch)
canc()
log.Println("cancelling")
wg.Wait()
```

When this program runs, we will see that one of the requests is completed successfully and gets sent to the channel, while the others are either cancelled or ignored.

当这个程序运行时，我们将看到其中一个请求成功完成并被发送到通道，而其他请求则被取消或忽略。


# HTTP server

The net/http package has several uses of context, including stopping the listener or being part of a request.

HTTP服务器
NET/HTTP包有几种上下文用途，包括停止侦听器或作为请求的一部分。

# Shutdown

http.Server allows us to pass a context for the shutdown operation. This allows to us to use some of the context capabilities, such as cancelling and timeout. We can define a new server with its mux and a cancellable context:

关闭
服务器允许我们传递关闭操作的上下文。这允许我们使用一些上下文功能，例如取消和超时。我们可以定义一个具有mux和可取消上下文的新服务器：

```
mux := http.NewServeMux()
server := http.Server{
    Addr: ":3000",
    Handler: mux,
}
ctx, canc := context.WithCancel(context.Background())
defer canc()
mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
    canc()
})


```

We can launch the server in a separate goroutine:
我们可以在单独的goroutine中启动服务器：

```
go func() {
    if err := server.ListenAndServe(); err != nil {
        if err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }
}()

```

The context will be complete when the shutdown endpoint is called and the cancellation function is invoked. We can wait for that event and then use another context with a timeout to call the shutdown method:

```
当调用shutdown端点并调用cancellation函数时，上下文将完成。我们可以等待该事件，然后使用另一个具有超时的上下文调用Shutdown方法：

select {
case <-ctx.Done():
    ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
    defer canc()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalln("Shutdown:", err)
    } else {
        log.Println("Shutdown:", "ok")
    }
}
```

This will allow us to terminate the server effectively within the timeout, after which it will terminate with an error.
这将允许我们在超时内有效地终止服务器，在此之后，服务器将以错误终止。


# Passing values

Another usage of context in a server is as a propagation of values and cancellation between different HTTP handlers. Let's look at an example where each request has a unique key that is an integer. We will use a couple of functions that are similar to the example where we had values using integers. The generation of a new key will be done with atomic:

``` 
传递值
服务器中上下文的另一种用法是在不同的HTTP处理程序之间传播值和取消。让我们来看一个例子，其中每个请求都有一个整数的唯一键。我们将使用一些类似于示例的函数，其中我们有使用整数的值。新密钥的生成将使用Atomic完成：

type keyType struct{}

var key = &keyType{}

var counter int32

func WithKey(ctx context.Context) context.Context {
    return context.WithValue(ctx, key, atomic.AddInt32(&counter, 1))
}

func GetKey(ctx context.Context) (int32, bool) {
    v := ctx.Value(key)
    if v == nil {
        return 0, false
    }
    return v.(int32), true
}
```

Now, we can define another function that takes any HTTP handler and creates the context, if necessary, and adds the key to it:
现在，我们可以定义另一个函数，它接受任何HTTP处理程序并在必要时创建上下文，并向其中添加键：

```
func AssignKeyHandler(h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        if ctx == nil {
            ctx = context.Background()
        }
        if _, ok := GetKey(ctx); !ok {
            ctx = WithKey(ctx)
        }
        h.ServeHTTP(w, r.WithContext(ctx))
    }
}
```

By doing this, we can define a very simple handler that serves files under a certain root directory. This function will use the key from the context to log information correctly. It will also check that the file exists before trying to serve it:

``` 
通过这样做，我们可以定义一个非常简单的处理程序，为某个根目录下的文件提供服务。此函数将使用上下文中的键来正确记录信息。它还将检查文件是否存在，然后再尝试为其提供服务：

func ReadFileHandler(root string) http.HandlerFunc {
    root = filepath.Clean(root)
    return func(w http.ResponseWriter, r *http.Request) {
        k, _ := GetKey(r.Context())
        path := filepath.Join(root, r.URL.Path)
        log.Printf("[%d] requesting path %s", k, path)
        if !strings.HasPrefix(path, root) {
            http.Error(w, "not found", http.StatusNotFound)
            log.Printf("[%d] unauthorized %s", k, path)
            return
        }
        if stat, err := os.Stat(path); err != nil || stat.IsDir() {
            http.Error(w, "not found", http.StatusNotFound)
            log.Printf("[%d] not found %s", k, path)
            return
        }
        http.ServeFile(w, r, path)
        log.Printf("[%d] ok: %s", k, path)
    }
}
```


We can combine those handlers to serve content from different folders, such as the home user or the temporary directory:

```
我们可以组合这些处理程序来提供来自不同文件夹（如主用户或临时目录）的内容：

home, err := os.UserHomeDir()
if err != nil {
    log.Fatal(err)
}
tmp := os.TempDir()
mux := http.NewServeMux()
server := http.Server{
    Addr: ":3000",
    Handler: mux,
}

mux.Handle("/tmp/", http.StripPrefix("/tmp/", AssignKeyHandler(ReadFileHandler(tmp))))
mux.Handle("/home/", http.StripPrefix("/home/", AssignKeyHandler(ReadFileHandler(home))))
if err := server.ListenAndServe(); err != nil {
    if err != http.ErrServerClosed {
        log.Fatal(err)
    }
}

```

We are using http.StipPrefix to remove the first part of the path and obtain the relative path, and pass it to the handler underneath. The resulting server will use context to pass the key value between handlers – this allows us to create another similar handler and use the AssignKeyHandler function to wrap the handler and GetKey(r.Context()) to access the key inside our handler.

我们正在使用http.stiprefix删除路径的第一部分并获取相对路径，然后将其传递给下面的处理程序。结果服务器将使用context在处理程序之间传递键值——这允许我们创建另一个类似的处理程序，并使用assignkeyhandler函数包装处理程序，使用getkey（r.context（））访问处理程序内的键。

# TCP dialing 

The network package offers context-related functionalities, such as dialing cancellation when we're dialing or listening to incoming connections. It allows us to use the timeout and cancellation capabilities of context when dialing a connection.

TCP拨号
网络包提供与上下文相关的功能，例如拨号或监听传入连接时取消拨号。它允许我们在拨号连接时使用上下文的超时和取消功能。

# Cancelling a connection

In order to test out the usage of context in a TCP connection, we can create a goroutine with a TCP server that will wait a period of time before starting the listener:

```
取消连接
为了测试TCP连接中上下文的使用情况，我们可以使用TCP服务器创建一个goroutine，该服务器将在启动侦听器之前等待一段时间：

addr := os.Args[1]
go func() {
    time.Sleep(time.Second)
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalln("Listener:", addr, err)
    }
    c, err := listener.Accept()
    if err != nil {
        log.Fatalln("Listener:", addr, err)
    }
    defer c.Close()
}()

```

We can use a context with a timeout that's lower than the server waiting time. We have to use net.Dialer in order to use the context in a dial operation:

```
我们可以使用超时时间低于服务器等待时间的上下文。我们必须使用net.dialer才能在拨号操作中使用上下文：

ctx, canc := context.WithTimeout(context.Background(),   
    time.Millisecond*100)
defer canc()
conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", os.Args[1])
if err != nil {
    log.Fatalln("-> Connection:", err)
}
log.Println("-> Connection to", os.Args[1])
conn.Close()
```

The application will try to connect for a short time, but will eventually give up when the context expires, returning an error.

In a situation where you want to establish a single connection from a series of endpoints, context cancellation would be a perfect use case. All the connection attempts would share the same context, and the first connection that dials correctly would call the cancellation, stopping the other attempts. We will create a single server that is listening to one of the addresses we will try to call:

```
应用程序将尝试短时间连接，但最终会在上下文过期时放弃，返回错误。
在您希望从一系列端点建立单个连接的情况下，上下文取消将是一个完美的用例。所有连接尝试将共享相同的上下文，正确拨号的第一个连接将调用取消，停止其他尝试。我们将创建一个正在监听我们将尝试呼叫的其中一个地址的单一服务器：

list := []string{
    "localhost:9090",
    "localhost:9091",
    "localhost:9092",
}
go func() {
    listener, err := net.Listen("tcp", list[0])
    if err != nil {
        log.Fatalln("Listener:", list[0], err)
    }
    time.Sleep(time.Second * 5)
    c, err := listener.Accept()
    if err != nil {
        log.Fatalln("Listener:", list[0], err)
    }
    defer c.Close()
}()

```

Then, we can try to dial all three addresses and cancel the context as soon as one connects. We will use a WaitGroup to synchronize with the end of the goroutines:

然后，我们可以尝试拨打所有三个地址，并在一个连接时立即取消上下文。我们将使用waitgroup与goroutines的结尾同步：

```
ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
defer canc()
wg := sync.WaitGroup{}
wg.Add(len(list))
for _, addr := range list {
    go func(addr string) {
        defer wg.Done()
        conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", addr)
        if err != nil {
            log.Println("-> Connection:", err)
            return
        }
        log.Println("-> Connection to", addr, "cancelling context")
        canc()
        conn.Close()
    }(addr)
}
wg.Wait()
```

What we will see in the output of this program is one connection succeeding, followed by the cancellation error of the other attempt.
我们将在这个程序的输出中看到的是一个连接成功，随后是另一个尝试的取消错误。



# Database operations
We aren't looking at the sql/database package in this book, but for the sake of completion, it is worth mentioning that it uses context too. Most of its operations have a counterpart with context, for instance:
Beginning a new transaction
Executing a query
Pinging the database
Preparing a query
This concludes the packages in the standard library that use context. Next, we are going to try to use context to build a package to allow the user of that package to cancel requests.

```
数据库操作
我们在这本书中没有讨论SQL/Database包，但是为了完成这一工作，值得一提的是它也使用上下文。它的大多数操作都与上下文对应，例如：
开始新交易
执行查询
Ping数据库
准备查询
这将结束标准库中使用上下文的包。接下来，我们将尝试使用上下文来构建一个包，以允许该包的用户取消请求。
```

# Experimental packages
A notable example in the experimental package that uses context is one we've already looked at – semaphore. Now that we have a better understanding of what context is for, it should be pretty obvious why the acquire operation also takes a context in as an argument.

When creating our application, we can provide a context with a timeout or cancellation and act accordingly:

```
实验包
在使用上下文的实验包中，我们已经看到了一个值得注意的例子——信号量。既然我们已经对上下文的用途有了更好的理解，那么很明显为什么Acquire操作也将上下文作为参数。
在创建应用程序时，我们可以为上下文提供超时或取消，并相应地执行以下操作：


func main() {
    s := semaphore.NewWeighted(int64(5))
    ctx, canc := context.WithTimeout(context.Background(), time.Second)
    defer canc()
    wg := sync.WaitGroup{}
    wg.Add(20)
    for i := 0; i < 20; i++ {
        go func(i int) {
            defer wg.Done()
            if err := s.Acquire(ctx, 1); err != nil {
                fmt.Println(i, err)
                return
            }
            go func(i int) {
                fmt.Println(i)
                time.Sleep(time.Second / 2)
                s.Release(1)
            }(i)
        }(i)
    }
    wg.Wait()
}
```

Running this application will show that the semaphore is acquired for the first second, but after that the context expires and all the remaining operations fail.

运行此应用程序将显示在第一秒获取信号量，但之后上下文将过期，所有剩余的操作都将失败。


# Context in your application

context.Context is the perfect tool to integrate into your package or application if it has operations that could take a long time and a user can cancel them, or if they should have time limitations such timeouts or deadlines.

```
应用程序中的上下文
Context.Context是集成到您的包或应用程序中的完美工具，如果它有可能需要很长时间的操作，并且用户可以取消它们，或者如果它们应该有时间限制，例如超时或截止日期。


```

# Things to avoid

Even though the context scope has been made very clear by the Go team, developers have been using it in various ways – some less orthodox than others. Let's check out some of them and which alternatives there are, instead of resorting to context.

```
要避免的事情
尽管Go团队已经明确了上下文范围，但开发人员一直在以各种方式使用它——有些方式比其他方式不那么正统。让我们来看看其中的一些，以及有哪些选择，而不是诉诸于上下文。

``` 
 
 
Wrong types as keys
The first practice to avoid is the usage of built-in types as keys. This is problematic because they can be overwritten because two interfaces with the same built-in values are considered the same, as shown in the following example:
 
```
键类型错误
要避免的第一个实践是使用内置类型作为键。这是有问题的，因为它们可以被覆盖，因为具有相同内置值的两个接口被认为是相同的，如下面的示例所示：


func main() {
    var a interface{} = "request-id"
    var b interface{} = "request-id"
    fmt.Println(a == b)

    ctx := context.Background()
    ctx = context.WithValue(ctx, a, "a")
    ctx = context.WithValue(ctx, b, "b")
    fmt.Println(ctx.Value(a), ctx.Value(b))
}
```
 
The first print instruction outputs true, and since the keys are compared by value, the second assignment shadows the first, resulting in the values for both keys being the same. A potential solution to this is to use an empty struct custom type, or an unexported pointer to a built-in value.

第一个打印指令输出为真，由于键是按值比较的，第二个赋值将第一个赋值隐藏起来，导致两个键的值相同。可能的解决方案是使用空结构自定义类型，或者使用指向内置值的未导入指针。 
 

# Passing parameters
It might so happen that you need to travel a long way through a stack of function calls. A very tempting solution would be to use a context to store that value and recall it only in the function that needs it. It is generally not a good idea to hide a required parameter that should be passed explicitly. It results in less readable code because it won't make it clear what influences the execution of a certain function.

It is still much better to pass the function down the stack. If the parameters list is getting too long, then it could be grouped into one or more structs in order to be more readable.

Let's have a look at the following function:

```
传递参数
可能会发生这样的情况，您需要在一堆函数调用中走很远的路。一个非常诱人的解决方案是使用一个上下文来存储这个值，并且只在需要它的函数中调用它。通常，隐藏应该显式传递的必需参数不是一个好主意。它会导致代码的可读性降低，因为它无法明确影响某个函数执行的因素。
在堆栈中传递函数仍然更好。如果参数列表太长，那么可以将其分组为一个或多个结构，以便更易于阅读。
让我们看一下以下函数：

func SomeFunc( ctx context.Context, 
    name, surname string, age int, 
    resourceID string, resourceName string) {}
```

The parameters could be grouped in the following way:

参数可以按以下方式分组：

```
type User struct {
    Name string
    Surname string
    Age int
}

type Resource struct {
    ID string
    Name string
}

func SomeFunc(ctx context.Context, u User, r Resource) {}
```


Optional arguments
Context should be used to pass optional parameters around, and also used as a sort of catch-all, like Python kwargs or JavaScript arguments. Using context as a substitute for behaviors can be very problematic because it could cause the shadowing of variables, like we saw in the example of context.WithValue.

Another big drawback of this approach is hiding what's happening and making the code more obscure. A much better approach when it comes to optional values is using a pointer to structure arguments – this allows you to avoid passing the structure at all with nil.

Let's say you had the following code:

```
可选参数
上下文应该用来传递可选参数，也可以用作一种“全部捕获”，比如python kwargs或javascript参数。使用上下文作为行为的替代可能是非常有问题的，因为它可能导致变量的隐藏，就像我们在context.withValue示例中看到的那样。
这种方法的另一个大缺点是隐藏正在发生的事情并使代码更加模糊。当涉及到可选值时，一个更好的方法是使用指向结构参数的指针——这允许您避免使用nil传递结构。
假设您有以下代码：


// This function has two mandatory args and 4 optional ones
func SomeFunc(ctx context.Context, arg1, arg2 int, 
    opt1, opt2, opt3, opt4 string) {}
```

By using Optional, you would have something like this:
通过使用可选选项，您可以获得如下内容：

```
type Optional struct {
    Opt1 string
    Opt2 string
    Opt3 string
    Opt4 string
}

// This function has two mandatory args and 4 optional ones
func SomeFunc(ctx context.Context, arg1, arg2 int, o *Optional) {}
```

Globals

Some global variables can be stored in a context so that they can be passed through a series of function calls. This is generally not good practice since globals are available in every point of the application, so using context to store and recall them is pointless and a waste of resources and performance. If your package has some globals, you can use the Singleton pattern we looked at in Chapter 12, Synchronization with sync and atomic, to allow access to them from any point of your package or application.

```
全局
一些全局变量可以存储在上下文中，以便通过一系列函数调用传递它们。这通常不是好的实践，因为全局变量在应用程序的每个点都可用，所以使用上下文来存储和调用它们是毫无意义的，浪费资源和性能。如果您的包有一些全局变量，您可以使用我们在第12章同步中看到的单例模式。使用Sync和Atomic进行电子化，以允许从包或应用程序的任何点访问它们。

```

# Main interface and usage
The signature of the package will include a context, the root folder, the search term, and a couple of optional parameters:

Search in contents: Will look for the string in the file's contents instead of the name
Exclude list: Will not search the files with the selected name/names
The function would look something like this:

```
主界面及用途
包的签名将包括上下文、根文件夹、搜索词和几个可选参数：
搜索内容：将查找文件内容中的字符串，而不是名称
排除列表：将不搜索具有所选名称的文件
该函数将如下所示：

type Options struct {
    Contents bool
    Exclude []string
}

func FileSearch(ctx context.Context, root, term string, o *Options)
```

Since it should be a concurrent function, the return type could be a channel of result, which could be either an error or a series of matches in a file. Since we can search for the names of content, the latter could have more than one match:

因为它应该是一个并发函数，所以返回类型可以是一个结果通道，它可以是一个错误，也可以是一个文件中的一系列匹配项。由于我们可以搜索内容的名称，后者可以有多个匹配项：

```
type Result struct {
    Err error
    File string
    Matches []Match
}

type Match struct {
    Line int
    Text string
}
```

The previous function will return a receive-only channel of the Result type:
上一个函数将返回结果类型的仅接收通道：

```
func FileSearch(ctx context.Context, root, term string, o *Options) <-chan Result
```

Here, this function would keep receiving values from the channel until it gets closed:
在这里，此函数将一直从通道接收值，直到关闭：

```
for r := range FileSearch(ctx, directory, searchTerm, options) {
    if r.Err != nil {
        fmt.Printf("%s - error: %s\n", r.File, r.Err)
        continue
    }
    if !options.Contents {
        fmt.Printf("%s - match\n", r.File)
        continue
    }
    fmt.Printf("%s - matches:\n", r.File)
    for _, m := range r.Matches {
        fmt.Printf("\t%d:%s\n", m.Line, m.Text)
    }
}
```


# Exit and entry points

The result channel should be closed by either the cancellation of the context, or by the search being over. Since a channel cannot be closed twice, we can use sync.Once to avoid closing the channel for the second time. To keep track of the goroutines that are running, we can use sync.Waitgroup:

```
出入口点
结果通道应该通过取消上下文或结束搜索来关闭。因为一个频道不能关闭两次，所以我们可以使用sync.one来避免第二次关闭频道。要跟踪正在运行的goroutine，我们可以使用sync.waitgroup:

ch, wg, once := make(chan Result), sync.WaitGroup{}, sync.Once{}
go func() {
    wg.Wait()
    fmt.Println("* Search done *")
    once.Do(func() {
        close(ch)
    })
}()
go func() {
    <-ctx.Done()
    fmt.Println("* Context done *")
    once.Do(func() {
        close(ch)
    })
}()
```

We could launch a goroutine for each file so that we can define a private function that we can use as an entry point and then use it recursively for subdirectories:

我们可以为每个文件启动一个goroutine，这样我们就可以定义一个可以用作入口点的私有函数，然后将其递归地用于子目录：

```
func fileSearch(ctx context.Context, ch chan<- Result, wg *sync.WaitGroup, file, term string, o *Options)
```

The main exported function will start by adding a value to the wait group. It will then launch the private function, starting it as an asynchronous process:

```
主导出函数将通过向等待组添加值开始。然后它将启动私有函数，将其作为异步进程启动：

wg.Add(1)
go fileSearch(ctx, ch, &wg, root, term, o)
```

The last thing each fileSearch should do is call WaitGroup.Done to mark the end of the current file.
每个文件搜索最不应该做的就是调用waitgroup.done来标记当前文件的结尾。

# Exclude list

The private function will decrease the wait group counter before it finishes using the Done method.. Besides that, the first thing it should do is check the filename so that it can skip it if it is in the exclusion list:

```
排除列表
在使用done方法完成之前，private函数将减少wait group计数器。除此之外，它应该做的第一件事是检查文件名，以便在排除列表中跳过它：

defer wg.Done()
_, name := filepath.Split(file)
if o != nil {
    for _, e := range o.Exclude {
        if e == name {
            return
        }
    }
}
```

If that is not the case, we can check the current file's information using os.Stat and send an error to the channel if we don't succeed. Since we cannot risk causing a panic by sending into a closed channel, we can check whether the context is done, and if not, send the error:

```
如果不是这样，我们可以使用os.stat检查当前文件的信息，如果不成功，则向通道发送错误。由于我们不能通过发送到封闭通道来冒险引起恐慌，因此我们可以检查上下文是否已完成，如果没有，则发送错误：

info, err := os.Stat(file)
if err != nil {
    select {
    case <-ctx.Done():
        return
    default:
        ch <- Result{File: file, Err: err}
    }
    return
}
```

# Handling directories
The information that's received will tell us whether the file is a directory or not. If it is a directory, we can get a list of files and handle the error, as we did earlier with os.Stat. Then, we can launch another series of searches, one for each file, if the context isn't already done. The following code sums up these operations:

```
处理目录
收到的信息将告诉我们文件是否是目录。如果它是一个目录，我们可以得到一个文件列表并处理错误，就像前面对os.stat所做的那样。然后，如果上下文尚未完成，我们可以启动另一系列搜索，每个文件一个。以下代码总结了这些操作：


if info.IsDir() {
    files, err := ioutil.ReadDir(file)
    if err != nil {
        select {
        case <-ctx.Done():
            return
        default:
            ch <- Result{File: file, Err: err}
        }
        return
    }
    select {
    case <-ctx.Done():
    default:
        wg.Add(len(files))
        for _, f := range files {
            go fileSearch(ctx, ch, wg, filepath.Join(file, 
        f.Name()), term, o)
        }
    }
    return
}

```

# Checking file names and contents
If the file is a regular file and not a directory, we can compare the file name or its contents, depending on the options that are specified. Checking the file name is pretty easy:

```
检查文件名和内容
如果文件是常规文件而不是目录，我们可以根据指定的选项比较文件名或其内容。检查文件名非常简单：

if o == nil || !o.Contents {
    if name == term {
        select {
        case <-ctx.Done():
        default:
            ch <- Result{File: file}
        }
    }
    return
}
```

If we are searching for the contents, we should open the file:
如果要搜索内容，应打开文件：

```
f, err := os.Open(file)
if err != nil {
    select {
    case <-ctx.Done():
    default:
        ch <- Result{File: file, Err: err}
    }
    return
}
defer f.Close()


```

Then, we can read the file line by line to search for the selected term. If the context expires while we are reading the file, we will stop all operations:

然后，我们可以一行一行地读取文件来搜索选定的术语。如果在读取文件时上下文过期，我们将停止所有操作：

```
scanner, matches, line := bufio.NewScanner(f), []Match{}, 1
for scanner.Scan() {
    select {
    case <-ctx.Done():
        break
    default:
        if text := scanner.Text(); strings.Contains(text, term) {
            matches = append(matches, Match{Line: line, Text: text})
        }
        line++
    }
}

```

Finally, we can check for errors from the scanner. If there's none and the search has results, we can send all the matches to the output channel:

```
最后，我们可以检查扫描仪的错误。如果没有，并且搜索有结果，我们可以将所有匹配项发送到输出通道：

select {
case <-ctx.Done():
    break
default:
    if err := scanner.Err(); err != nil {
        ch <- Result{File: file, Err: err}
        return
    }
    if len(matches) != 0 {
        ch <- Result{File: file, Matches: matches}
    }
}
```

In less than 200 lines, we created a concurrent file search function that uses one goroutine per file. It takes advantage of a channel to send results and synchronization primitives in order to coordinate operations.

在不到200行中，我们创建了一个并发文件搜索函数，每个文件使用一个goroutine。它利用通道发送结果和同步原语以协调操作。



# Summary
In this chapter, we looked at what one of the newer packages, context, is all about. We saw that Context is a simple interface that has four methods, and should be used as the first argument of a function. Its main scope is to handle cancellation and deadlines to synchronize concurrent operations and provide the user with functionality to cancel an operation.

We saw how the default contexts, Background and TODO, don't allow cancellation, but they can be extended using various functions of the package to add timeouts or cancellation. We also talked about the capabilities of context when it comes to holding values and how this should be used carefully in order to avoid shadowing and other problems.

Then, we dived into the standard package to see where context is already used. This included the HTTP capabilities of requests, where it can be used for values, cancellation, and timeout, and the server shutdown operation. We also saw how the TCP package allows us to use it in a similar fashion with a practical example, and we also listed the operations in the database package that allow us to use context to cancel them.

Before building our own functionality using context, we went into some of the uses that should be avoided, from using the wrong types of keys to using context to pass values around that should be in a function or method signature instead. Then, we proceeded to create a function that searches files and contents, using what we have learned about concurrency from the last three chapters.

The next chapter will conclude the concurrency section of this book by showing off the most common Go concurrency patterns and their usage. This will allow us to put together all that we have learned so far about concurrency in some very common and effective configurations.

```
总结
在本章中，我们研究了一个较新的包context的全部内容。我们看到Context是一个简单的接口，它有四个方法，应该用作函数的第一个参数。它的主要作用是处理取消操作和同步并发操作的截止日期，并为用户提供取消操作的功能。
我们看到了默认上下文background和todo如何不允许取消，但是可以使用包的各种功能来扩展它们，以添加超时或取消。我们还讨论了上下文在保存值时的功能，以及如何小心使用它以避免隐藏和其他问题。
然后，我们深入到标准包中，看看上下文在哪里被使用。这包括请求的HTTP功能，可用于值、取消和超时，以及服务器关闭操作。我们还看到了TCP包如何允许我们以类似的方式使用它和一个实际的例子，我们还列出了数据库包中允许我们使用上下文来取消它们的操作。
在使用上下文构建我们自己的功能之前，我们讨论了一些应该避免的用法，从使用错误类型的键到使用上下文传递应该在函数或方法签名中的值。然后，我们继续创建一个搜索文件和内容的函数，使用我们在最后三章中所学到的关于并发性的知识。
下一章将通过展示最常见的go并发模式及其用法来结束本书的并发部分。这将允许我们在一些非常常见和有效的配置中汇总迄今为止所学的关于并发性的所有知识。
```


Questions
What is a context in Go?
What's the difference between cancellation, deadline, and timeout?
What are the best practices when passing values with a context?
Which standard packages already use context?

```go

问题
Go中的上下文是什么？
取消、截止日期和超时之间有什么区别？
通过上下文传递值时，最佳实践是什么？
哪些标准包已经使用上下文？
```

# Implementing Concurrency Patterns

This chapter will be about concurrency patterns and how to use them to build robust system applications. We have already looked at all the tools that are involved in concurrency (goroutines and channels, sync and atomic, and context) so now we will look at some common ways of combining them in patterns so that we can use them in our programs.

The following topics will be covered in this chapter:

Beginning with generators
Sequencing with pipelines
Muxing and demuxing
Other patterns
Resource leaking

```
实现并发模式
本章将讨论并发模式以及如何使用它们构建健壮的系统应用程序。我们已经研究了所有涉及并发的工具（goroutines和channels、sync和atomic以及context），所以现在我们将研究一些将它们组合在模式中的常见方法，以便在程序中使用它们。
本章将讨论以下主题：
从生成器开始
管道排序
Muxing和Demuxing
其他模式
资源泄漏
```

# Beginning with generators
A generator is a function that returns the next value of a sequence each time it is called. The biggest advantage of using generators is the lazy creation of new values of the sequence. In Go, this can either be expressed with an interface or with a channel. One of the upsides of generators when they're used with channels is that they produce values concurrently, in a separate goroutine, leaving the main goroutine capable of executing other kinds of operations on them.

It can be abstracted with a very simple interface:

```
从生成器开始
生成器是一个函数，每次调用它时都返回序列的下一个值。使用生成器的最大优势是懒惰地创建序列的新值。在Go中，这可以通过接口或通道来表示。生成器与通道一起使用的好处之一是，它们在单独的goroutine中同时生成值，从而使主goroutine能够对其执行其他类型的操作。
它可以用一个非常简单的接口抽象出来：

type Generator interface {
    Next() interface{}
}

type GenInt64 interface {
    Next() int64
}
```

The return type of the interface will depend on the use case, which is int64 in our case. The basic implementation of it could be something such as a simple counter:

```
接口的返回类型将取决于用例，在我们的例子中是int64。它的基本实现可以是一个简单的计数器：

type genInt64 int64

func (g *genInt64) Next() int64 {
    *g++
    return int64(*g)
}
```

This implementation is not thread-safe, so if we try to use it with goroutines, we could lose some elements on the way:

```
此实现不是线程安全的，因此如果我们尝试将其与goroutine一起使用，可能会在使用过程中丢失一些元素：

func main() {
    var g genInt64
    for i := 0; i < 1000; i++ {
        go func(i int) {
            fmt.Println(i, g.Next())
        }(i)
    }
    time.Sleep(time.Second)
}
```

A simple way to make a generator concurrent is to execute atomic operations on the integer.

This will make the concurrent generator thread-safe, with very few changes to the code needing to happen:

```
使生成器并发的一个简单方法是对整数执行原子操作。
这将使并发生成器线程安全，只需对代码进行很少的更改：


type genInt64 int64

func (g *genInt64) Next() int64 {
    return atomic.AddInt64((*int64)(g), 1)
}


```

This will avoid race conditions in the application. However, there is another implementation that's possible, but this requires the use of channels. The idea is to produce the value in a goroutine and then pass it in a shared channel to the next method, as shown in the following example:
这将避免应用程序中的竞争条件。然而，还有另一种可能的实现，但这需要使用通道。其思想是在goroutine中生成值，然后在共享通道中将其传递给下一个方法，如下面的示例所示：

```
type genInt64 struct {
    ch chan int64
}

func (g genInt64) Next() int64 {
    return <-g.ch
}

func NewGenInt64() genInt64 {
    g := genInt64{ch: make(chan int64)}
    go func() {
        for i := int64(0); ; i++ {
            g.ch <- i
        }
    }()
    return g
}


```

The loop will go on forever, and will be blocking in the send operation when the generator user stops requesting new values with the Next method.

The code was structured this way because we were trying to implement the interface we defined at the beginning. We could also just return a channel and use it for receiving:

``` 
循环将永远持续下去，并且在发送操作中，当生成器用户停止使用下一个方法请求新值时，循环将被阻塞。
代码的结构是这样的，因为我们试图实现我们在开始时定义的接口。我们还可以返回一个频道并使用它接收：

func GenInt64() <-chan int64 {
 ch:= make(chan int64)
    go func() {
        for i := int64(0); ; i++ {
            ch <- i
        }
    }()
    return ch
}

```

The main advantage of using the channel directly is the possibility of including it in select statements in order to choose between different channel operations. The following shows a select between two different generators:

```
直接使用通道的主要优点是可以将其包含在select语句中，以便在不同的通道操作之间进行选择。下面显示两个不同的生成器之间的选择：

func main() {
    ch1, ch2 := GenInt64(), GenInt64()
    for i := 0; i < 20; i++ {
        select {
        case v := <-ch1:
            fmt.Println("ch 1", v)
        case v := <-ch2:
            fmt.Println("ch 2", v)
        }
    }
}

```


# Avoiding leaks 避免泄漏

It is a good idea to allow the loop to end in order to avoid goroutine and resource leakage. Some of these issues are as follows:


允许循环结束是一个好主意，以避免Goroutine和资源泄漏。其中一些问题如下：

When a goroutine hangs without returning, the space in memory remains used, contributing to the application's size in memory. The goroutine and the variables it defines in the stack will get collected by the GC only when the goroutine returns or panics.
If a file remains open, this can prevent other processes from executing operations on it. If the number of files that are open reaches the limit imposed by the OS, the process will not be able to open other files (or accept network connections).
An easy solution to this problem is to always use context.Context so that you have a well-defined exit point for the goroutine:

```
当Goroutine挂起而不返回时，内存中的空间将继续使用，这将导致应用程序在内存中的大小。只有当goroutine返回或崩溃时，GC才会收集goroutine及其在堆栈中定义的变量。
如果文件保持打开状态，这可能会阻止其他进程对其执行操作。如果打开的文件数达到操作系统所规定的限制，进程将无法打开其他文件（或接受网络连接）。
此问题的一个简单解决方案是始终使用context.context，以便为goroutine定义一个明确的出口点：

func NewGenInt64(ctx context.Context) genInt64 {
    g := genInt64{ch: make(chan int64)}
    go func() {
        for i := int64(0); ; i++ {
            select {
            case g.ch <- i:
                // do nothing
            case <-ctx.Done():
                close(g.ch)
                return
            }
        }
    }()
    return g
}
```

This can be used to generate values until there is a need for them and cancel the context when there's no need for new values. The same pattern can be applied to a version that returns a channel. For instance, we could use the cancel function directly or set a timeout on the context:

``` 
这可用于生成值，直到需要这些值为止，并在不需要新值时取消上下文。同样的模式可以应用于返回通道的版本。例如，我们可以直接使用cancel函数或在上下文上设置超时：

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
    defer cancel()
    g := NewGenInt64(ctx)
    for i := range g.ch {
        go func(i int64) {
            fmt.Println(i, g.Next())
        }(i)
    }
    time.Sleep(time.Second)
}
```

The generator will produce numbers until the context that's provided expires. At this point, the generator will close the channel.
生成器将生成数字，直到提供的上下文过期。此时，发电机将关闭通道。


# Sequencing with pipelines
A pipeline is a way of structuring the application flow, and is obtained by splitting the main execution into stages that can talk with one another using certain means of communication. This could be either of the following:

External, such as a network connection or a file
Internal to the application, like Go's channels
The first stage is often referred to as the producer, while the last one is often called the consumer.

The set of concurrency tools that Go offers allows us to efficiently use multiple CPUs and optimize their usage by blocking input or output operations. Channels in particular are the perfect tools for internal pipeline communication. They can be represented by functions that receive an inbound channel and return an outbound one. The base structure would look something like this:

``` 
管道排序
管道是构建应用程序流的一种方式，它是通过将主执行划分为多个阶段来获得的，这些阶段可以使用某种通信方式彼此进行通信。这可以是以下任一项：
外部，如网络连接或文件
应用程序内部，如Go的渠道
第一个阶段通常被称为生产者，而最后一个阶段通常被称为消费者。
Go提供的一组并发工具允许我们高效地使用多个CPU，并通过阻塞输入或输出操作来优化它们的使用。尤其是通道是内部管道通信的理想工具。它们可以由接收入站通道并返回出站通道的函数表示。基本结构如下所示：



func stage(in <-chan interface{}) <-chan interface{} {
    var out = make(chan interface{})
    go func() {
        for v := range in {
            v = v.(int)+1 // some operation
            out <- v
        }
        close(out)
    }()
    return out
}

```

We create a channel of the same type as the input channel and return it. In a separate goroutine, we receive data from the input channel, perform an operation on the data, and we send it to the output channel.

This pattern can be further improved by the use of context.Context so that we have greater control of the application flow. It would look something like the following code:

```
我们创建一个与输入通道类型相同的通道并返回它。在单独的goroutine中，我们从输入通道接收数据，对数据执行操作，然后将其发送到输出通道。
通过使用context.context可以进一步改进此模式，以便更好地控制应用程序流。它看起来像下面的代码：

func stage(ctx context.Context, in <-chan interface{}) <-chan interface{} {
    var out = make(chan interface{})
    go func() {
        defer close(out)
        for v := range in {
            v = v.(int)+1 // some operation
            select {
                case out <- v:
                case <-ctx.Done():
                    return
            }
        }
    }()
    return out
}
```

There are a couple of general rules that we should follow when designing pipelines:

Intermediate stages will receive an inbound channel and return another one.
The producer will not receive any channel, but return one.
The consumer will receive a channel without returning one.
Each stage will close the channel on creation when it's done sending messages.
Each stage should keep receiving from the input channel until it's closed.
Let's create a simple pipeline that filters lines from a reader using a certain string and prints the filtered lines, highlighting the search string. We can start with the first stage – the source – which will not receive any channel in its signature but will use a reader to scan lines. We take a context in for reacting to an early exit request (context cancellation) and a bufio scanner to read line by line. The following code shows this:

```
在设计管道时，我们应该遵循以下几个一般规则：
中间阶段将接收一个入站通道并返回另一个通道。
制作人不会收到任何频道，但会返回一个。
消费者将收到一个没有返回的频道。
当完成发送消息后，每个阶段都将在创建时关闭通道。
每个阶段都应该保持从输入通道接收，直到它关闭。
让我们创建一个简单的管道，它使用特定的字符串过滤读卡器中的行，并打印过滤后的行，突出显示搜索字符串。我们可以从第一个阶段开始——源——它不会接收到它的签名中的任何通道，而是使用读卡器扫描行。我们采用一个上下文来响应早期退出请求（上下文取消），使用bufio扫描器逐行读取。下面的代码显示了这一点：


func SourceLine(ctx context.Context, r io.ReadCloser) <-chan string {
    ch := make(chan string)
    go func() {
        defer func() { r.Close(); close(ch) }()
        s := bufio.NewScanner(r)
        for s.Scan() {
            select {
            case <-ctx.Done():
                return
            case ch <- s.Text():
            }
        }
    }()
    return ch
}
```

We can split the remaining operations into two phases: a filtering phase and a writing phase. The filtering phase will simply filter from a source channel to an output channel. We are still passing context in order to avoid sending extra data if the context is already complete. This is the text filter implementation:

```
我们可以将剩余的操作分为两个阶段：过滤阶段和写入阶段。滤波相位将简单地从源信道过滤到输出信道。我们仍然在传递上下文，以避免在上下文已经完成时发送额外的数据。这是文本筛选实现：

func TextFilter(ctx context.Context, src <-chan string, filter string) <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        for v := range src {
            if !strings.Contains(v, filter) {
                continue
            }
            select {
            case <-ctx.Done():
                return
            case ch <- v:
            }
        }
    }()
    return ch
}
```

Finally, we have the final stage, the consumer, which will print the output in a writer and will also use a context for early exit:

```
最后，我们还有最后一个阶段，消费者，它将在一个编写器中打印输出，并且还将使用上下文来提前退出：

func Printer(ctx context.Context, src <-chan string, color int, highlight string, w io.Writer) {
    const close = "\x1b[39m"
    open := fmt.Sprintf("\x1b[%dm", color)
    for {
        select {
        case <-ctx.Done():
            return
        case v, ok := <-src:
            if !ok {
                return
            }
            i := strings.Index(v, highlight)
            if i == -1 {
                panic(v)
            }
            fmt.Fprint(w, v[:i], open, highlight, close, v[i+len(highlight):], "\n")
        }
    }
}
```

The use of this function is as follows:

```
此功能的使用如下：

func main() {
    var search string
    ...
    ctx := context.Background()
    src := SourceLine(ctx, ioutil.NopCloser(strings.NewReader(sometext)))
    filter := TextFilter(ctx, src, search)
    Printer(ctx, filter, 31, search, os.Stdout)
}
```
With this approach, we learned how to split a complex operation into a simple task that's executed by stages and connected using channels.

```
通过这种方法，我们学习了如何将一个复杂的操作拆分为一个简单的任务，该任务由阶段执行，并使用通道进行连接。
```

Muxing and demuxing
Now that we are familiar with pipelines and stages, we can introduce two new concepts:

Muxing (multiplexing) or fan-out: Receiving from one channel and sending to multiple channels
Demuxing (demultiplexing) or fan-in: Receiving from multiple channels and sending through one channel
This pattern is very common and allows us to use the power of concurrency in different ways. The most obvious way is to distribute data from a channel that is quicker than its following step, and create more than one instance of such steps to make up for the difference in speed.

```
Muxing和Demuxing
现在我们已经熟悉了管道和阶段，我们可以引入两个新概念：
多路复用或扇出：从一个信道接收并发送到多个信道
解复用或扇入：从多个信道接收并通过一个信道发送
这种模式非常常见，允许我们以不同的方式使用并发的能力。最明显的方法是从一个比下面的步骤更快的通道中分发数据，并创建这些步骤的多个实例来弥补速度上的差异。

```

Fan-out
The implementation of multiplexing is pretty straightforward. The same channel needs to be passed to different stages so that each one will be reading from it.

Each goroutine is competing for resources during the runtime schedule, so if we want to reserve more of them, we can use more than one goroutine for a certain stage of the pipeline, or for a certain operation in our application.

We can create a small application that counts the occurrence of words which appear in a piece of text using such an approach. Let's create an initial producer stage that reads from a writer and returns a slice of words for that line:

```
扇出
多路复用的实现非常简单。同一个通道需要传递到不同的阶段，以便每个阶段都能从中读取信息。
每个goroutine在运行时计划期间都在竞争资源，因此如果我们想保留更多的资源，可以在管道的某个阶段或应用程序中的某个操作中使用多个goroutine。
我们可以创建一个小应用程序，用这种方法计算出现在一段文本中的单词的出现次数。让我们创建一个初始的生产者阶段，它从一个作者处读取并返回该行的一部分单词：


func SourceLineWords(ctx context.Context, r io.ReadCloser) <-chan []string {
    ch := make(chan []string)
    go func() {
        defer func() { r.Close(); close(ch) }()
        b := bytes.Buffer{}
        s := bufio.NewScanner(r)
        for s.Scan() {
            b.Reset()
            b.Write(s.Bytes())
            words := []string{}
            w := bufio.NewScanner(&b)
            w.Split(bufio.ScanWords)
            for w.Scan() {
                words = append(words, w.Text())
            }
            select {
            case <-ctx.Done():
                return
            case ch <- words:
            }
        }
    }()
    return ch
}
```

In order to use the first stage as a source for more than one instance of the second stage, we just need to create more than one counting stage with the same input channel:

```
为了将第一阶段用作第二阶段的多个实例的源，我们只需要创建具有相同输入通道的多个计数阶段：

ctx, canc := context.WithCancel(context.Background())
defer canc()
src := SourceLineWords(ctx,   
    ioutil.NopCloser(strings.NewReader(cantoUno)))
count1, count2 := WordOccurrence(ctx, src), WordOccurrence(ctx, src)
```

# Fan-in
Demuxing is a little bit more complicated because we don't need to receive data blindly in one goroutine or another – we actually need to synchronize a series of channels. A good approach to avoid race conditions is to create another channel where all the data from the various input channels will be received. We also need to make sure that this merge channel gets closed once all the channels are done. We also have to keep in mind that the channel will be closed if the context is cancelled. We are using sync.Waitgroup here to wait for all the channels to finish:

解组有点复杂，因为我们不需要在一个或另一个goroutine中盲目地接收数据——我们实际上需要同步一系列通道。避免竞争条件的一个好方法是创建另一个通道，在该通道中将接收来自不同输入通道的所有数据。我们还需要确保在完成所有通道后关闭此合并通道。我们还必须记住，如果上下文被取消，通道将被关闭。我们在这里使用sync.waitgroup等待所有频道完成：

```
wg := sync.WaitGroup{}
merge := make(chan map[string]int)
wg.Add(len(src))
go func() {
    wg.Wait()
    close(merge)
}()
```

The problem is that we have two possible triggers for closing the channel: regular transmission ending and context cancelling.

We have to make sure that if the context ends, no message is sent to the outbound channel. Here, we are collecting the values from the input channels and sending them to the merge channel, but only if the context isn't complete. We do this in order to avoid a send operation being sent to a closed channel, which would make our application panic:

问题是我们有两个关闭通道的可能触发器：常规传输结束和上下文取消。
我们必须确保如果上下文结束，就不会向出站通道发送消息。这里，我们从输入通道收集值并将其发送到合并通道，但前提是上下文不完整。我们这样做是为了避免发送操作被发送到一个关闭的通道，这将使我们的应用程序恐慌：

```
for _, ch := range src {
    go func(ch <-chan map[string]int) {
        defer wg.Done()
        for v := range ch {
            select {
            case <-ctx.Done():    
                return
            case merge <- v:
            }
        }
    }(ch)
}
```

Finally, we can focus on the last operation, which uses the merge channel to execute our final word count:
最后，我们可以将重点放在最后一个操作上，它使用合并通道来执行最终的字数计数：

```
count := make(map[string]int)
for {
    select {
    case <-ctx.Done():
        return count
    case c, ok := <-merge:
        if !ok {
            return count
        }
        for k, v := range c {
            count[k] += v
        }
    }
}

```

The application's main function, with the addition of the fan-in, will look as follows:
应用程序的主要功能，加上扇入，如下所示：

```
func main() {
    ctx, canc := context.WithCancel(context.Background())
    defer canc()
    src := SourceLineWords(ctx, ioutil.NopCloser(strings.NewReader(cantoUno)))
    count1, count2 := WordOccurrence(ctx, src), WordOccurrence(ctx, src)
    final := MergeCounts(ctx, count1, count2)
    fmt.Println(final)
}
```

We can see that the fan-in is the most complex and critical part of the application. Let's recap the decisions that helped build a fan-in function that is free from panic or deadlock:

Use a merge channel to collect values from the various input.
Have sync.WaitGroup with a counter equal to the number of input channels.
Use it in a separate goroutine and wait for it to close the channel.
For each input channel, create a goroutine that transfers the values to the merge channel.
Ensure that you send the record only if the context is not complete.
Use the wait group's done function before exiting such a goroutine.
Following the preceding steps will allow us to use the merge channel with a simple range. In our example, we are also checking whether the context is complete before receiving from the channel in order to allow for an early exit from the goroutine.

```
我们可以看到，扇入是应用程序中最复杂和最关键的部分。让我们回顾一下有助于构建无恐慌或死锁的扇入函数的决策：
使用合并通道从各种输入中收集值。
使sync.waitgroup的计数器等于输入通道的数目。
在单独的goroutine中使用它，并等待它关闭通道。
对于每个输入通道，创建一个将值传输到合并通道的goroutine。
确保仅在上下文不完整时发送记录。
在退出此类goroutine之前，请使用wait group的done函数。
按照前面的步骤，我们可以使用具有简单范围的合并通道。在我们的示例中，我们还将在从通道接收之前检查上下文是否完整，以允许提前退出Goroutine。

```

# Producers and consumers

Channels allow us to easily handle a scenario in which multiple consumers receive data from one producer and vice versa.

The case with a single producer and one consumer, as we have already seen, is pretty straightforward:

```
渠道使我们能够轻松处理多个消费者从一个生产者接收数据的场景，反之亦然。
正如我们已经看到的，单一生产商和一个消费者的情况非常简单：

func main() {
    // one producer
    var ch = make(chan int)
    go func() {
        for i := 0; i < 100; i++ {
            ch <- i
        }
        close(ch)
    }()
    // one consumer
    var done = make(chan struct{})
    go func() {
        for i := range ch {
            fmt.Println(i)
        }
        close(done)
    }()
    <-done
}
```

# Multiple producers (N * 1)

Having multiple producers or consumers can be easily handled using wait groups. In the case of multiple producers, all the goroutines will share the same channel:

```
拥有多个生产者或消费者可以使用等待组轻松处理。在多个生产商的情况下，所有Goroutine将共享同一渠道：


// three producer
var ch = make(chan string)
wg := sync.WaitGroup{}
wg.Add(3)
for i := 0; i < 3; i++ {
    go func(n int) {
        for i := 0; i < 100; i++ {
            ch <- fmt.Sprintln(n, i)
        }
        wg.Done()
    }(i)
}
go func() {
    wg.Wait()
    close(ch)
}()
```

They will use sync.WaitGroup to wait for each producer to finish before closing the channel.
他们将使用sync.waitgroup在关闭频道之前等待每个制作人完成。

# Multiple consumers (1 * M)
The same reasoning applies with multiple consumers – they all receive from the same channel in different goroutines:

同样的道理也适用于多个消费者——他们都从同一渠道以不同的方式接收：

```
func main() {
    // three consumers
    wg := sync.WaitGroup{}
    wg.Add(3)
    var ch = make(chan string)

    for i := 0; i < 3; i++ {
        go func(n int) {
            for i := range ch {
                fmt.Println(n, i)
            }
            wg.Done()
        }(i)
    }

    // one producer
    go func() {
        for i := 0; i < 10; i++ {
            ch <- fmt.Sprintln("prod-", i)
        }
        close(ch)
    }()

    wg.Wait()
}

```

In this case, sync.WaitGroup is used to wait for the application to end.
在这种情况下，sync.waitgroup用于等待应用程序结束。

# Multiple consumers and producers (N*M)
The last scenario is where we have an arbitrary number of producers (N) and another arbitrary number of consumers (M).

In this case, we need two waiting groups: one for the producer and another for the consumer:

最后一个场景是，我们有任意数量的生产者（n）和另一个任意数量的消费者（m）。
在这种情况下，我们需要两个等待组：一个用于生产者，另一个用于消费者：

```
const (
    N = 3
    M = 5
)
wg1 := sync.WaitGroup{}
wg1.Add(N)
wg2 := sync.WaitGroup{}
wg2.Add(M)
var ch = make(chan string)

```

This will be followed by a series of producers and consumers, each one in their own goroutine:
接下来是一系列的生产者和消费者，每个人都有自己的产品：

```
for i := 0; i < N; i++ {
    go func(n int) {
        for i := 0; i < 10; i++ {
            ch <- fmt.Sprintf("src-%d[%d]", n, i)
        }
        wg1.Done()
    }(i)
}

for i := 0; i < M; i++ {
    go func(n int) {
        for i := range ch {
            fmt.Printf("cons-%d, msg %q\n", n, i)
        }
        wg2.Done()
    }(i)
}
```

The final step is to wait for the WaitGroup producer to finish its work in order to close the channel.
Then, we can wait for the consumer channel to let all the messages be processed by the consumers:

``` 
最后一步是等待waitgroup生产者完成其工作以关闭频道。
然后，我们可以等待消费者通道让消费者处理所有消息：
```

```
wg1.Wait()
close(ch)
wg2.Wait()
```

# Other patterns

So far, we've looked at the most common concurrency patterns that can be used. Now, we will focus on some that are less common but are worth mentioning.
其他模式
到目前为止，我们已经研究了可以使用的最常见的并发模式。现在，我们将关注一些不太常见但值得一提的问题。

# Error groups
The power of sync.WaitGroup is that it allows us to wait for simultaneous goroutines to finish their jobs. We have already looked at how sharing context can allow us to give the goroutines an early exit if it's used correctly. The first concurrent operation, such as send or receive from a channel, is in the select block, together with the context completion channel:

```
sync.waitgroup的强大之处在于，它允许我们等待同时进行的goroutine完成它们的工作。我们已经研究了共享上下文如何允许我们在正确使用时提前退出Goroutines。第一个并发操作（如从通道发送或接收）与上下文完成通道一起位于选择块中：

func main() {
    ctx, canc := context.WithTimeout(context.Background(), time.Second)
    defer canc()
    wg := sync.WaitGroup{}
    wg.Add(10)
    var ch = make(chan int)
    for i := 0; i < 10; i++ {
        go func(ctx context.Context, i int) {
            defer wg.Done()
            d := time.Duration(rand.Intn(2000)) * time.Millisecond
            time.Sleep(d)
            select {
            case <-ctx.Done():
                fmt.Println(i, "early exit after", d)
                return
            case ch <- i:
                fmt.Println(i, "normal exit after", d)
            }
        }(ctx, i)
    }
    go func() {
        wg.Wait()
        close(ch)
    }()
    for range ch {
    }
}
```

An improvement on this scenario is offered by the experimental golang.org/x/sync/errgroup package.

The built-in goroutines are always of the func() type, but this package allows us to execute func() error concurrently and return the first error that's received from the various goroutines.

This is very useful in scenarios where you launch more goroutines together and receive the first error. The errgroup.Group type can be used as a zero value, and its Do method takes func() error as an argument and launches the function concurrently.

The Wait method either waits for all the functions to finish successfully and returns nil, or it returns the first error that comes from any of the functions.

Let's create an example that defines a URL visitor, that is, a function that gets a URL string and returns func() error, which makes the call:

```
实验golang.org/x/sync/errgroup包提供了对这种情况的改进。
内置goroutines始终是func（）类型，但此包允许我们同时执行func（）错误，并返回从各种goroutines接收到的第一个错误。
这在一起启动更多goroutine并接收第一个错误的场景中非常有用。errGroup.Group类型可以用作零值，其do方法将func（）错误作为参数，并同时启动函数。
wait方法要么等待所有函数成功完成并返回nil，要么返回来自任何函数的第一个错误。
让我们创建一个定义URL访问者的示例，即获取URL字符串并返回func（）错误的函数，该函数使调用：


func visitor(url string) func() error {
    return func() (err error) {
        s := time.Now()
        defer func() {
            log.Println(url, time.Since(s), err)
        }()
        var resp *http.Response
        if resp, err = http.Get(url); err != nil {
            return
        }
        return resp.Body.Close()
    }
}
```

We can use it directly with the Go method and wait. This will return the error that was caused by the invalid URL:
我们可以直接使用Go方法并等待。这将返回由无效URL引起的错误：

```
func main() {
    eg := errgroup.Group{}
    var urlList = []string{
        "http://www.golang.org/",
        "http://invalidwebsite.hey/",
        "http://www.google.com/",
    }
    for _, url := range urlList {
        eg.Go(visitor(url))
    }
    if err := eg.Wait(); err != nil {
        log.Fatalln("Error:", err)
    }
}
```

The error group also allows us to create a group, along with a context, with the WithContext function. This context gets cancelled when the first error is received. The context's cancellation enables the Wait method to return right away, but it also allows an early exit in the goroutines in your functions.

We can create a similar func() error creator that will send values into a channel until the context is closed. We will introduce a small chance (1%) of raising an error:

错误组还允许我们使用withContext函数创建一个组以及一个上下文。当收到第一个错误时，此上下文将被取消。上下文的取消允许wait方法立即返回，但它也允许在函数的goroutines中提前退出。
我们可以创建一个类似的func（）错误创建者，将值发送到一个通道中，直到上下文关闭。我们将引入引发错误的小概率（1%）：

```
func sender(ctx context.Context, ch chan<- string, n int) func() error {
    return func() (err error) {
        for i := 0; ; i++ {
            if rand.Intn(100) == 42 {
                return errors.New("the answer")
            }
            select {
            case ch <- fmt.Sprintf("[%d]%d", n, i):
            case <-ctx.Done():
                return nil
            }
        }
    }
}


```

We will generate an error group and a context with the dedicated function and use it to launch several instances of the function. We will receive this in a separate goroutine while we wait for the group. After the wait is over, we will make sure that there are no more values being sent to the channel (this would cause a panic) by waiting an extra second:

```
我们将使用专用函数生成一个错误组和一个上下文，并使用它来启动函数的几个实例。我们将在单独的Goroutine中接收此消息，同时等待团队。等待结束后，我们将等待一秒钟，以确保没有更多的值发送到通道（这将导致恐慌）：

func main() {
    eg, ctx := errgroup.WithContext(context.Background())
    ch := make(chan string)
    for i := 0; i < 10; i++ {
        eg.Go(sender(ctx, ch, i))
    }
    go func() {
        for s := range ch {
            log.Println(s)
        }
    }()
    if err := eg.Wait(); err != nil {
        log.Println("Error:", err)
    }
    close(ch)
    log.Println("waiting...")
    time.Sleep(time.Second)
}
```

As expected, thanks to the select statement within the context, the application runs seamlessly and does not panic.
正如预期的那样，多亏了上下文中的select语句，应用程序可以无缝运行，不会死机。

# Leaky bucket
We saw how to build a rate limiter using ticker in the previous chapters: by using time.Ticker to force a client to await its turn in order to get served. There is another take on rate limiting of services and libraries that's known as the leaky bucket. The name evokes an image of a bucket with a few holes in it. If you are filling it, you have to be careful to not put too much water into the bucket, otherwise it's going to overflow. Before adding more water, you need to wait for the level to drop – the speed at which this happens will depend on the size of the bucket and the number of the holes it has. We can easily understand what this concurrency pattern does by taking a look at the following analogy:

 The water going through the holes represents requests that have been completed.
The water that's overflowing from the bucket represents the requests that have been discarded.
The bucket will be defined by two attributes:

Rate: The ideal amount of requests per time if the frequency of requests is lower.
Capacity: The number of requests that can be done at the same time before the resource turns unresponsive temporarily.
The bucket has a maximum capacity, so when requests are made with a frequency higher than the rate specified, this capacity starts dropping, just like when you're putting too much water in and the bucket starts to overflow. If the frequency is zero or lower than the rate, the bucket will slowly gain its capacity, and so the water will be slowly drained.

The data structure of the leaky bucket will have a capacity and a counter for the requests that are available. This counter will be the same as the capacity on creation, and will drop each time requests are executed. The rate specifies how often the status needs to be reset to the capacity:

```
漏桶
在前面的章节中，我们看到了如何使用ticker构建一个限速器：通过使用time.ticker来强制客户等待它的轮到以获得服务。还有另一种对服务和库的速率限制，称为漏桶。这个名字让人想起一个有几个洞的桶。如果你在装水，你必须小心不要往桶里放太多的水，否则水会溢出来。在添加更多的水之前，您需要等待水位下降——下降的速度将取决于水桶的大小和水桶的孔数。通过查看以下类比，我们可以很容易地理解这种并发模式的作用：
穿过孔的水代表已完成的请求。
从桶中溢出的水表示已丢弃的请求。
桶将由两个属性定义：
速率：如果请求频率较低，则每次请求的理想数量。
容量：在资源暂时失去响应之前，可以同时执行的请求数。
桶有一个最大的容量，所以当请求的频率高于指定的速率时，这个容量就会开始下降，就像你往桶里放了太多水，桶就开始溢出一样。如果频率为零或低于速率，则桶将缓慢地获得其容量，因此水将缓慢地排出。
漏桶的数据结构将有一个容量和一个计数器，用于处理可用的请求。此计数器将与创建时的容量相同，并将在每次执行请求时删除。速率指定状态需要重置为容量的频率：

type bucket struct {
    capacity uint64
    status uint64
}
```

When creating a new bucket, we should also take care of the status reset. We can use a goroutine for this and use a context to terminate it correctly. We can create a ticker using the rate and then use these ticks to reset the status. We need to use the atomic package to ensure it is thread-safe:

```
当创建一个新的bucket时，我们还应该注意状态重置。我们可以为此使用goroutine，并使用上下文来正确地终止它。我们可以使用速率创建一个ticker，然后使用这些ticks重置状态。我们需要使用原子包来确保它是线程安全的：

func newBucket(ctx context.Context, cap uint64, rate time.Duration) *bucket {
    b := bucket{capacity: cap, status: cap}
    go func() {
        t := time.NewTicker(rate)
        for {
            select {
            case <-t.C:
                atomic.StoreUint64(&b.status, b.capacity)
            case <-ctx.Done():
                t.Stop()
                return
            }
        }
    }()
    return &b
}
```

When we're adding to the bucket, we can check the status and act accordingly:

If the status is 0, we cannot add anything.
If the amount to add is higher than the availability, we add what we can.
We add the full amount otherwise:

```
当我们添加到bucket时，我们可以检查状态并相应地执行以下操作：
如果状态为0，则无法添加任何内容。
如果要添加的数量高于可用性，我们将添加可以添加的内容。
我们加上全部金额，否则：

func (b *bucket) Add(n uint64) uint64 {
    for {
        r := atomic.LoadUint64(&b.status)
        if r == 0 {
            return 0
        }
        if n > r {
            n = r
        }
        if !atomic.CompareAndSwapUint64(&b.status, r, r-n) {
            continue
        }
        return n
    }
}
```

We are using a loop to try atomic swap operations until they succeed to ensure that what we get with the Load operation doesn't change when we are doing a compare and swap (CAS).

The bucket can be used in a client that will try to add a random amount to the bucket and will log its result:

```
我们使用一个循环来尝试原子交换操作，直到它们成功为止，以确保在进行比较和交换（CAS）时，加载操作所获得的内容不会改变。
该存储桶可用于将尝试向存储桶中添加随机数量并记录其结果的客户端：


type client struct {
    name string
    max int
    b *bucket
    sleep time.Duration
}

func (c client) Run(ctx context.Context, start time.Time) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            n := 1 + rand.Intn(c.max-1)
            time.Sleep(c.sleep)
            e := time.Since(start).Seconds()
            a := c.b.Add(uint64(n))
            log.Printf("%s tries to take %d after %.02fs, takes  
                %d", c.name, n, e, a)
        }
    }
}
```

We can use more clients concurrently so that having concurrent access to resources will have the following result:

Some goroutines will be adding what they expect to the bucket.
One goroutine will finally fill the bucket by adding a quantity that is equal to the remaining capacity, even if the amount that they are trying to add is higher.
The other goroutines will not be able to add to the bucket until the capacity is reset:

```
我们可以同时使用更多的客户机，这样并发访问资源将产生以下结果：
一些Goroutine将把他们期望的添加到桶中。
一个goroutine最终会通过添加一个等于剩余容量的数量来填充bucket，即使他们尝试添加的数量更高。
在重置容量之前，其他goroutine将无法添加到存储桶：


func main() {
    ctx, canc := context.WithTimeout(context.Background(), time.Second)
    defer canc()
    start := time.Now()
    b := newBucket(ctx, 10, time.Second/5)
    t := time.Second / 10
    for i := 0; i < 5; i++ {
        c := client{
            name: fmt.Sprint(i),
            b: b,
            sleep: t,
            max: 5,
        }
        go c.Run(ctx, start)
    }
    <-ctx.Done()
}
```

Sequencing
In concurrent scenarios with multiple goroutines, we may need to have a synchronization between goroutines, such as in a scenario where each goroutine needs to wait for its turn after sending. 

A use case for this scenario could be a turn-based application wherein different goroutines are sending messages to the same channel, and each one of them has to wait until all the others have finished before they can send it again.

A very simple implementation of this scenario can be obtained using private channels between the main goroutine and the senders. We can define a very simple structure that carries both messages and a Wait channel. It will have two methods – one for marking the transaction as done and another one that waits for such a signal – when it uses a channel underneath. The following method shows this:

```
排序
在具有多个goroutine的并发场景中，我们可能需要在goroutine之间进行同步，例如在发送后每个goroutine需要等待轮到它的场景中。
这个场景的一个用例可以是一个基于turn的应用程序，其中不同的goroutine将消息发送到同一个通道，并且每个goroutine必须等到其他所有goroutine完成后才能再次发送消息。
这个场景的一个非常简单的实现可以通过主goroutine和发送者之间的专用通道获得。我们可以定义一个非常简单的结构，它同时承载消息和等待通道。当它使用下面的通道时，它将有两种方法——一种是将事务标记为已完成，另一种是等待这样的信号。以下方法显示了这一点：

type msg struct {
    value string
    done chan struct{}
}

func (m *msg) Wait() {
    <-m.done
}

func (m *msg) Done() {
    m.done <- struct{}{}
}
```

We can create a source of messages with a generator. We can use a random delay with the send operation. After each send, we wait for the signal that is obtained by calling the Done method. We always use context to keep everything free from leaks:
我们可以使用生成器创建消息源。我们可以对发送操作使用随机延迟。每次发送后，我们都等待通过调用done方法获得的信号。我们始终使用上下文来防止所有内容泄漏：

```
func send(ctx context.Context, v string) <-chan msg {
    ch := make(chan msg)
    go func() {
        done := make(chan struct{})
        for i := 0; ; i++ {
            time.Sleep(time.Duration(float64(time.Second/2) * rand.Float64()))
            m := msg{fmt.Sprintf("%s msg-%d", v, i), done}
            select {
            case <-ctx.Done():
                close(ch)
                return
            case ch <- m:
                m.Wait()
            }
        }
    }()
    return ch
}
```

We can use a fan-in to put all of the channels into one, singular channel:
我们可以使用扇入将所有通道放在一个单独的通道中：

```
func merge(ctx context.Context, sources ...<-chan msg) <-chan msg {
    ch := make(chan msg)
    go func() {
        <-ctx.Done()
        close(ch)
    }()
    for i := range sources {
        go func(i int) {
            for {
                select {
                case v := <-sources[i]:
                    select {
                    case <-ctx.Done():
                        return
                    case ch <- v:
                    }
                }
            }
        }(i)
    }
    return ch
}
```

The main application will be receiving from the merged channel until it's closed. When it receives one message from each channel, the channel will be blocked, waiting for the Done method signal to be called by the main goroutine.

This specific configuration will allow the main goroutine to receive just one message from each channel. When the message count reaches the number of goroutines, we can call Done from the main goroutine and reset the list so that the other goroutines will be unlocked and be able to send messages again:

```
主应用程序将从合并的通道接收，直到关闭。当它从每个通道接收到一条消息时，通道将被阻塞，等待主goroutine调用done方法信号。
这种特定的配置将允许主goroutine仅从每个通道接收一条消息。当消息计数达到goroutine的数目时，我们可以从主goroutine调用done并重置列表，以便其他goroutine将被解锁并能够再次发送消息：

func main() {
    ctx, canc := context.WithTimeout(context.Background(), time.Second)
    defer canc()
    sources := make([]<-chan msg, 5)
    for i := range sources {
        sources[i] = send(ctx, fmt.Sprint("src-", i))
    }
    msgs := make([]msg, 0, len(sources))
    start := time.Now()
    for v := range merge(ctx, sources...) {
        msgs = append(msgs, v)
        log.Println(v.value, time.Since(start))
        if len(msgs) == len(sources) {
            log.Println("*** done ***")
            for _, m := range msgs {
                m.Done()
            }
            msgs = msgs[:0]
            start = time.Now()
        }
    }
}
```

Running the application will result in all the goroutines sending a message to the main one once. Each of them will be waiting for everyone to send their message. Then, they will start sending messages again. This results in messages being sent in rounds, as expected. 

```
运行应用程序将导致所有goroutine向主应用程序发送一次消息。他们每个人都会等待每个人发送他们的信息。然后，他们将再次开始发送消息。这将导致消息按预期的顺序发送。
```

Summary
In this chapter, we looked at some specific concurrency patterns for our applications. We learned that generators are functions that return channels, and also feed such channels with data and close them when there is no more data. We also saw that we can use a context to allow the generator to exit early.

Next, we focused on pipelines, which are stages of execution that use channels for communication. They can either be source, which doesn't require any input; destination, which doesn't return a channel; or intermediate, which receives a channel as input and returns one as output.

Another pattern is the multiplexing and demultiplexing one, which consists of spreading a channel to different goroutines and combining several channels into one. It is often referred to as fan-out fan-in, and it allows us to execute different operations concurrently on a set of data.

Finally, we learned how to implement a better version of the rate limiter called leaky bucket, which limits the number of requests in a specific amount of time. We also looked at the sequencing pattern, which uses a private channel to signal to all of the sending goroutines when they are allowed to send data again.

In this next chapter, we are going to introduce the first of two extra topics that were presented in the Sequencing section. It is here that we will demonstrate how to use reflection to build generic code that adapts to any user-provided type.

总结
在本章中，我们研究了应用程序的一些特定并发模式。我们了解到生成器是返回通道的函数，也向这些通道提供数据，并在没有更多数据时关闭它们。我们还看到，我们可以使用上下文来允许生成器提前退出。
接下来，我们将重点放在管道上，管道是使用通道进行通信的执行阶段。它们可以是源，不需要任何输入；目的地，不返回通道；或者中间，接收通道作为输入，返回通道作为输出。
另一种模式是多路复用和解复用模式，它包括将一个信道扩展到不同的goroutine，并将多个信道组合成一个信道。它通常被称为扇出扇入，它允许我们对一组数据同时执行不同的操作。
最后，我们学习了如何实现一个更好版本的速率限制器LeakyBucket，它在特定的时间内限制了请求的数量。我们还研究了排序模式，它使用一个专用通道向所有发送goroutine发送信号，当它们被允许再次发送数据时。
在下一章中，我们将介绍排序部分中介绍的另外两个主题中的第一个。在这里，我们将演示如何使用反射来构建适合任何用户提供类型的通用代码。


# Questions
What is a generator? What are its responsibilities?
How could you describe a pipeline?
What type of stage gets a channel and returns one?
What is the difference between fan-in and fan-out?


问题
什么是生成器？它的职责是什么？
你如何描述一个管道？
哪种类型的舞台可以获得频道并返回频道？
扇入和扇出有什么区别？

