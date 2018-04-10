import time

import random
from multiprocessing import Process


def calculatePrimeFactors(n):
    # 这是对给定数字n的所有质因数分解
    """质数因子"""
    primfac = []
    d = 2
    while d * d <= n:
        while (n % d) == 0:
            primfac.append(d)  # 假设你想重复多个因素。
            n //= d
        d += 1

        if n > 1:
            primfac.append(n)
        return primfac


# We split our workload from one batch of 10,000 calculations
# into 10 batches of 1,000 calculations
# 我们将工作负载从1万次计算中分离出来。
# 10个批次的1000次计算。
def exec_Proc():
    for i in range(1000):
        rand = random.randint(20000, 100000000)
        print(calculatePrimeFactors(rand))

def main():
    print("Starting number crunching")
    t0 = time.time()

    procs = []
    #  Here we create our processes and kick them off 在这里，我们创建processes并将其启动。
    for i in range(10):
        proc = Process(target=exec_Proc, args=())
        procs.append(proc)
        proc.start()

    # Again we use the .join() method in order to wait for
    # execution to finish for all of our processes
    # 再次，我们使用.join()  方法来等待。
    # 执行完成我们所有的processes。

    for proc in procs:
        proc.join()

    t1 = time.time()
    totalTime = t1 - t0
    # we print out the total execution time for our 10 procs.
    print("Execution Time: {}".format(totalTime))


if __name__ == '__main__':
    main()

"""
现在让我们来看看如何利用多个进程来提高这个程序的性能。

为了让我们将这个工作负载分解，我们将定义一个executeProc函数，它将生成1000个随机数，而不是生成10,000个随机数。我们将创建10个进程，并执行这个函数10次，所以计算的总数应该与我们执行顺序测试时的完全相同:

"""
