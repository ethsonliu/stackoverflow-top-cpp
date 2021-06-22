<https://stackoverflow.com/questions/98650/what-is-the-strict-aliasing-rule>

## 问题

在这个问题 [common undefined behavior in C](https://stackoverflow.com/questions/98340/what-are-the-common-undefinedunspecified-behavior-for-c-that-you-run-into) 中提到的 strict aliasing 是什么意思？

## 回答

> 译注：以下转自参考一，并做稍微修改。

### 什么是 Aliasing？

[Understanding Strict Aliasing](https://cellperformance.beyond3d.com/articles/2006/06/understanding-strict-aliasing.html) 一文中这样描述：

当两个指针指向同一块区域或对象时，我们称一个指针 alias 另一个指针。

[strict aliasing](http://dbp-consulting.com/tutorials/StrictAliasing.html) 一文中这样描述：

Aliasing 是指多于一个的左值指向同一块区域。

比如：

```c++
// 例一
int i;
int *pi = &i; // pi alias i

// 例二
int i;
void foo(int &i1, int &i2){}
foo(i,i); // i1 alias i2

// 例三
int i;
float *pf = (float*)&i;
*pf = 0; // pf alias i 违反了 strict alias 么
```


### 什么是 Strict Aliasing？

按照 [Understanding Strict Aliasing](https://cellperformance.beyond3d.com/articles/2006/06/understanding-strict-aliasing.html) 一文描述：

Strict aliasing 是 C 或 C++ 编译器的一种假设：不同类型的指针绝对不会指向同一块内存区域。

暂且不管这句话，我们看个例子：

```c
#include <stdio.h>

int a;
int f(float *b)
{
    a = 1;
    *b = 0;
    return a;
}

int main()
{
    printf("%d\n", f((float*)&a));
    return 0;
}
```

用 GCC4.4 加 -O3 编译：

```
debao@ubuntu:~/ttt$ gcc-4.4 -Wall -O3 hello.c
hello.c: In function ‘main’:
hello.c:7: warning: dereferencing pointer ‘a.16’ does break strict-aliasing rules
hello.c:13: note: initialized from here
```

运行程序，结果为 1。

```
debao@ubuntu:~/ttt$ ./a.out
1
```

而如果不加 -O3，程序结果为 0。

```
debao@ubuntu:~/ttt$ gcc-4.4 -Wall  hello.c
debao@ubuntu:~/ttt$ ./a.out
0
```

原因是警告信息说：对指针 a.16 的解引用打破了 strict-aliasing 规则。

```c
int a;
int f(float *b)
{
    a = 1;
    *b = 0;
    return a;
}
```

按照 strict aliasing 规则，编译器认为 `a` 和 `*b` 绝不会指向同一块存储区域，故而优化后返回 a 时直接返回了 1。

那么哪些 alias 是不会破坏规则的呢？或者说规则在哪儿呢？

strict aliasing 一文中将这些条文可以总结如下：

1. 兼容类型或差别仅在于 signed、unsigned、const、volatile 的类型（比如 `const unsigned long *` 和 `long*`）
2. 聚合类型(struct 或 class)或联合类型( unio n)可以 alias 它们所包含的类型（比如 int 和 包含有 int 的结构体(包括间接包含)）
3. 字符类型(`char *`、`signed char*`、`unsinged char*`)可以 alias 任何类型的指针
4. C++ 基类的类型(可能带有 const、volatile 等 cv 修饰)可以 alias 派生类的类型

为什么要 strict alias？主要目的应该就是为了使编译器能生成更高效的代码。考虑下面的代码：

```c
int a;
void f( double * b )
{
    a = 1;
    *b = 2.0;
    g(a);
}
```

如果没有 strict alias 假定，编译器必须始终假设 b 可能会指向 a 所在的地址，从而不能将 g(a) 调用优化成 g(1)。

strict alias 规则中，为什么允许 char 类型可以 alias 任何对象呢？

Character pointer types are often used in the bytewise manipulation of objects;a byte stored through such a character pointer may well end up in an object of any type.

为什么 class、struct、union 可以 alias 它们包含的对象类型呢？

Structure and union types also have problematic aliasing properties:

```c
struct fi{ float f; int i;};
void f( struct fi * fip, int * ip )
{
```

参考：

- [C/C++ Strict Alias 小记](https://blog.csdn.net/dbzhang800/article/details/6720141)
- [What is the Strict Aliasing Rule and Why do we care?](https://gist.github.com/shafik/848ae25ee209f698763cffee272a58f8)
