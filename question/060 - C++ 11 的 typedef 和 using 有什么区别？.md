<https://stackoverflow.com/questions/10747810/what-is-the-difference-between-typedef-and-using-in-c11>

## 问题

在 C++11 中下面的两条语句表达的都是一个意思，

```c++
typedef int MyInt;

using MyInt = int;
```

同时 `using` 还可在模板中使用，

```c++
emplate<class T>
using MyType = AnotherType<T, MyAllocatorType>;
```

那么 `typedef` 和 `using` 到底还有什么其它区别？

## 回答

没有了，除了 `using` 还可以在模板中使用，其它的都是等同的。
