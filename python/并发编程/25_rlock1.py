import threading
import time


class myWorker():
    def __init__(self):
        self.a = 1
        self.b = 2
        self.lock = threading.Lock()

    def modifyA(self):
        with self.lock:
            print("Modifying A : RLock Acquired: {}".format(self.lock._is_owned()))
            print("{}".format(self.lock))
            self.a = self.a + 1
            time.sleep(5)

    def modifyB(self):
        with self.lock:
            print("Modifying B : Lock Acquired: {}".format(self.lock._is_owned()))
            print("{}".format(self.lock))
            self.b = self.b - 1
            time.sleep(5)

    def modifyBoth(self):
        with self.lock:
            print("lock acquired, modifying A and B")
            print("{}".format(self.lock))
            self.modifyA()
            print("{}".format(self.lock))
            self.modifyB()
        print("{}".format(self.lock))


workerA = myWorker()
workerA.modifyBoth()


"""
RLocks与常规锁
如果我们尝试使用传统的锁原语来执行相同的前一个程序，那么您应该注意到程序实际上从来没有到达执行我们的modifyA()函数的点。
实际上，我们的程序会陷入一种僵局，因为我们还没有实现一种允许我们的线程继续前进的释放机制。这在以下代码示例中显示:

本质上，RLocks允许我们以递归方式获得某种形式的线程安全，而无需实现复杂的获取，并在代码中释放锁逻辑。
它们允许我们编写更简单的代码，从而更易于遵循，因此，在代码生成之后，更容易维护。

"""