<https://stackoverflow.com/questions/154136/why-use-apparently-meaningless-do-while-and-if-else-statements-in-macros>

## 问题

比如下面的语句，

```c++
#define FOO(X) do { f(X); g(X); } while (0)
#define FOO(X) if (1) { f(X); g(X); } else
```

为什么不直接写成这样，

```c++
#define FOO(X) f(X); g(X)
```

## 回答

其实是为了把这个宏模拟成一条真实的语句。

```c++
bool x;
...
if (x)
    FOO(1);
```

宏被替换就会变成这样，

```c++
if (x)
    f(1); g(1);
```

很明显，这是不符合预期的，你不得不这么做，

```c++
if (x)
    { FOO(1); }
```

但总有人会忘记加上 `{}`。所以加上 do-while 或者 if-else 就可以解决这个问题。

```c++
if (x)
    do { f(1); g(1); } while (0);
```
