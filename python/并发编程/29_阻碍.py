import threading
import time
import random


class myThread(threading.Thread):
    def __init__(self, barrier):
        threading.Thread.__init__(self)
        self.barrier = barrier

    def run(self):
        print("Thread {} working on something".format(threading.current_thread()))
        time.sleep(random.randint(1, 10))
        print("Thread {} is joining {} waiting on Barrier".format(threading.current_thread(),
                                                                  self.barrier.n_waiting))
        self.barrier.wait()

        print("Barrier has been lifted, continuing with work")


barrier = threading.Barrier(4)
threads = []
for i in range(4):
    thread = myThread(barrier)
    thread.start()
    threads.append(thread)
for t in threads:
    t.join()


"""
把它分解
如果我们看一下前面的代码，我们已经定义了一个自定义类myThread，它继承了thread . thread。在这个类中，
我们定义了标准__init__函数和run函数。我们的__init__函数接受我们的barrier对象，以便以后可以引用它。

在我们的运行函数中，我们模拟我们的线程在1到10秒之间随机地做一些工作，然后我们开始在barrier上等待。

根据我们的类定义，我们首先通过调用barrier = thread . barrier(4)来创建barrier对象。
我们将其作为一个参数传递的4表示在它将被取消之前必须在barrier上等待的线程数。然后我们继续定义四个不同的线程，并将它们全部连接起来。



输出
如果您在您的系统上运行前面的程序，您应该希望看到类似如下的输出。

你会看到我们的四个线程打印出它们正在处理某个东西，然后，一个接一个地，它们会随机地开始等待我们的barrier对象。一旦第4个线程开始等待，程序几乎立即完成，因为所有的四个线程都完成了最后的打印声明，现在这个屏障已经被解除了。
"""