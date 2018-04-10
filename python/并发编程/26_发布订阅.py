"""
我们的出版商
我们的发布者类有两个定义在其中的函数，它包含了整数数组和条件原语的引用。

运行函数本质上是在调用时进入一个永久循环，然后在0到1000之间生成一个随机整数。一旦我们生成了这个数字，我们就会获得条件，然后将这个新生成的整数追加到我们的整数数组中。

在我们添加到数组之后，我们首先通知我们的订阅者，在这个数组中添加了一个新项，然后我们释放这个条件。
"""
import random
import threading

import time


class Publisher(threading.Thread):
    def __init__(self, integers, condition):
        self.condition = condition
        self.integers = integers
        threading.Thread.__init__(self)

    def run(self):
        while True:
            integer = random.randint(0, 1000)
            self.condition.acquire()
            print("Condition Acquired by Publisher: {}".format(self.name))
            self.integers.append(integer)
            self.condition.notify()
            print("Condition Released by Publisher: {}".format(self.name))
            self.condition.release()
            time.sleep(1)


"""
我们的用户
同样，订阅者类有两个定义在其中的函数:构造函数和运行函数。构造函数包含两个东西，第一个是它将要消耗的整数数组的引用，第二个是条件同步原语。

在我们的运行函数中，我们启动一个循环，该循环不断尝试获取传入它的条件。当我们设法获取这个锁时，我们打印出线程已经获得它的事实，然后我们继续尝试“弹出”第一个整数，从我们传入的整数数组中。一旦我们成功地管理了这个，我们就会释放条件原语，并且，再一次，开始尝试重新获得这个条件。
"""


class Subscriber(threading.Thread):
    def __init__(self, integers, condition):
        self.integers = integers
        self.condition = condition
        threading.Thread.__init__(self)

    def run(self):
        while True:
            self.condition.acquire()
            print("Condition Acquired by Consumer: {}".format(self.name))
            while True:
                if self.integers:
                    integer = self.integers.pop()
                    print("{} Popped from list by Consumer: {}".format(integer, self.name))
                    break
                print("Condition Wait by {}".format(self.name))
                self.condition.wait()
            print("Consumer {} Releasing Condition".format(self.name))
            self.condition.release()


"""
踢掉
在该程序的主要功能中，我们首先声明将充当消息队列的整数数组。然后我们声明我们的条件原语。

最后，我们定义一个发布者和两个不同的订阅者。然后，我们启动这些发布者和订阅者并加入线程，这样我们的程序就不会在线程有机会执行之前立即终止。

"""

def main():
     integers = []
     condition = threading.Condition()
     # Our Publisher
     pub1 = Publisher(integers, condition)
     pub1.start()
     # Our Subscribers
     sub1 = Subscriber(integers, condition)
     sub2 = Subscriber(integers, condition)
     sub1.start()
     sub2.start()
     ## Joining our Threads
     pub1.join()
     consumer1.join()
     consumer2.join()
if __name__ == '__main__':
    main()


"""
当我们运行这个程序时，您应该看到一个类似于下面的输出。您应该看到，当发布者获得条件时，它向数组追加一个数字，然后通知条件并释放它。

823 Popped from list by Consumer: Thread-3
Consumer Thread-3 Releasing Condition
Condition Acquired by Consumer: Thread-3
Condition Wait by Thread-3
Condition Acquired by Publisher: Thread-1
Condition Released by Publisher: Thread-1
262 Popped from list by Consumer: Thread-2
Consumer Thread-2 Releasing Condition
Condition Acquired by Consumer: Thread-2
Condition Wait by Thread-2
Condition Acquired by Publisher: Thread-1
Condition Released by Publisher: Thread-1
685 Popped from list by Consumer: Thread-3
Consumer Thread-3 Releasing Condition
Condition Acquired by Consumer: Thread-3
Condition Wait by Thread-3

在通知的条件下，两个订阅者之间的战斗开始，他们都试图首先获得这个条件。当一个人赢得这场战斗，它就会简单地从数组中“弹出”这个数字。
"""