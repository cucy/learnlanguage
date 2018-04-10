"""
启动一个线程
在Python中，有许多不同的方法可以启动线程。当我们有一个相对简单的任务时，我们希望多线程，然后我们可以选择将它定义为一个单一的函数。在下面的例子中，我们有一个非常简单的函数，它只在一个随机的时间间隔内睡眠。这是一个非常简单的功能，它是封装一个简单函数的理想，然后将这个简单的函数作为新线程的目标。线程对象如下代码所示:
"""

import threading
import time
import random


def executeThread(i):
    print(f"thread {i} started")
    sleep_time = random.randint(1, 10)
    time.sleep(sleep_time)
    print(f"thread {i} finished executing")


for i in range(10):
    thread = threading.Thread(target=executeThread, args=(i,))
    thread.start()
    print("Active Threads:", threading.enumerate())
