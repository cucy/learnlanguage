"""The Python implementation of the gRPC helloworld.Greeter client."""

from __future__ import print_function
import logging

import grpc

from protos import helloworld_pb2
from protos import helloworld_pb2_grpc
import default_value_client_interceptor


def run():
    default_value = helloworld_pb2.HelloResponse(message='Hello from your local interceptor, 本地拦截器!')
    default_value_interceptor = default_value_client_interceptor.DefaultValueClientInterceptor(default_value)
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051') as channel:
        intercept_channel = grpc.intercept_channel(channel, default_value_interceptor)
        stub = helloworld_pb2_grpc.GreeterStub(intercept_channel)
        response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
    print("接收到服务端返回数据: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()
