A student sent me the below note so I included the student's excellent solution to this project:

Todd, I reviewed your solution to 22_go-routines/04_challenge-solution in your Golang Programming class.

I believe your solution might not run 100 factorial computations concurrently and in parallel as the following statement will calculate the factorials sequentially, since it is receiving them - one by one -  from the pipeline:
out <- fact(n)

I used the sync package to create a solution that I believe will accomplish your goal. Let me know if I am missing something:

https://github.com/arsanjea/udemyTraining/blob/master/Exercises/exercise38.go

/////

# Definitions

## concurrency
a design pattern
go uses goroutines to create concurrent design patterns

## parallelism
running code from a program on more than one cpu
parallelism implies you have used concurrent design patterns

## sequentially
one thing happening after another, in sequence

# here are my thoughts

The original solution uses concurrent design patterns. It uses two different goroutines. The [control flow](https://en.wikipedia.org/wiki/Control_flow) of the program is no longer a straight top-to-bottom sequence. Different goroutines have been launched.

The program may or may not run in parallel. If the program is run on a machine with two or more cpus, the program has the potential to run in parallel. Each of our three goroutines (the two goroutines we launched, and main) could be running on different cpu cores.

Jean-Marc Arsan is correct: the program IS still running sequentially. Even though calculations are occuring in different goroutines, and potentially on different CPU cores, the sequence in which they occur is still sequential.

goroutines allow synchronization of code.

In this original example, the code is synchronized.

Thank you, Jean-Marc, for your comment on this code!

I appreciate these discussions and hope the notes and your new code sample are helpful to everyone!






一个学生给我发了下面的说明，所以我把这个学生的优秀解决方案包括在这个项目中:

托德，我在你的Golang编程课上复习了你的解决方案:22_go /04_challeng -solution。

我相信您的解决方案可能不会同时运行100的阶乘计算，并且并行的，因为下面的语句将按顺序计算阶乘，因为它是一个接一个地从管道接收它们:
<——事实(n)

我使用同步包来创建一个解决方案，我相信这将实现您的目标。如果我遗漏了什么，请告诉我:

https://github.com/arsanjea/udemyTraining/blob/master/Exercises/exercise38.go

/ / / / /

#定义

# #并发
一种设计模式
go使用goroutines创建并发的设计模式。

# #并行性
在多个cpu上运行程序的代码。
并行性意味着您使用了并发设计模式。

# #顺序
一件事一件接一件地发生。

这是我的想法。

原始解决方案使用并发设计模式。它使用了两个不同的goroutines。该程序的[控制流](https://en.wikipedia.org/wiki/Control_flow)不再是一个直接的从上到下的顺序。不同的goroutines已经发射。

程序可以并行运行，也可以不并行运行。如果程序运行在具有两个或多个cpu的机器上，程序就有可能并行运行。我们的三个goroutines(我们推出的两个goroutines和main)都可以运行在不同的cpu内核上。

Jean-Marc Arsan是正确的:程序仍在按顺序运行。尽管计算在不同的goroutines中发生，并且可能在不同的CPU内核中发生，但是它们发生的顺序仍然是顺序的。

goroutines允许代码同步。

在这个原始示例中，代码是同步的。

谢谢你，Jean-Marc，谢谢你对这段代码的评论!

我很感激这些讨论，希望这些注释和您的新代码示例对每个人都有帮助!