import time
import random
import threading


def myThread(i):
    print("Thread {}: started".format(i))
    time.sleep(random.randint(1, 5))
    print("Thread {}: finished".format(i))


def main():
    for i in range(random.randint(2, 50)):
        thread = threading.Thread(target=myThread, args=(i,))
        thread.start()
    time.sleep(4)
    print("Total Number of Active Threads: {}".format(threading.active_count()))


if __name__ == '__main__':
    main()



"""
把它分解
在前面的例子中，我们所做的只是在2到50之间启动一个随机的线程数，然后让它们在随机的时间间隔内休眠。
一旦启动了所有给定的线程，我们就会休眠4秒，调用thread .active_count()，并将其输出到格式化的print语句中。

"""
