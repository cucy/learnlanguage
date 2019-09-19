"""Example gRPC client that gets/sets metadata (HTTP2 headers)"""

from __future__ import print_function
import logging

import grpc

import helloworld_pb2
import helloworld_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response, call = stub.SayHello.with_call(
            helloworld_pb2.HelloRequest(name='张三'),
            metadata=(
                # ("client客户端请求头", "c请求头值"),  # 不能用中文
                ('initial-metadata-1', 'The value should be str'),
                ('binary-metadata-bin', b'With -bin surffix, the value can be bytes'),
                ('accesstoken', 'gRPC Python is great'),
            ))

    print("客户端接收到: " + response.message)
    for key, value in call.trailing_metadata():
        print('客户端接收到添加的元数据metadata: key=%s value=%s' %
              (key, value))


if __name__ == '__main__':
    logging.basicConfig()
    run()
