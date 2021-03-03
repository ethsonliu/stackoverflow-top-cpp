<https://stackoverflow.com/questions/1433204/how-do-i-use-extern-to-share-variables-between-source-files>

## 问题

如何使用 extern 关键字在不同的源文件间共享变量？

## 回答

首先需要知道 **声明** 和 **定义** 的区别。声明并不分配内存，定义才会。

```c++
extern int a; // 声明，a 的定义可能在其它的文件
int b; // 定义，b 占有实际的内存
```

下面是一个用法，

```c++
// file1.cpp
#include <iostream>

extern int a; // a 的定义在另一个文件

void func()
{
    a++;
}
```

```c++
// file2.cpp
#include <iostream>

int a = 1;

void func()
{
    a++;
}
```
