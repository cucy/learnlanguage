# 概述

- Service definition (服务定义)

类其他RPC一样, grpc基于定义服务的思想,需指定可使用其参数和返回类型远程调用的方法(指定方法名, 调用参数(参数类型),返回数据(数据类型))。
grpc默认使用[protocol buffers](https://developers.google.com/protocol-buffers/) 作为其接口定义语言(IDL) (类,方法)来描述有效负载(payload message)既可使用其参数和返回类型远程调用的方法(指定方法名, 调用参数(参数类型),返回数据(数据类型))
如果需要，可以使用其他替代方案。

```proto

// 定义接口名(类名)
service HelloService {
  // 定义方法名, 调用参数, 返回参数
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

// 定义数据类型
message HelloRequest {
  string greeting = 1;
}
// 定义数据类型
message HelloResponse {
  string reply = 1;
}
```


> grpc 可以定义四种不同方法

* 传统的请求响应类型 (一问一答: request,response)

一元rpc，客户端向服务器发送一个请求并得到一个响应，就像普通的函数调用一样。

```proto
rpc SayHello(HelloRequest) returns (HelloResponse){
}
```

* Server streaming RPCs  (客户端发送一次请求,  服务端可以流式的响应信息,且保证响应是有序的)

服务器流rpc，客户端向服务器发送请求，并获取流以读取消息序列。客户端从返回的流中读取，直到没有更多的消息为止。grpc保证在单个rpc调用中进行消息排序。

```proto
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
}
```


* Client streaming RPCs 

客户端流式rpc，其中客户端再次使用提供的流写入一系列消息并将它们发送到服务器。一旦客户端完成了消息的写入，它就会等待服务器读取它们并返回其响应。grpc再次保证了单个rpc调用中的消息排序。

```proto
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
}
```


*  Bidirectional streaming RPCs

双向流rpc，其中双方使用读写流发送一系列消息。这两个流独立运行，因此客户机和服务器可以按照它们喜欢的顺序读写.
例如，服务器可以在写入响应之前等待接收所有客户机消息，或者它可以交替地读一条消息，然后再写一条消息，或是其他一些读和写。保留每个流中消息的顺序。

```proto
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
}
```



## Using the API surface(表面) 


Starting from a service definition in a .proto file, gRPC provides protocol buffer compiler plugins that generate client- and server-side code. gRPC users typically call these APIs on the client side and implement the corresponding API on the server side.
GRPC从.proto文件中的服务定义开始，提供生成客户端和服务器端代码的协议缓冲区编译器插件。grpc用户通常在客户端调用这些api，并在服务器端实现相应的api。 (gRPC通过.proto定义服务,提供生成工具生成对应不同语言的代码, client调用proto定义的方法,服务端实现proto定义的方法)

On the server side, the server implements the methods declared by the service and runs a gRPC server to handle client calls. The gRPC infrastructure decodes incoming requests, executes service methods, and encodes service responses.
在服务器端，服务器实现服务声明的方法，并运行grpc服务器来处理客户端调用。GRPC基础设施对传入的请求进行解码，执行服务方法，并对服务响应进行编码。 (接收请求 -> 解码 -> 执行方法 -> 编码 -> 发送响应请求)


On the client side, the client has a local object known as stub (for some languages, the preferred term is client) that implements the same methods as the service. The client can then just call those methods on the local object, wrapping the parameters for the call in the appropriate protocol buffer message type - gRPC looks after sending the request(s) to the server and returning the server’s protocol buffer response(s).
在客户端，客户端有一个名为stub的本地对象（对于某些语言，首选术语是client），它实现与服务相同的方法。然后，客户机可以在本地对象上调用这些方法，将调用的参数包装在适当的协议缓冲区消息类型中—GRPC在将请求发送到服务器并返回服务器的协议缓冲区响应后进行处理。
(客户端封装方法, 将参数封装成proto文件所定义的格式发送请求, 服务端进行处理响应后, 客户端再做进一步的操作)


## Synchronous vs. asynchronous 同步 异步

Synchronous RPC calls that block until a response arrives from the server are the closest approximation to the abstraction of a procedure call that RPC aspires to. On the other hand, networks are inherently asynchronous and in many scenarios it’s useful to be able to start RPCs without blocking the current thread.

The gRPC programming surface in most languages comes in both synchronous and asynchronous flavors. You can find out more in each language’s tutorial and reference documentation (complete reference docs are coming soon).

同步的rpc调用在响应到达服务器之前一直阻塞，这与rpc所期望的过程调用的抽象最为接近。另一方面，网络本质上是异步的，在许多情况下，能够在不阻塞当前线程的情况下启动rpc是很有用的。
大多数语言中的grpc编程界面都有同步和异步两种风格。您可以在每种语言的教程和参考文档中找到更多信息（即将提供完整的参考文档）。


## RPC life cycle

Now let’s take a closer look at what happens when a gRPC client calls a gRPC server method. We won’t look at implementation details, you can find out more about these in our language-specific pages.
rpc生命周期
现在让我们更仔细地看一下当grpc客户机调用grpc服务器方法时会发生什么。我们不会查看实现的详细信息，您可以在我们的特定语言页面中找到更多关于这些的信息。

## Unary RPC
First let’s look at the simplest type of RPC, where the client sends a single request and gets back a single response.

Once the client calls the method on the stub/client object, the server is notified that the RPC has been invoked with the client’s metadata for this call, the method name, and the specified deadline if applicable.
The server can then either send back its own initial metadata (which must be sent before any response) straight away, or wait for the client’s request message - which happens first is application-specific.
Once the server has the client’s request message, it does whatever work is necessary to create and populate its response. The response is then returned (if successful) to the client together with status details (status code and optional status message) and optional trailing metadata.
If the status is OK, the client then gets the response, which completes the call on the client side.

一元RPC
首先，让我们看看最简单的rpc类型，其中客户机发送一个请求并返回一个响应。
一旦客户机调用存根/客户机对象上的方法，就会通知服务器已使用此调用的客户机元数据、方法名和指定的截止日期（如果适用）调用了rpc。
然后，服务器可以直接发送回它自己的初始元数据（必须在任何响应之前发送），或者等待客户机的请求消息（首先发生的是特定于应用程序的消息）。
一旦服务器收到客户机的请求消息，它就会执行创建和填充其响应所需的任何工作。然后将响应（如果成功）连同状态详细信息（状态代码和可选状态消息）和可选的尾部元数据返回给客户端。
如果状态为ok，则客户端将获得响应，从而在客户端完成调用。

## Server streaming RPC
A server-streaming RPC is similar to our simple example, except the server sends back a stream of responses after getting the client’s request message. After sending back all its responses, the server’s status details (status code and optional status message) and optional trailing metadata are sent back to complete on the server side. The client completes once it has all the server’s responses.

服务器流式RPC
服务器流式rpc与我们的简单示例类似，只是服务器在获取客户机的请求消息后发送回响应流。在发送回其所有响应之后，服务器的状态详细信息（状态代码和可选状态消息）和可选的尾部元数据将发送回服务器端以完成。一旦客户端拥有服务器的所有响应，它就完成了。

## Client streaming RPC
A client-streaming RPC is also similar to our simple example, except the client sends a stream of requests to the server instead of a single request. The server sends back a single response, typically but not necessarily after it has received all the client’s requests, along with its status details and optional trailing metadata.


客户端流式RPC
客户端流式rpc也类似于我们的简单示例，只是客户端向服务器发送请求流，而不是单个请求。服务器发送回一个响应，通常但不一定是在接收到客户端的所有请求及其状态详细信息和可选的尾部元数据之后。

## Bidirectional streaming RPC
In a bidirectional streaming RPC, again the call is initiated by the client calling the method and the server receiving the client metadata, method name, and deadline. Again the server can choose to send back its initial metadata or wait for the client to start sending requests.

What happens next depends on the application, as the client and server can read and write in any order - the streams operate completely independently. So, for example, the server could wait until it has received all the client’s messages before writing its responses, or the server and client could “ping-pong”: the server gets a request, then sends back a response, then the client sends another request based on the response, and so on.

双向流式rpc
在双向流式rpc中，调用再次由调用该方法的客户端和接收客户端元数据、方法名称和截止日期的服务器发起。同样，服务器可以选择发送回其初始元数据或等待客户端开始发送请求。
接下来会发生什么取决于应用程序，因为客户端和服务器可以按任何顺序读写-流完全独立地运行。因此，例如，服务器可以等到接收到客户端的所有消息后再写入其响应，或者服务器和客户端可以“乒乓”：服务器获取请求，然后发送回响应，然后客户端根据响应发送另一个请求，依此类推。


## Deadlines/Timeouts
gRPC allows clients to specify how long they are willing to wait for an RPC to complete before the RPC is terminated with the error DEADLINE_EXCEEDED. On the server side, the server can query to see if a particular RPC has timed out, or how much time is left to complete the RPC.

How the deadline or timeout is specified varies from language to language - for example, not all languages have a default deadline, some language APIs work in terms of a deadline (a fixed point in time), and some language APIs work in terms of timeouts (durations of time).

截止日期/超时
GRPC允许客户端指定在超过错误截止日期的情况下终止RPC之前，他们愿意等待RPC完成的时间。在服务器端，服务器可以查询特定的rpc是否超时，或者还有多少时间来完成rpc。
指定截止时间或超时的方式因语言而异—例如，并非所有语言都有默认的截止时间，有些语言API按截止时间（固定时间点）工作，有些语言API按超时时间（时间段）工作。


## RPC termination
In gRPC, both the client and server make independent and local determinations of the success of the call, and their conclusions may not match. This means that, 
for example, you could have an RPC that finishes successfully on the server side (“I have sent all my responses!”) but fails on the client side (“The responses arrived after my deadline!“). It’s also possible for a server to decide to complete before a client has sent all its requests.

rpc终止
在grpc中，客户机和服务器都对呼叫的成功进行独立和本地的判断，它们的结论可能不匹配。这意味着，
例如，您可以有一个在服务器端成功完成的rpc（“我已经发送了所有响应！）但是在客户端失败了（“回复是在我的截止日期之后到达的！“”。服务器也可以在客户端发送其所有请求之前决定完成。


## Cancelling RPCs
Either the client or the server can cancel an RPC at any time. A cancellation terminates the RPC immediately so that no further work is done. It is not an “undo”: changes made before the cancellation will not be rolled back.

取消RPC
客户端或服务器可以随时取消rpc。取消操作会立即终止rpc，这样就不会做进一步的工作。这不是“撤消”：取消前所做的更改不会回滚。

## Metadata
Metadata is information about a particular RPC call (such as authentication details) in the form of a list of key-value pairs, where the keys are strings and the values are typically strings (but can be binary data). Metadata is opaque to gRPC itself - it lets the client provide information associated with the call to the server and vice versa.

Access to metadata is language-dependent.

元数据
元数据是有关特定rpc调用（如身份验证详细信息）的信息，以键-值对列表的形式显示，其中键是字符串，值通常是字符串（但可以是二进制数据）。元数据对grpc本身是不透明的，它允许客户端提供与服务器调用相关联的信息，反之亦然。
对元数据的访问依赖于语言。

## Channels
A gRPC channel provides a connection to a gRPC server on a specified host and port and is used when creating a client stub (or just “client” in some languages). Clients can specify channel arguments to modify gRPC’s default behaviour, such as switching on and off message compression. A channel has state, including connected and idle.

How gRPC deals with closing down channels is language-dependent. Some languages also permit querying channel state.

渠道
GRPC通道提供到指定主机和端口上的GRPC服务器的连接，并在创建客户端存根（或某些语言中的“客户端”）时使用。客户端可以指定通道参数来修改grpc的默认行为，例如打开和关闭消息压缩。通道具有状态，包括已连接和空闲。
GRPC如何关闭频道取决于语言。有些语言还允许查询通道状态。


---

# Authentication

本文档概述了GRPC身份验证，包括我们内置的受支持的身份验证机制、如何插入您自己的身份验证系统，以及如何在受支持的语言中使用GRPC身份验证的示例。


概述
GRPC设计用于各种身份验证机制，使得安全地使用GRPC与其他系统进行通信变得容易。您可以使用我们支持的机制-有或没有基于google令牌的身份验证的ssl/tls-或者通过扩展我们提供的代码插入您自己的身份验证系统。
grpc还提供了一个简单的身份验证api，允许您在创建通道或进行调用时提供所有必要的身份验证信息作为凭据。


支持的身份验证机制
GRPC内置以下身份验证机制：
ssl/tls:grpc集成了ssl/tls，并提倡使用ssl/tls对服务器进行身份验证，并加密客户端和服务器之间交换的所有数据。可选的机制可供客户端提供用于相互身份验证的证书。
使用google基于令牌的身份验证：grpc提供了一种通用机制（如下所述），用于将基于元数据的凭据附加到请求和响应。通过grpc访问googleapi时获取访问令牌（通常是oauth2令牌）的额外支持是为某些身份验证流提供的：您可以在下面的代码示例中看到这是如何工作的。一般来说，这个机制必须与通道上的ssl/tls一起使用-google不允许没有ssl/tls的连接，而且大多数grpc语言实现不允许您在未加密的通道上发送凭据。
警告：google凭据只能用于连接到google服务。向非google服务发送google颁发的oauth2令牌可能会导致此令牌被盗，并用于将客户端模拟为google服务。
身份验证API
grpc提供了一个简单的认证api，它基于credentials对象的统一概念，可以在创建整个grpc通道或单个调用时使用。

凭据类型
凭据可以是两种类型：
附加到通道的通道凭据，如SSL凭据。
调用凭据，附加到一个调用（或C++中的ClientContext）。
您还可以在compositechannelcredentials中组合这些内容，例如，允许您指定通道的ssl详细信息以及通道上每个调用的调用凭据。compositechannelcredentials将channelcredentials和callcredentials关联起来以创建新的channelcredentials。结果将发送与组合的callcredentials相关联的身份验证数据以及在通道上进行的每个调用。
例如，可以从sslcredentials和accesstokencredentials创建channelcredentials。应用于通道时的结果将为此通道上的每个调用发送适当的访问令牌。
也可以使用compositeCallCredentials组合单个CallCredentials。在调用中使用生成的CallCredentials将触发与两个CallCredentials关联的身份验证数据的发送。




---



# 环境安装


pip install grpcio
pip install grpcio-tools

## 基础

### 定义proto格式IDL(数据交互语言)

```proto
syntax = "proto3";

option java_multiple_files = true;
option java_package = "io.grpc.examples.helloworld";
option java_outer_classname = "HelloWorldProto";
option objc_class_prefix = "HLW";

// 定义包名(模块名)
package helloworld;

// 定义发起请求的数据格式
message HelloRequest {
    string name = 1;
}
message HelloResponse {
    string message = 1;
}
// 定义类名(接口)
service Greeter {
    // 定义方法
    rpc SayHello (HelloRequest) returns (HelloResponse) {
    }
}
```


### `生成相对语言代码`

python -m grpc_tools.protoc -I./protos --python_out=./protos --grpc_python_out=./protos  ./protos/helloworld.proto


### 服务端实现

```python
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


# 注册服务
def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port(":50051")
    server.start()

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    server()

```


### 客户端实现

```python
from __future__ import print_function, absolute_import

import grpc
from protos import helloworld_pb2
from protos import helloworld_pb2_grpc


def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(helloworld_pb2.HelloRequest(name="张三"))
        print("接收到服务端的响应: %s" % resp.message)


if __name__ == '__main__':
    run()

```