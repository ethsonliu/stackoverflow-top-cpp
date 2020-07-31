<https://stackoverflow.com/questions/18335861/why-is-enum-class-preferred-over-plain-enum>

## 问题

我听到一些人建议使用 `enum class`，因为它是类型安全（type safety）的。这到底是什么意思？

## 回答

C++ 有两种枚举（enum），

1. enum class
2. enum

它们的使用也很简单，例如，

```c++
enum class Color { red, green, blue }; // enum class
enum Animal { dog, cat, bird, human }; // enum 
```

两者的区别如下，

**一：作用域不同**

enum 中的 { } 大括号并没有将枚举成员的可见域限制在大括号内，导致 enum 成员曝露到了上一级作用域中。

```c++
#include <iostream>

enum color {red, blue};

int main()
{
    std::cout << blue << std::endl; // 可以直接访问
    std::cin.get();
    return 0;
}
```

就如上面的代码，color 的成员被泄露到了该文件的全局作用域中（虽然它尚不具备外部链接性），可以直接访问，而不需要域运算符的帮助。这样带来的问题就是无法定义同名的枚举成员，例如，

```c++
enum Color1 { red, green, blue }; // 编译报错，重定义
enum Color2 { red, green, blue };
```

再来看看 enum class，

```c++
enum class Color1 { red, green, blue }; // 没问题，可以编译使用
enum class Color2 { red, green, blue };

std::cout << blue << std::endl; // 报错，不可以直接访问
std::cout << Color1::blue << std::endl; // 通过，使用域运算符来访问
```

**二：隐式转换**

```c++
enum Color {red, blue};
enum class Animal {dog, cat};

int main()
{
    Color myColor = red;
    Animal myAnimal = Animal::dog;
    
    int number = myColor; // 可以隐式转换
    number = myAnimal; // 错误，不允许隐式转换
    number = static_cast<int>(myAnimal); // 正确，显示转换
}
```

**三：指定底层所使用的数据类型**

enum 无法指定数据类型，导致我们无法明确枚举类型所占的内存大小。这种麻烦在结构体当中尤为突出，特别是当我们需要内存对齐和填充处理的时候。

其次，当我们使用 enum 时，我们无法决定编译器底层是如何对待 enum 的（比如：signed 和 unsigned）。

而标准规定，enum class 默认的底层数据类型是 int，也可以自己手动指定数据类型，语法如下，

```c++
enum class color:unsigned char {red, blue};
enum calss colorb:long long {yellow, black};
```

## 参考

- [C++11 的 enum class & enum struct & enum](https://blog.csdn.net/sanoseiichirou/article/details/50180533)
