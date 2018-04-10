import threading
import urllib.request
import time


def downloadImage(imagePath, fileName):
    """
        抽象出生成文件名函数, 并将其应用到executeThread函数中
    """
    print(f"从下载{imagePath}图片")
    urllib.request.urlretrieve(imagePath, fileName)
    print(f"{fileName}下载完成")


def executeThread(i):
    imageName = f"img/image_{i}.jpg"
    downloadImage("http://lorempixel.com/400/200/sports", imageName)


def main():
    t0 = time.time()

    # create an array which will store a reference to
    #  all of our threads
    threads = []
    # create 10 threads, append them to our array of threads   创建10个线程，将它们附加到我们的线程数组中。
    # and start them off开始
    for i in range(10):
        thread = threading.Thread(target=executeThread, args=(i,))
        threads.append(thread)
        thread.start()

    # ensure that all the threads in our array have completed    确保数组中的所有线程都已完成。
    # their execution before we log the total time to complete   在我们记录总时间完成之前，它们的执行。

    for i in threads:
        i.join()

    #  calculate the total execution time
    t1 = time.time()
    total_time = t1 - t0
    print(f'总执行时间: {total_time}')


if __name__ == '__main__':
    main()
