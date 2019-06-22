<https://stackoverflow.com/questions/1452721/why-is-using-namespace-std-considered-bad-practice>

## 问题

有人告诉我在代码里直接使用`using namespace std;`这样很不好，应该这么用，`std::cout`、`std::cin`等等。

但是为什么不好呢？

影响性能？命名冲突？

#### 最佳回答

首先这跟性能是没有关系的。

现在考虑你正在使用两个库，分别是`Foo`和`Bar`，

```c++
using namespace foo;
using namespace bar;
```

不管是调用`foo`里的函数`Blah()`，还是调用`bar`里的`Quux()`，都没啥问题。

然后有一天你的库`Foo`要升级了，里边新加了一个函数`Quux()`，这样就出现问题了，因为它和命名空间`bar`里的`Quux()`冲突了。想一想，如果很多函数名都冲突了，
你是不是得一个一个去解决，费时间费力气。

所以如果你当初没有全局导入这些名称，这些问题就都不存在了。

因此，不建议全局导入命名空间，而是你要用到哪个就显示指定哪个命名空间，这样的代码本身阅读性也更好。
