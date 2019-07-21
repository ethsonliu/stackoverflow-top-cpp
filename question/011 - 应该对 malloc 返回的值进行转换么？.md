<https://stackoverflow.com/questions/605845/do-i-cast-the-result-of-malloc>

## 问题

在这个 [问题](http://stackoverflow.com/questions/571945/getting-a-stack-overflow-exception-when-declaring-a-large-array) 里，有人在
[评论](http://stackoverflow.com/questions/571945/getting-a-stack-overflow-exception-when-declaring-a-large-array#comment388297_571961) 里建议不要对`malloc`返回的值进行转换。举个例子，

应该这样，

```c++
int *sieve = malloc(sizeof(int) * length);
```

而不是，

```c++
int *sieve = (int *) malloc(sizeof(int) * length);
```

谁能说下为什么？

## 回答

1. C 中，从 void* 到其它类型的指针是自动转换的，所以无需手动加上类型转换。

2. 在旧式的 C 编译器里，如果一个函数没有原型声明，那么编译器会认为这个函数返回 int。那么，如果碰巧代码里忘记包含头文件 <stdlib.h>，那么编译器看到malloc 调用时，会认为它返回一个 int。
  
   在实际运行时，malloc 的返回值（一个v oid* 指针），会被直接解释成一个 int。如果这时强制转换这个值，实际就是将 int 直接转换为 void* 。这里就有 2 个问题：一，void* 和 int 可能不能无损地相互转换，例如它们长度不同，或者编译器不支持这种转换。二， 即使可以相互转换，它们的表示也可能不同，即需要显示转换。而直接将 void* 当成 int 来用，然后再把这个 int 转换回 void* ，在这种情况下肯定是会有问题的。
  
   如果这时没有强转 malloc 的返回值，编译器看到要把 int 转换为 int* ，就会发出一条警告。而如果强转了 malloc 的返回值，编译器就不会做警告了，在运行时就可能出问题。
  
3. 强制转换 malloc 的返回值并没有错，但画蛇添足！例如，日后你有可能把 double* 改成 int* ，这时，你就要把所有相关的`(double*)malloc(sizeof(double))`改成`(int*)malloc(sizeof(int))`，如果改漏了，那么你的程序就存在 bug。

注意，以上都是以 C 语言为基础上成立的，在 C++ 中则是不一样，C++ 是不允许 void* 隐式转换为其它类型的，所以需要显示转换，一般用 static_cast。

## 参考

- <https://blog.csdn.net/bestone0213/article/details/40829203>
