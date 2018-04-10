import threading
import time


def myThread():
    print("Thread {} starting".format(threading.currentThread().getName()))
    time.sleep(10)
    print("Thread {} ending".format(threading.currentThread().getName()))


for i in range(4):
    threadName = "线程Thread-" + str(i)
    thread = threading.Thread(name=threadName, target=myThread)
    thread.start()
print("{}".format(threading.enumerate()))



"""
分解
在前面的代码中，我们主要做的是定义一个名为myThread的函数。在这个函数中，我们使用thread . currentthread (). getname () getter
方法来检索当前线程的名称，并在线程执行时和结束时将其打印出来。
然后，我们继续启动for循环，并创建4个线程对象，它们接受name参数，我们将其定义为“thread -”+ str(i)，以及myThread函数作为该线程执行的目标。最后，我们继续打印当前正在运行的所有活动线程。这应该打印出如下内容:

"""