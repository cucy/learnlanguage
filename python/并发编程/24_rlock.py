import threading
import time


class myWorker():


    def __init__(self):
        self.a = 1
        self.b = 2
        self.Rlock = threading.RLock()


    def modifyA(self):
        with self.Rlock:
            print("Modifying A : RLock Acquired:  {}".format(self.Rlock._is_owned()))
            print("{}".format(self.Rlock))
            self.a = self.a + 1
            time.sleep(5)


    def modifyB(self):
        with self.Rlock:
            print("Modifying B : RLock Acquired:  {}".format(self.Rlock._is_owned()))
            print("{}".format(self.Rlock))
            self.b = self.b - 1
            time.sleep(5)


    def modifyBoth(self):
        with self.Rlock:
            print("Rlock acquired, modifying A and B")
            print("{}".format(self.Rlock))
            self.modifyA()
            self.modifyB()
        print("{}".format(self.Rlock))


workerA = myWorker()
workerA.modifyBoth()


"""
Rlock acquired, modifying A and B
<locked _thread.RLock object owner=140735793988544 count=1 at 
0x10296e6f0>
Modifying A : RLock Acquired: True
<locked _thread.RLock object owner=140735793988544 count=2 at 
0x10296e6f0>
<locked _thread.RLock object owner=140735793988544 count=1 at 
0x10296e6f0>
Modifying B : RLock Acquired: True
<locked _thread.RLock object owner=140735793988544 count=2 at 
0x10296e6f0>
<unlocked _thread.RLock object owner=0 count=0 at 0x10296e6f0>

把它分解
在前面的代码中，我们看到了一个典型的例子，就是一个RLock在我们的单线程程序中工作的方式。我们定义了一个名为myWorker的类，它具有四个函数，它们是初始化我们的Rlock和a和b变量的构造函数。

然后我们继续定义两个分别修改a和b的函数。它们都首先使用with语句获取类Rlock，然后对内部变量执行任何必要的修改。

最后，我们有了modifyBoth函数，它在调用modifyA和modifyB函数之前执行初始的Rlock捕获。

在每一步中，我们打印出Rlock的状态。我们看到，当它在modifyBoth函数中获得之后，它的所有者被设置为主线程，并且它的计数增加为1。当我们下一次调用modifyA时，Rlocks计数器会再次增加一个，并且在modifyA之前进行必要的计算，然后释放Rlock。在Rlock的modifyA函数释放上，我们看到计数器减量为1，然后再通过我们的modifyB函数立即递增为2。

最后，当modifyB完成它的执行时，它会释放Rlock，然后，我们的modifyBoth函数也会释放。当我们从Rlock对象中进行最后打印时，我们看到所有者被设置为0，并且我们的计数也被设置为0。从理论上讲，只有在这一点上，另一个线程才能获得这个锁。

"""