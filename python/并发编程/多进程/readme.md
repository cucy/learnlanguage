# 进程间通信
当涉及到多个子进程之间的同步时，我们有许多不同的选择可以利用:

队列:这是您的标准FIFO队列，在第5章中讨论了线程之间的通信。
管道:这是一个新概念，稍后我们会详细介绍。
Manager:这些为我们提供了一种创建数据的方法，然后，在我们的Python应用程序中，在不同的进程之间共享这些数据。
Ctypes:这些对象使用一个共享内存，这些内存随后可以被子进程访问。
前面提到的四个选项代表了在多进程应用程序中可以使用的非常多的不同通信机制。为了确保你在交流时做出正确的选择，在这个话题上花点时间是绝对值得的。

不幸的是，这本书没有足够的空间来涵盖几乎无限的可能的解决方案，这些解决方案可以使用这些选项来实现。我建议，如果您对深入到这个主题有兴趣，那么您应该选择高性能Python，由Ian Ozsvald;Micha Gorelick，因为它有一个很好的章节，更详细地涵盖了这些概念。


# Communicating sequential processes
Communicating Sequential Processes, or CSP for short, is used to describe how systems that feature multiple concurrent models should interact with one another. It, typically, relies heavily on using channels as a medium for passing messages between two or more concurrent processes, and is the underlying mantra of languages like clojure and golang.

It's a concept that is certainly growing in popularity, and there are a number of fantastic talks and books on CSP that you should definitely check out.

I'd recommend checking out Communicating Sequential Processes by C.A.R. Hoare, which was published in May of 2015, the link for which is this: http://www.usingcsp.com/cspbook.pdf.
After some brief research on the topic, it's fascinating to see how certain problem sets can be abstracted out quite easily using this style of programming as opposed to your more traditional object-oriented setups.