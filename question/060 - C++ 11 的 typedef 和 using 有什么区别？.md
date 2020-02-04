<https://stackoverflow.com/questions/10747810/what-is-the-difference-between-typedef-and-using-in-c11>

## 问题

在 C++11 中下面的两条语句是等同的，

```c++
typedef int MyInt;

using MyInt = int;
```

同时 `using` 还可在模板中使用，

```c++
emplate<class T>
using MyType = AnotherType<T, MyAllocatorType>;
```

`typedef` 和 `using` 有什么区别？

## 回答

除了 `using` 还可以在模板中使用，其它的都是等同的。
