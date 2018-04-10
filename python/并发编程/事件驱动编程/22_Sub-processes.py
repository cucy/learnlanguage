"""
We've seen in the past that single process programs sometimes cannot meet the demands required of them in order for our software to function properly. We looked at various mechanisms in the previous chapters on how we can improve performance using multiple processes, and thankfully, asyncio comes with the ability for us to still leverage the power of sub-processes within our event-driven based programs.

I am not a fan of using this mechanism for improving performance as it can drastically heighten the complexity of your programs. However, this isn't to say that there aren't situations where this would be useful, and as such, I should make you aware of the official documentation, which can be found at https://docs.python.org/3/library/asyncio-subprocess.html


过去我们已经看到，单进程程序有时不能满足它们的要求，以便我们的软件能够正常工作。在前面几章中，我们讨论了如何使用多个流程来提高性能的各种机制，所幸的是，asyncio有能力在我们的基于事件驱动的程序中利用子流程的力量。

我不喜欢使用这种机制来提高性能，因为它可以极大地提高程序的复杂性。但是，这并不是说没有这样的情况会有用，因此，我应该让您了解官方文档，这些文档可以在https://docs.python.org/3/library/asyncio-subprocess.html中找到。


"""