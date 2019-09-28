# gRPC Python Example for Wait-for-ready

The default behavior of an RPC is to fail instantly if the server is not ready yet. This example demonstrates how to change that behavior.

### Definition of 'wait-for-ready' semantics

> If an RPC is issued but the channel is in TRANSIENT_FAILURE or SHUTDOWN states, the RPC is unable to be transmitted promptly. By default, gRPC implementations SHOULD fail such RPCs immediately. This is known as "fail fast," but the usage of the term is historical. RPCs SHOULD NOT fail as a result of the channel being in other states (CONNECTING, READY, or IDLE).
>
> gRPC implementations MAY provide a per-RPC option to not fail RPCs as a result of the channel being in TRANSIENT_FAILURE state. Instead, the implementation queues the RPCs until the channel is READY. This is known as "wait for ready." The RPCs SHOULD still fail before READY if there are unrelated reasons, such as the channel is SHUTDOWN or the RPC's deadline is reached.
>
> From https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md 

### Use cases for 'wait-for-ready'

When developers spin up gRPC clients and servers at the same time, it is very like to fail first couple RPC calls due to unavailability of the server. If developers failed to prepare for this situation, the result can be catastrophic. But with 'wait-for-ready' semantics, developers can initialize the client and server in any order, especially useful in testing.

Also, developers may ensure the server is up before starting client. But in some cases like transient network failure may result in a temporary unavailability of the server. With 'wait-for-ready' semantics, those RPC calls will automatically wait until the server is ready to accept incoming requests.

### DEMO Snippets

```Python
# Per RPC level
stub = ...Stub(...)

stub.important_transaction_1(..., wait_for_ready=True)
stub.unimportant_transaction_2(...)
stub.important_transaction_3(..., wait_for_ready=True)
stub.unimportant_transaction_4(...)
# The unimportant transactions can be status report, or health check, etc.
```



#等待就绪的grpc python示例
如果服务器还没有准备好，rpc的默认行为是立即失败。这个例子演示了如何改变这种行为。
###“等待就绪”语义的定义
>如果发出了rpc，但通道处于瞬时故障或关闭状态，则无法立即传输rpc。默认情况下，grpc实现应立即使此类rpc失败。这被称为“快速失败”，但这个词的使用是历史性的。rpc不应由于通道处于其他状态（连接、就绪或空闲）而失败。
>
>GRPC实现可以提供每个RPC选项，以避免由于通道处于瞬时故障状态而导致RPC失败。相反，实现对rpc进行排队，直到通道准备就绪。这就是所谓的“等待就绪”。如果有不相关的原因，例如通道关闭或达到了rpc的截止日期，rpc在就绪之前应该仍然失败。
>
>来自https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md
###“等待就绪”的用例
当开发人员同时启动grpc客户机和服务器时，很可能会因为服务器不可用而导致前两个rpc调用失败。如果开发人员未能为这种情况做好准备，结果可能是灾难性的。但是使用“等待就绪”语义，开发人员可以按任何顺序初始化客户机和服务器，特别是在测试中非常有用。
另外，开发人员可以在启动客户机之前确保服务器已启动。但在某些情况下，如短暂的网络故障，可能会导致服务器暂时不可用。使用“wait for ready”语义，这些rpc调用将自动等待服务器准备好接受传入请求。
###演示片段