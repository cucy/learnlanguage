from __future__ import print_function, absolute_import

import grpc
from protos import helloworld_pb2
from protos import helloworld_pb2_grpc


def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(helloworld_pb2.HelloRequest(name="张三"))
        print("接收到服务端的响应: %s" % resp.message)

        resp_sayHelloAgain = stub.SayHelloAgain(helloworld_pb2.HelloRequest(name="李四"))
        print("接收到服务端的响应: %s" % resp_sayHelloAgain.message)


if __name__ == '__main__':
    run()
