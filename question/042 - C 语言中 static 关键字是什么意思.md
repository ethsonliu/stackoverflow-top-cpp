<https://stackoverflow.com/questions/572547/what-does-static-mean-in-c>

## 问题

我在不少 C 语言程序中见到 `static` 这个关键词，这和 C# 中的 `static` 含义一样么？谁可以解释一下？

## 回答

**函数内静态变量在多次调用时都可以保留其值。**

```c
#include <stdio.h>

void foo()
{
    int a = 10;
    static int sa = 10;

    a += 5;
    sa += 5;

    printf("a = %d, sa = %d\n", a, sa);
}


int main()
{
    int i;

    for (i = 0; i < 10; ++i)
        foo();
}
```

输出如下，

```
a = 15, sa = 15
a = 15, sa = 20
a = 15, sa = 25
a = 15, sa = 30
a = 15, sa = 35
a = 15, sa = 40
a = 15, sa = 45
a = 15, sa = 50
a = 15, sa = 55
a = 15, sa = 60
```

**一个静态全局变量只对本编译单元可见。**

假如有两个 c 文件，

```c
#include "a.h"

int a;

void func_a()
{
}
```
```c
#include "b.h"

int a;

void func_b()
{
}
```

直接这样编译，那么两个编译单元在链接的时候就会变量 a 重定义报错。

而如果在其中一个 c 文件里将变量 a 设置为 static 类型，就会避免这个报错，因为 static 全局变量是内部链接属性。

