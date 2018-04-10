import threading
import time


def myChildThread():
    print("Child Thread Starting")
    time.sleep(5)
    print("Current Thread ----------")
    print(threading.current_thread())
    print("-------------------------")
    print("Main Thread -------------")
    print(threading.main_thread())
    print("-------------------------")
    print("Child Thread Ending")


child = threading.Thread(target=myChildThread)
child.start()
child.join()


"""
把它分解
在前面的代码中，我们定义了一个名为myChildThread的简单函数。这将是我们为演示目的创建的线程对象的目标。
在这个函数中，我们简单地打印当前线程和主线程。

然后，我们继续创建一个线程对象，然后启动并加入这个新创建的线程。在输出中，您应该看到如下内容:

"""