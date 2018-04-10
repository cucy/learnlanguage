import threading
import time


def myThread(myEvent):
    while not myEvent.is_set():
        print("Waiting for Event to be set")
        time.sleep(1)
    print("myEvent has been set")


def main():
    myEvent = threading.Event()
    thread1 = threading.Thread(target=myThread, args=(myEvent,))
    thread1.start()
    time.sleep(10)
    myEvent.set()


if __name__ == '__main__':
    main()


"""
把它分解
在前面的代码中，我们定义了一个myThread函数;在这个函数中，我们有一个while循环，
它只在我们传入该函数的事件对象未设置时运行。在这个循环中，我们简单地打印出我们正在等待事件被设置为1秒的间隔。

我们定义我们的事件对象，我们将传递给我们的主函数中的所有子线程。
为此，我们简单地调用myEvent = thread . event()，并为我们实例化一个事件对象的新实例。

然后我们实例化我们的线程对象，它接收我们的myEvent对象，然后启动它。
然后在设置myEvent信号之前，我们继续休眠10秒，这样我们的子线程就可以完成它的执行。

"""