from multiprocessing import Pool
import time

def myTask(n):
    time.sleep(n/2)
    return n*2

def main():
    print("mapping array to pool")
    with Pool(4) as p:
        print(p.map(myTask, [4,3,2,1]))

if __name__ == '__main__':
    main()



"""
map_async


from multiprocessing import Pool
import time

def myTask(n):
    time.sleep(n/2)
    return n*2

def main():
    with Pool(4) as p:
        print(p.map_async(myTask, [4,3,2,1]).get())

if __name__ == '__main__':
    main()


"""