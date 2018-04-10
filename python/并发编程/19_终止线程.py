from multiprocessing import Process
import time


def myWorker():
    t1 = time.time()
    print("Process started at: {}".format(t1))
    time.sleep(20)


myProcess = Process(target=myWorker)
print("Process {}".format(myProcess))
myProcess.start()
print("Terminating Process...")
myProcess.terminate()
myProcess.join()
print("Process Terminated: {}".format(myProcess))



"""
在前面的示例中，我们定义了一个简单的myWorker()函数，它打印出启动时间，然后休眠20秒。
然后我们继续声明myProcess，这是一个类型过程，我们将myWorker函数作为其执行的目标。

我们启动进程，然后使用终止方法立即终止它。您应该注意到，在输出中，这个程序几乎是立即完成的，而myProcess进程不会阻塞整个20秒。

"""
