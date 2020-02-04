<https://stackoverflow.com/questions/6500313/why-should-c-programmers-minimize-use-of-new>

## 问题

我看到一个问题 - [Memory leak with std::string when using std::list< std::string >](https://stackoverflow.com/q/3428750/211563)，其中的一个 [回答](https://stackoverflow.com/q/3428750/211563#comment3570156_3428750) 表述 C++ 程序员应尽量避免使用 `new`。

我不太明白为什么他那么说？

## 回答

C++ 并不带自动 GC。任何的 `new` 都需要有对应的 `delete`，否则就会有内存泄漏。

```c++
std::string *someString = new std::string(...);

// Do stuff

delete someString;
```

上面的语句看着没什么问题，但有个陷阱 - 如果 `Do stuff` 里发生异常了呢？`delete` 就不会被调用。

我们应该尽量这么用，

```c++
std::string someString(...);

// Do stuff
```

这就是 [RAII](http://en.wikipedia.org/wiki/Resource_Acquisition_Is_Initialization) 技术。当离开它的作用域的时候，`someString` 便会自动析构。

而且 C++11 完善了智能指针，旨在可以更方便地帮助我们实现 RAII，我们可以适当地加以利用。
