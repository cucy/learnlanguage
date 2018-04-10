import time
import random
import threading


def calculatePrimeFactors(n):
    primfac = []
    d = 2
    while d * d <= n:
        while (n % d) == 0:
            primfac.append(d)
            n //= d
        d += 1
    if n > 1:
        primfac.append(n)
    return primfac


def executeProc():
    for i in range(1000):
        rand = random.randint(20000, 100000000)
        print(calculatePrimeFactors(rand))


def main():
    print("Starting number crunching")
    t0 = time.time()

    threads = []

    for i in range(10):
        thread = threading.Thread(target=executeProc)
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()

    t1 = time.time()
    totalTime = t1 - t0
    print("Execution Time: {}".format(totalTime))


if __name__ == '__main__':
    main()



"""
把它分解
前面的示例几乎与第1章中的顺序质因数分解相同，加快了速度!但是，您应该注意到，在main函数中，我们定义了10个不同的线程，而不是定义10个进程并加入它们。

如果您现在运行这个程序，您应该会看到我们的程序的整体性能大大降低，与单线程和多处理的程序相比。其结果如下:


正如您从上表的结果中看到的，通过启动多个线程并将它们抛出一个问题，与单线程解决方案相比，我们实际上已经实现了大约7%的减速，
并且与我们的多处理解决方案相比，几乎是100%的慢下来。
"""