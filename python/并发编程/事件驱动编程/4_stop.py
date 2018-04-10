"""
stop()方法再次按照它在tin上说的做。它将导致一个循环，
该循环正在通过run_forever()方法无限期地运行，以在下一个合适的机会中停止:
"""

import asyncio
import time


async def myWork():
    print("Starting Work")
    time.sleep(5)
    print("Ending Work")


def main():
    loop = asyncio.get_event_loop()
    loop.run_until_complete(myWork())
    loop.stop()
    print("Loop Stopped")
    loop.close()


if __name__ == '__main__':
    main()
