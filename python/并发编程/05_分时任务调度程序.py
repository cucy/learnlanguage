import threading

import time
import random

counter = 1


def workerA():
    global counter

    while counter < 1000:
        counter += 1
        print(f"Worker A +  is incrementing counter to {counter}")
        seelp_time = random.randint(0, 1)
        time.sleep(seelp_time)

def workerB():
    global counter

    while counter > -1000:
        counter -= 1
        print(f"Worker B - is decrementing  counter to {counter}")
        seelp_time = random.randint(0, 1)
        time.sleep(seelp_time)



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
在前面的代码中，我们在Python中有两个相互竞争的线程，每个线程都试图实现各自的目标，要么将计数器减到1,000，要么反过来将其递增到1,000。在单核处理器,有工人管理完成其任务的可能性之前,工人有机会执行,和工人也是如此。然而,还有第三个潜在的可能性,那就是任务调度器继续工人A和工人B之间切换B无限次数和从来没有完成。

顺便说一下，前面的代码也显示了多线程访问共享资源的危险，没有任何形式的同步。没有准确的方法来确定我们的计数器将会发生什么，因此，我们的程序可能被认为是不可靠的。
"""