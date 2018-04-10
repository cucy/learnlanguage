"""
TicketSeller类
首先，我们将实现我们的TicketSeller类。这个类将包含它自己的内部计数器，它已经卖出了多少张票。在我们的构造函数中，我们初始化线程并获取信号量的引用。在我们的运行函数中，我们尝试获得这个信号量，如果我们可供出售的票数量小于或等于0;如果它大于0，那么我们就增加ticketSeller卖出的票的数量，并减少1。然后我们释放信号量并打印出我们的进度。
"""
import random
import threading

import time


class TicketSeller(threading.Thread):
    ticketsSold = 0

    def __init__(self, semaphore):
        threading.Thread.__init__(self)
        self.sem = semaphore
        print("Ticket Seller Started Work")

    def run(self):
        global ticketsAvailable
        running = True
        while running:
            self.randomDelay()

            self.sem.acquire()
            if (ticketsAvailable <= 0):
                running = False
            else:
                self.ticketsSold = self.ticketsSold + 1
                ticketsAvailable = ticketsAvailable - 1
                print("{} Sold One ({} left)".format(self.getName(), ticketsAvailable))
            self.sem.release()
        print("Ticket Seller {} Sold {} tickets in total".format(self.getName(), self.ticketsSold))

    def randomDelay(self):
        time.sleep(random.randint(0, 1))


# our sempahore primitive
semaphore = threading.Semaphore()
# Our Ticket Allocation
ticketsAvailable = 10
# our array of sellers
sellers = []
for i in range(4):
    seller = TicketSeller(semaphore)
    seller.start()
    sellers.append(seller)
# joining all our sellers
for seller in sellers:
    seller.join()



"""
在前面的代码中，我们定义了TicketSeller类。这个类具有一个构造函数，它接受我们的全局信号量对象的引用，并初始化我们的线程。在我们的运行函数中，我们定义了一个while循环，它在0到1秒之间模拟阻塞，然后尝试获取信号量。在成功获得信号量之后，它会检查是否有出售的票。如果有的话，它会增加ticketsSold和decments ticketsAvailable的数量，然后将它打印到控制台。

现在我们已经定义了TicketSeller类，我们需要首先创建我们的信号量对象，该对象将被传递给TicketSerllers的所有实例，如下所示:

输出
在运行前面的程序时，希望看到类似如下的输出。在这个特别的运行中，我们看到了在四个不同的线程之间销售的几乎均匀的票。当其中一个线程阻塞了一个不确定的时间量时，另一个线程就会获得信号量并尝试出售他们的票。

"""