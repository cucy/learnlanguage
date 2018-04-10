import threading
import queue
import random
import time


def mySubscriber(queue):
    while not queue.empty():
        item = queue.get()
        if item is None:
            break
        print("{} removed {} from the queue".format(threading.current_thread(), item))
        queue.task_done()
        time.sleep(1)


myQueue = queue.Queue()
for i in range(10):
    myQueue.put(i)

print("Queue Populated")

threads = []
for i in range(4):
    thread = threading.Thread(target=mySubscriber, args=(myQueue,))
    thread.start()
    threads.append(thread)

for thread in threads:
    thread.join()

"""
把它分解
在前面的示例中，我们导入了必要的队列Python模块。然后，我们继续定义一个mySubscriber函数，
它将作为我们的多个线程的目标，这些线程将从我们的队列中消耗。

在下面的mySubscriber函数声明中，我们通过调用myQueue = queue.Queue()来声明我们的队列，然后我们继续填充从0到9的数字。

最后，我们继续声明并实例化将从线程安全队列中消耗的大量线程。我们启动这些线程，注意在新声明的队列对象中传入它们的args，然后，随后加入它们。

"""
