from multiprocessing import Pool
import time

def myTask(n):
    time.sleep(n/2)
    return n*2

def main():
    with Pool(4) as p:
        print(p.apply(myTask, (4,)))
        print(p.apply(myTask, (3,)))
        print(p.apply(myTask, (2,)))
        print(p.apply(myTask, (1,)))

if __name__ == '__main__':
    main()