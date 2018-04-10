import threading, time, random


def executeThread(i):
    print(f'thread  {i} starded')
    sleep_time = random.randint(1, 10)
    time.sleep(sleep_time)
    print(f'thread {i} 执行结束')


for i in range(10):
    thread = threading.Thread(target=executeThread, args=(i,))
    thread.start()
    print("active threads", threading.enumerate())



"""
把它分解
在前面的代码中，我们定义了一个名为executeThread的简单函数，它将i作为它的惟一参数。在这个函数中，我们只需调用time.sleep()函数，并在1到10之间随机生成一个整数。

然后，我们继续声明一个for循环，该循环从1到10循环，它创建一个线程对象，然后在传入线程的args时启动它。当您运行这个脚本时，您应该看到如下内容:
thread  0 starded
active threads [<_MainThread(MainThread, started 9028)>, <Thread(Thread-1, started 11164)>]
thread  1 starded
active threads [<_MainThread(MainThread, started 9028)>, <Thread(Thread-1, started 11164)>, <Thread(Thread-2, started 10928)>]
thread  2 starded
"""