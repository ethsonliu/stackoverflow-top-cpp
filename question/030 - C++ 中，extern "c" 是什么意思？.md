<https://stackoverflow.com/questions/1041866/what-is-the-effect-of-extern-c-in-c>

## 问题

C++ 中在代码中的 `extern "C"` 是什么意思？

比如，

```c++
extern "C" {
   void foo();
}
```

## 回答

C++ 支持函数的重载，重载这个特性给我们带来了很大的便利。为了支持函数重载的这个特性，C++ 编译器实际上将下面这些重载函数

```c++
void print(int i);
void print(char c);
void print(float f);
void print(char* s);
```

编译为

```
_print_int
_print_char
_print_float
_pirnt_string
```

这样的函数名，来唯一标识每个函数。（不同的编译器实现可能不一样，但都是利用这种机制。）所以当链接是调用 `print(3)` 时，它会去查找 `_print_int(3)` 这样的函数。

C++ 中的变量，编译也类似，如全局变量可能编译 g_xx，类变量编译为 c_xx 等，链接也是按照这种机制去查找相应的变量。

而 C 语言中并没有重载和类这些特性，故并不像 C++ 那样 `print(int i)`，会被编译为 `_print_int`，而是直接编译为 `_print` 等。因此如果直接在 C++ 中调用 C 的函数会失败，因为链接时调用 C 中的 `print(3)` 时，它会去找 `_print_int(3)`。

因此 `extern "C"` 的作用就体现出来了。它用来告诉 **C++ 编译器**，这部分代码要按照 C 语言的方式去编译和链接。
