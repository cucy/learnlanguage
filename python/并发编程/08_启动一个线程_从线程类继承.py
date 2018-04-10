"""
对于需要更多代码的场景，我们可以定义一个从线程本地类继承目录的类。

这是理想的场景，因为代码的复杂性对于单个函数来说太大了，而需要将其分解为多个函数。虽然这确实给我们在处理线程时提供了更大的灵活性，但是我们必须考虑到我们现在必须在这个类中管理线程的事实。

为了让我们定义一个新的线程，子类化原生Python线程类，我们需要至少做以下操作:

将线程类传递给我们的类定义。
为了让线程初始化，在构造函数中调用thread .__init__(self)。
定义一个run()函数，该函数将在线程启动时调用:

"""

from threading import Thread


class my_work_thread(Thread):
    def __init__(self):
        print("hello world")
        Thread.__init__(self)

    def run(self):
        print("Thread is now running")


myThread = my_work_thread()
print("Created my Thread Object")
myThread.start()
print("Started my thread")
myThread.join()
print("My Thread finished")



"""
把它分解
在前面的代码中，我们定义了一个非常简单的类myWorkerThread，它继承了线程类。在我们的构造函数中，我们调用必要的线程。__init__(self)函数。

然后，我们还定义了当我们开始写神话时将会调用的run函数。在这个运行函数中，我们简单地调用print函数将我们的状态打印到控制台，
然后我们的线程实际上终止了。

"""