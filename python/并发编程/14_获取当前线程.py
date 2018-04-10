import threading
import time


def threadTarget():
    print("Current Thread: {}".format(threading.current_thread()))


threads = []
for i in range(10):
    thread = threading.Thread(target=threadTarget)
    thread.start()
    threads.append(thread)
for thread in threads:
    thread.join()



"""
把它分解
在前面的示例中，我们定义了一个函数threadTarget，它打印出当前线程。然后，我们继续为线程创建一个空的数组，
然后用10个不同的线程对象填充这个数组。然后我们依次加入这些线程，这样我们的程序就不会立即退出。前一个程序的输出应该是这样的:
"""