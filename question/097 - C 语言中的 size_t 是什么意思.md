<https://stackoverflow.com/questions/2550774/what-is-size-t-in-c>

## 问题

我知道 `size_t` 是作为 `sizeof` 的返回类型，但这个类型到底是什么？干什么用的？

比如下面的 for 循环，我是用 `int` 还是 `size_t`？

```c++
for (i = 0; i < some_size; i++)
```

## 回答

`size_t` 定义在头文件 `stddef.h` 中，标准规定它是一个至少 16 位的无符号整型。在我的机器上它是这样的，

```c
typedef unsigned long size_t;
```
