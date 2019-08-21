
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
```


