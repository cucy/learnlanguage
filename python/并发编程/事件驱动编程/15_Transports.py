"""
Transports
Transports are classes that come included within the asyncio module that allow you to implement various types of communication. In total, there are four distinct types of transports that each inherit from the BaseTransport class:

    ReadTransport
    WriteTransport
    DatagramTransport
    BaseSubprocesTransport
This BaseTransport class has five methods that are subsequently transient across all four transport types listed earlier:

    close(): This closes the transport
    is_closing(): This returns true if the transport is closing or is already closed
    get_extra_info(name, default=None): This returns optional transport information.
    set_protocol(protocol): This does exactly what it says on the tin
    get_protocol(): This returns the current protocol

传输
传输是包含在asyncio模块内的类，它允许您实现各种类型的通信。总共有四种不同类型的传输，它们分别从BaseTransport类继承:

ReadTransport
WriteTransport
DatagramTransport
BaseSubprocesTransport
这个BaseTransport类有五种方法，它们在前面列出的所有四种传输类型中都是短暂的:

close():这将关闭传输。
is_closed():如果传输关闭或已经关闭，则返回true。
get_extra_info(name, default=None):返回可选的传输信息。
set_protocol(协议):这正是它在tin上说的。
get_protocol():它返回当前协议。
"""