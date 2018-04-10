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


myQueue = queue.LifoQueue()

for i in range(10):
    myQueue.put(i)

print("Queue Populated")

threads = []
for i in range(2):
    thread = threading.Thread(target=mySubscriber, args=(myQueue,))
    thread.start()
    threads.append(thread)

for thread in threads:
    thread.join()

print("Queue is empty")


"""
把它分解
前面的代码与我们用于正常FIFO队列的代码并不相似。惟一真正的区别是，当我们声明myQueue对象时，我们使用queue.LifoQueue()来声明它，而不是使用正常的queue.Queue()。
"""