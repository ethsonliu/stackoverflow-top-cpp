<https://stackoverflow.com/questions/3279543/what-is-the-copy-and-swap-idiom>

## 问题

我发现 copy-and-swap 这个名词在很多地方都出现，

- [What are your favorite C++ Coding Style idioms: Copy-swap](https://stackoverflow.com/questions/276173/what-are-your-favorite-c-coding-style-idioms/2034447#2034447)
- [Copy constructor and = operator overload in C++: is a common function possible?](https://stackoverflow.com/questions/1734628/copy-constructor-and-operator-overload-in-c-is-a-common-function-possible/1734640#1734640)
- [What is copy elision and how it optimizes copy-and-swap idiom](https://stackoverflow.com/questions/2143787/what-is-copy-elision-and-how-it-optimizes-copy-and-swap-idiom)
- [C++: dynamically allocating an array of objects?](https://stackoverflow.com/questions/255612/c-dynamically-allocating-an-array-of-objects/255744#255744)

它到底是什么意思？怎么用？在 C++ 11 中它又有什么变化？

## 回答

为什么需要 copy-and-swap 呢？ 任何资源管理类（比如智能指针）都需要遵循一个规则：[三法则](https://github.com/Hapoa/stackoverflow-top-cpp/blob/master/question/014%20-%20%E4%BB%80%E4%B9%88%E6%98%AF%E2%80%9C%E4%B8%89%E6%B3%95%E5%88%99%E2%80%9D%EF%BC%9F.md)。其中拷贝构造函数和析构函数实现起来比较容易，但是拷贝赋值操作要复杂许多，而 copy-and-swap 就是实现拷贝赋值操作完美的解决方案。它既能避免代码冗余，还可以提供 [强异常安全保证](https://en.wikipedia.org/wiki/Exception_safety)。

那 copy-and-swap 是怎么实现的呢？大致思路是：先用拷贝构造函数创建一个副本，然后利用函数`swap`交换其成员数据，当作用域退出，副本的析构函数会自动调用。这里有三个注意点：一，拷贝构造函数应该是可用的；二，这里的`swap`并非指`std::swap`，而是需要我们自己写的，而且需要保证`swap`不会抛出异常；三：析构函数也应该是可用的。


