<https://stackoverflow.com/questions/154469/unnamed-anonymous-namespaces-vs-static-functions>

## 问题

匿名命名空间和关键词 static 都可以让其声明的变量或函数变为内部链接属性，那么它们之间的区别是什么？

## 回答

首先，匿名命名空间比关键词 static 的功能更丰富，因此我们更推荐使用前者。比如，static 无法修饰类（class/struct），

```c++
// 非法代码
static class sample_class { /* class body */ };
static struct sample_struct { /* struct body */ };


// 合法代码
namespace 
{  
    class sample_class { /* class body */ };
    struct sample_struct { /* struct body */ };
}
```

参考：

- [Superiority of unnamed namespace over static?](https://stackoverflow.com/questions/4422507/superiority-of-unnamed-namespace-over-static)
- [Why an unnamed namespace is a “superior” alternative to static?](https://stackoverflow.com/questions/4977252/why-an-unnamed-namespace-is-a-superior-alternative-to-static)
