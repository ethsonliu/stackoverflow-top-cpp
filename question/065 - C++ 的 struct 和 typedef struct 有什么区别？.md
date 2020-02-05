<https://stackoverflow.com/questions/612328/difference-between-struct-and-typedef-struct-in-c>

## 问题

C++ 中下面的两条语句有什么区别么？

```c++
struct Foo { ... };

typedef struct { ... } Foo;
```

## 回答

在 C++ 中没什么区别，主要是 C 语言中有区别。

C 标准（[C89 §3.1.2.3](http://port70.net/~nsz/c/c89/c89-draft.txt), [C99 §6.2.3](http://port70.net/~nsz/c/c99/n1256.html#6.2.3), [C11 §6.2.3](http://port70.net/~nsz/c/c11/n1570.html#6.2.3)）把不同类型的标识符（identifier）分到不同的命名空间（namespace）。

例如标签标识符（tag identifiers）struct/union/enum，普通标识符（ordinary identifiers）以 typedef 修饰和其它常见的。

因此 C 语言中，

```c
struct Foo { ... };
Foo x;
```

这样的用法会报错，因为 Foo 定义在标签命名空间，必须得显示表明 `struct`，即 `struct Foo x`。但每次都加 `struct Foo` 太繁琐了，所以可以加个 `typedef` 来声明别名，这个别名就是普通标识符（C++ 中，所有的 struct/union/enum/class 都是普通标识符）。

```c
struct Foo { ... };
typedef struct Foo Foo;

Foo x;
```

`typedef` 的别名不能在另一个文件前置声明来使用，只能通过 `#include`。

另外，在 C 语言中下面两种定义有一个注意点，

```c
typedef struct Foo { ... } Foo; // 1
typedef struct { ... } Foo;     // 2
```

第 1 个，定义一个名称是 Foo 的结构体，并别名 Foo；第 2 个，定义一个匿名的结构体，并别名 Foo。

两者的区别就是第二个无法被前置声明，因为 Foo 只在 typedef 命名空间。
