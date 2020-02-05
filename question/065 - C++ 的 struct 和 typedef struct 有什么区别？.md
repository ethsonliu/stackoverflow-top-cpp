<https://stackoverflow.com/questions/612328/difference-between-struct-and-typedef-struct-in-c>

## 问题

C++ 中下面的两条语句有什么区别么？

```c++
struct Foo { ... };

typedef struct { ... } Foo;
```

## 回答

在 C++ 中没什么区别，主要是 C 语言中有区别。

C 语言中，

```c
struct Foo { ... };
Foo x;
```

这样的用法会报错，必须得显示表明 `struct`，即 `struct Foo x`。但每次都加 `struct Foo` 太繁琐了，所以可以加个 `typedef` 来避免。

```c
struct Foo { ... };
typedef struct Foo Foo;

Foo x;
```
