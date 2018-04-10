import threading
import time


def standardThread():
    print("Starting my Standard Thread启动标准线程")
    time.sleep(10)
    print("Ending my standard thread 结束标准线程")


def daemonThread():
    while True:
        print("Sending Out Heartbeat Signal发送心跳信号")
        time.sleep(2)


if __name__ == '__main__':
    standardThread = threading.Thread(target=standardThread)
    daemonThread = threading.Thread(target=daemonThread)
    daemonThread.setDaemon(True)
    daemonThread.start()

    standardThread.start()



"""
把它分解
在前面的代码示例中，我们定义了两个函数，它们将作为我们正常的、非守护线程和daemonThread的目标。我们的标准线程函数本质上就是输出它的状态，休眠20秒，以模拟一个较长的程序。

daemonThread函数进入一个永久的while循环，并简单地打印出每2秒发出的心跳信号。这只是您选择的任何心跳机制的一个占位符。

在我们的main函数中，我们创建了两个线程，我们的标准线程和守护线程，我们使用相同的start()方法开始。您将注意到，我们还在daemonThread对象上使用了setDaemon函数。这只是将线程对象的守护标志设置为我们传入该函数的任何内容，并且仅用于引用。

"""