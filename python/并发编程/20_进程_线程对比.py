import threading
from multiprocessing import Process
import time
import os


def MyTask():
    print("Starting")
    time.sleep(2)


t0 = time.time()
threads = []
for i in range(10):
    thread = threading.Thread(target=MyTask)
    thread.start()
    threads.append(thread)
t1 = time.time()
print("Total Time for Creating 10 Threads: {} seconds".format(t1 - t0))
for thread in threads:
    thread.join()
t2 = time.time()
procs = []
for i in range(10):
    process = Process(target=MyTask)
    process.start()
    procs.append(process)
t3 = time.time()
print("Total Time for Creating 10 Processes: {} seconds".format(t3 - t2))
for proc in procs:
    proc.join()

"""
Starting
Starting
Starting
Starting
Starting
Starting
Starting
Starting
Starting
Starting
Total Time for Creating 10 Threads: 0.0019259452819824219 seconds
Starting
Starting
Starting
Starting
Starting
Starting
Starting
Total Time for Creating 10 Processes: 0.047638654708862305 seconds
Starting
Starting
Starting




您将在前面的示例中看到，我们定义了一个MyTask函数，它将是我们将要创建的线程和进程的目标。

我们首先将启动时间存储在t0变量中，然后继续创建一个名为threads的空数组，它将方便地将引用存储到我们创建的所有线程对象中。
然后我们继续创建，然后在记录时间之前启动这些线程，这样我们就可以计算执行创建和启动所需的总时间。

然后我们继续按照与以前相同的创建和启动过程，但是这次，使用进程而不是线程。我们再次记录时间，并计算它们之间的差异。
在我的机器上运行这个脚本时，创建和启动的两个记录时间是一个数量级。创建和启动进程需要花费10倍的时间来创建和启动普通线程。
这个特定程序的输出在我的机器上是这样的:


现在，虽然为我们的相对轻量级的例子做这两项任务的时间可能是最小的，但是考虑一下您将看到的性能影响，如果您在巨大的服务器机架上启动了成百上千的进程或线程。

我们可以解决这个问题的一种方法是，在开始时做所有的进程或线程创建，并将它们存储在池中，这样它们就可以坐下来等待进一步的指令，而无需我们承担这些巨大的创建成本。在第7章中，我们将更深入地研究线程池和过程池的概念。
"""
