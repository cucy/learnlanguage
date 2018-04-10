import random
import threading

import time


class Philosopher(threading.Thread):

    def __init__(self, leftFork, rightFork):
        print("Our Philosopher Has Sat Down At the Table")
        threading.Thread.__init__(self)
        self.leftFork = leftFork
        self.rightFork = rightFork

    def run(self):
        print("Philosopher: {} has started thinking".format(threading.current_thread()))
        while True:
            time.sleep(random.randint(1, 5))
            print("Philosopher {} has finished thinking".format(threading.current_thread()))
            self.leftFork.acquire()
            time.sleep(random.randint(1, 5))
            try:
                print("Philosopher {} has acquired the left fork".format(threading.current_thread()))
                self.rightFork.acquire()
                try:
                    print("Philosopher {} has attained both forks, currentlyeating".format(threading.current_thread()))
                finally:
                    self.rightFork.release()
                    print("Philosopher {} has released the right fork".format(threading.current_thread()))
            finally:
                self.leftFork.release()
                print("Philosopher {} has released the left fork".format(threading.current_thread()))


t = Philosopher(1,2)
t.run()