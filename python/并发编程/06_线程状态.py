import threading, time


# A very simple method for our thread to execute
def threadWorker():
    # it is only at the point where the thread starts executing    只有在线程开始执行的时候。
    # that it's state goes from 'Runnable' to a 'Running'   它的状态从“Runnable”到“Running”
    # state
    print("My Thread has entered the 'Running' State")

    # If we call the time.sleep() method then our thread  如果我们调用time.sleep()方法，那么我们的线程。
    # goes into a not-runnable state. We can do no further work  进入一个不能运行的状态。我们不能再做进一步的工作了。
    # on this particular thread          在这个特定的线程上。
    time.sleep(10)

    # Thread then completes its tasks and terminates 线程然后完成任务并终止。
    print("My Thread is terminating终止")


# At this point in time, the thread has no state  此时，线程没有状态。
#  it hasn't been allocated any system resources   它还没有分配任何系统资源。
myThread = threading.Thread(target=threadWorker)

# When we call myThread.start(), Python allocates the necessary system  当我们调用myThread.start()时，Python会分配必要的系统。
# resources in order for our thread to run and then calls the thread's   为我们的线程运行并调用线程的资源。
# run method. It goes from 'Starting' state to 'Runnable' but not running 运行方法。它从“开始”状态变为“可运行”但不运行。
myThread.start()

# Here we join the thread and when this method is called    这里我们加入线程，当调用此方法时。
# our thread goes into a 'Dead' state. It has finished the    我们的线程进入了“死”状态。它已经完成了
# job that it was intended to do.     它本来就是要做的工作。
myThread.join()
print("My Thead has entered a 'Dead' state")



"""
把它分解
在前面的代码示例中，我们定义了一个函数threadWorker，它将是我们将要创建的线程的调用目标。这个threadWorker函数所做的一切就是打印出它的当前状态，然后通过调用time.sleep(10)来休眠10秒。在定义了threadWorker之后，我们接着在这一行中创建一个新的线程对象:

myThread = threading.Thread(target=threadWorker)

"""

"""
此时，我们的线程对象当前处于新的线程状态，并且还没有分配它需要运行的任何系统资源。这只发生在我们调用这个函数时

myThread.start()

"""

"""
此时，我们的线程分配了所有的资源，并且调用了线程的运行函数。线程现在进入“可运行”状态。它接着打印出自己的状态，然后通过调用time.sleep(10)来阻塞10秒。在该线程休眠的10秒内，线程被认为处于“不运行”状态，其他线程将被调度在此线程上运行。

最后，在10秒的周期结束后，我们的线程被认为已经结束并且处于“死”状态。它不再需要它所分配的任何资源，它将被垃圾收集器清理。
"""
