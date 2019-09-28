from concurrent import futures

import time
import logging

import grpc
from protos import helloworld_pb2
from protos import helloworld_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Greeter(helloworld_pb2_grpc.GreeterServicer):
    #  实现方法
    def SayHello(self, request, context):
        return helloworld_pb2.HelloResponse(message="Hello,%s!" % request.name)

    # 实现另一个方法
    def SayHelloAgain(self, request, context):
        return helloworld_pb2.HelloResponse(message="SayHelloAgain,%s!" % request.name)


# 注册服务
def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port("127.0.0.1:50051")
    server.start()

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    server()
