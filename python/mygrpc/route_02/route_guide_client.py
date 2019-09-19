import os
import random
import logging
import time

import grpc

from protos import routeguide_pb2
from protos import routeguide_pb2_grpc

import route_guide_resources

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


def make_route_note(message, latitude, longitude):
    return routeguide_pb2.RouteNote(
        message=message,
        location=routeguide_pb2.Point(latitude=latitude, longitude=longitude))


def guide_get_one_feature(stub, point):
    feature = stub.GetFeature(point)

    if not feature.location:
        print("Server returned incomplete feature")
        return

    if feature.name:
        print("查找: %s \n经纬度:\n%s" % (feature.name, feature.location))
    else:
        print("没有匹配到 %s" % feature.location)

    # #  对getFeature的异步调用与此类似，但与在线程池中异步调用本地方法类似：
    # feature_future = stub.GetFeature.future(point)
    # feature = feature_future.result()


# 一问一答
def guide_get_feature(stub):
    """
    此RPC方法是同步调用,几乎和调用本地方法(函数)一样, client发起调用后一直等到服务端响应, 服务端一般是返回正常响应或者是异常响应
    :param stub:
    :return:
    """
    print("--------------最简单的rpc  GetFeature --------------")
    guide_get_one_feature(stub, routeguide_pb2.Point(latitude=409146138, longitude=-746188906))
    guide_get_one_feature(stub, routeguide_pb2.Point(latitude=0, longitude=0))


# 服务端流式处理
def guide_list_features(stub):
    # 服务端流式处理
    print("--------------服务端流式  ListFeatures --------------")
    rectangle = routeguide_pb2.Rectangle(
        lo=routeguide_pb2.Point(latitude=400000000, longitude=-750000000),
        hi=routeguide_pb2.Point(latitude=420000000, longitude=-730000000))
    print("Looking for features between 40, -75 and 42, -73")
    #  调用响应流列表功能类似于处理序列类型：
    features = stub.ListFeatures(rectangle)

    for feature in features:
        print("Feature called %s at %s" % (feature.name, feature.location))


def generate_route(feature_list):
    for _ in range(0, 10):
        random_feature = feature_list[random.randint(0, len(feature_list) - 1)]
        print("访问点 Visiting point %s" % random_feature.location)
        yield random_feature.location


#  客户端流式
def guide_record_route(stub):
    #  客户端流式
    #  从json文件中读取列表
    print("--------------客户端流式 RecordRoute --------------")
    #  从本地加载数据
    feature_list = route_guide_resources.read_route_guide_database()
    #  Calling the request-streaming RecordRoute is similar to passing an iterator to a local method. Like the simple RPC above that also returns a single response, it can be called synchronously or asynchronously:
    # 调用请求流记录路由类似于将迭代器传递给本地方法。与上面也返回单个响应的简单rpc一样，可以同步或异步调用它：

    route_iterator = generate_route(feature_list)
    route_summary = stub.RecordRoute(route_iterator)  # 同步
    print("完成%s个传输" % route_summary.point_count)
    print("传递了%s个功能" % route_summary.feature_count)
    print("行驶了%s米" % route_summary.distance)
    print("总共花费%s秒" % route_summary.elapsed_time)

    # 同步
    ''' 
    route_summary = stub.RecordRoute(point_iterator)
    '''
    # 异步模式
    '''
    route_summary_future = stub.RecordRoute.future(point_iterator)
    route_summary = route_summary_future.result()
    '''


def generate_messages():
    messages = [
        make_route_note("第一条消息", 0, 0),
        make_route_note("第二条消息", 0, 1),
        make_route_note("第三条消息", 1, 0),
        make_route_note("第四条消息", 0, 0),
        make_route_note("第五条消息", 1, 0),
    ]
    for msg in messages:
        print("发送 %s at %s" % (msg.message, msg.location))
        yield msg


# 客户端,服务端都是流式
def guide_route_chat(stub):
    """
    双向流式rpc
    调用双向流路由具有请求流和响应流语义的组合（与服务端的情况一样）
    for received_route_note in stub.RouteChat(sent_route_note_iterator):
    """
    print("--------------请求,响应都是流式  RouteChat --------------")
    responses = stub.RouteChat(generate_messages())
    for response in responses:
        print("接收到消息 %s at %s" % (response.message, response.location))


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('192.168.1.183:50051') as channel:
        stub = routeguide_pb2_grpc.RouteGuideStub(channel)

        guide_get_feature(stub)
        os._exit(1)
        guide_list_features(stub)
        guide_record_route(stub)
        guide_route_chat(stub)


if __name__ == '__main__':
    logging.basicConfig()
    run()
    time.sleep(2)
