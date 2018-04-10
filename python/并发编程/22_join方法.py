"""
的连接方法
在开发不可思议的重要企业系统时，能够决定某些任务的执行顺序是非常重要的。幸运的是，Python的线程对象允许我们保留某种形式的控制，因为它们附带了一个连接方法。

实际上，join方法阻止父线程继续前进，直到线程确认它已经终止。这可能是通过自然的结束，或者是当线程抛出一个未处理的异常时。让我们通过下面的例子来理解这一点:
"""

import threading
import time


def ourThread(i):
    print("Thread {} Started".format(i))
    time.sleep(i * 2)
    print("Thread {} Finished".format(i))


def main():
    thread1 = threading.Thread(target=ourThread, args=(1,))
    thread1.start()
    print("Is thread 1 Finished?")
    thread2 = threading.Thread(target=ourThread, args=(2,))
    thread2.start()
    thread2.join()
    print("Thread 2 definitely finished")


if __name__ == '__main__':
    main()



"""
把它分解
前面的代码示例展示了如何利用这个连接方法使线程程序的流具有一定的确定性。

我们首先定义一个非常简单的函数myThread，它接受一个参数。所有这些函数都是在它启动时打印出来的，不管什么值被传递到它的时间2，然后在它完成执行时打印出来。

在我们的main函数中，我们定义了两个线程，第一个线程我们将其称为thread1，并将值1作为其惟一参数传递。然后我们启动这个线程并执行print语句。需要注意的是，第一个print语句在我们的thread1完成之前执行。

然后我们创建第二个线程对象，并以想象的方式调用这个thread2，并以2作为唯一的参数。关键的区别在于，我们在启动线程后立即调用thread2.join()。通过调用thread2，我们可以保存执行打印语句的顺序，在线程2结束后，您可以在输出中看到线程2确实完成了。



把它放在一起
虽然join方法可能非常有用，并且为您提供了一种快速而干净的方法来确保代码内的顺序，但也很重要的一点是，您可以通过将我们的代码在第一个地方进行多线程处理，从而消除我们所取得的所有成果。

在前面的示例中考虑我们的thread2对象——通过多线程，我们究竟得到了什么?我知道这是一个相当简单的程序，但关键是我们在启动它之后立即加入它，并且基本上，阻塞了我们的主线程，直到thread2完成了它的执行。实际上，在执行thread2的过程中，我们将多线程应用程序单线程化。
"""