"""
在前面的示例中，我们看到两个线程在不断地竞争以增加或减少计数器。通过添加锁，
我们可以确保这些线程能够以确定和安全的方式访问我们的计数器。
"""

import threading
import time
import random

counter = 1
lock = threading.Lock()


def workerA():
    global counter
    lock.acquire()  # 上锁
    try:
        while counter < 1000:
            counter += 1
            print("Worker A is incrementing counter to {}".format(counter))
            sleepTime = random.randint(0, 1)
            time.sleep(sleepTime)
    finally:
        lock.release()  # 释放锁


def workerB():
    global counter
    lock.acquire()
    try:
        while counter > -1000:
            counter -= 1
            print("Worker B is decrementing counter to {}".format(counter))
            sleepTime = random.randint(0, 1)
            time.sleep(sleepTime)
    finally:
        lock.release()


def main():
    t0 = time.time()
    thread1 = threading.Thread(target=workerA)
    thread2 = threading.Thread(target=workerB)
    thread1.start()
    thread2.start()
    thread1.join()
    thread2.join()
    t1 = time.time()
    print("Execution Time {}".format(t1 - t0))


if __name__ == '__main__':
    main()


"""
把它分解
在前面的代码中，我们添加了一个非常简单的锁原语，它封装了两个worker函数中的while循环。当线程第一次启动时，它们都争着获取锁，以便执行它们的目标，并尝试将计数器增加到1,000或-1,000，而不必与其他线程竞争。只有在一个线程完成它们的目标并释放锁之后，另一个线程才能获得该锁，并尝试增加或递减计数器。

前面的代码将执行得非常缓慢，因为它主要用于演示目的。如果您在while循环中删除了time.sleep()调用，那么您应该会注意到这段代码几乎是立即执行的。



"""