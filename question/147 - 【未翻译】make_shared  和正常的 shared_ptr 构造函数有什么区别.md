<https://stackoverflow.com/questions/20895648/difference-in-make-shared-and-normal-shared-ptr-in-c>

## 问题

```c++
std::shared_ptr<Object> p1 = std::make_shared<Object>("foo");
std::shared_ptr<Object> p2(new Object("foo"));
```

我看到很多人都推荐使用 make_shared，因为它比 shared_ptr 构造函数来的更高效，但我搞不懂是为什么？

有人能为我详细解释下么？

## 回答

它们的区别在于 make_shared 只有一次内存申请操作，而 shared_ptr 构造函数会有两次。

shared_ptr 对象会管理两部分内容，

- 控制块，比如引用计数、deleter 等等
- 要被管理的对象

当调用 make_shared 的时候，会申请一份足够大的内存同时给控制块和对象使用。而 shared_ptr 构造函数会分别为控制块和对象调用内存申请，详情可以参考 [cpprefrence - implementation notes](http://en.cppreference.com/w/cpp/memory/shared_ptr)。

当然 make_shared 也是有弊端的。当 shared_ptr 都离开了各自的作用域，被管理的对象也无法被析构。只有所有的 weak_ptr 也都离开了各自的作用域，这时候，一次申请的内存才会被释放掉。
