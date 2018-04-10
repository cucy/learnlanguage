import threading
import time
import random


def myThread(i):
    print("Thread {}: started".format(i))
    time.sleep(random.randint(1, 5))
    print("Thread {}: finished".format(i))


def main():
    for i in range(4):
        thread = threading.Thread(target=myThread, args=(i,))
        thread.start()
    print("Enumerating: {}".format(threading.enumerate()))


if __name__ == '__main__':
    main()


"""
把它分解
在前面的示例中，我们首先定义一个非常简单的函数myThread，它将是我们将要创建的线程的目标。在这个函数中，
我们只需打印线程已经启动，然后在打印前等待1到5秒之间的随机间隔。

然后我们定义了一个主函数，它创建了四个不同的线程对象，然后启动它们。
一旦我们完成了创建和启动这些线程，我们就会打印出thread .enumerate()的结果，它应该输出如下内容:
"""
