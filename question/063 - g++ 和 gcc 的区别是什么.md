<https://stackoverflow.com/questions/172587/what-is-the-difference-between-g-and-gcc>

## 问题

g++ 和 gcc 的区别是什么？一般 c++ 开发应该用哪一个？

## 回答

`gcc` 和 `g++` 都是 GNU 编译器*套件*（以前只是 GNU C *编译器*）的编译驱动程序。

即使它们可以自行根据文件类型决定使用哪种后端（`cc1`、`cc1plus` ……）除非指定了 `-x language`，但它们还是有一些区别。

最重要的区别可能就是默认上它们会自动链接到不同的库。

根据 GCC 的在线文档 [3.15 Options for Linking](https://gcc.gnu.org/onlinedocs/gcc/Link-Options.html) 以及 [g++ 是如何被调用的](https://gcc.gnu.org/onlinedocs/gcc/Invoking-G_002b_002b.html)，`g++` 与 `gcc -xc++ -lstdc++ -shared-libgcc` （第一个是编译器选项，后面两个是链接器选项）是等价的。通过带 `-v` 参数（显示将要运行的后端工具链指令）运行这两个指令，可以证明它们确实等价。

## 回答

GCC：GNU 编译器套件

它包含所有支持不同语言的 GNU 编译器。

`gcc`：GNU C 编译器
`g++`：GNU C++ 编译器

主要区别：

1.  `gcc` 会编译：`*.c`、`*.cpp`文件，分别作为 C 语言和 C++ 语言。
2.  `g++` 会编译：`*.c`、`*.cpp` 文件，但都作为 C++ 语言文件对待。
3.  若用 `g++` 当作链接器，它会自动链接标准 C++ 库（`gcc` 不这么做，可以添加 `--lstdc++`）。
4.  `gcc` 编译 C 文件时会有更少的预定义宏。
5.  `gcc` 编译 `*.cpp` 文件以及`g++` 编译 `*.c`、`*.cpp` 时会有额外的宏。

当编译 `*.cpp` 文件时的额外宏：

```c
#define __GXX_WEAK__ 1
#define __cplusplus 1
#define __DEPRECATED 1
#define __GNUG__ 4
#define __EXCEPTIONS 1
#define __private_extern__ extern
```
