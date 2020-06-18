<https://stackoverflow.com/questions/1674032/static-const-vs-define-vs-enum>

## 问题

在 C 语言中，下面的用法哪个最好？

```c
static const int var = 5;
```

```c
#define var 5
```

```c
enum { var = 5 };
```

## 回答

取决于你用来干什么。

1. `static const int var = 5`
2. `#define var 5`
3. `enum { var = 5 }`

- 如果需要传指针，那只能用 (1)
- (1) 不能作为全局作用域下数组的维数定义，而 (2)(3) 可以
- (1) 不能作为函数作用域下静态数组的维数定义，而 (2)(3) 可以
- (1) 不能在 switch 语句下使用，而 (2)(3) 可以
- (1) 不能用来初始化另一个静态常量，而 (2)(3) 可以
- (2) 可以用预处理器判断是否已存在，而 (1)(3) 不可以

大多场景下，enum 是最佳选择。

**如果是 C++ 语言，那么自始至终都应该使用 (1)。**
