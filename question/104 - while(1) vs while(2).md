<https://stackoverflow.com/questions/24848359/which-is-faster-while1-or-while2>

## 问题

下面的代码哪个更快？

```c
while(1) {
    // Some code
}
```

```c
while(2) {
    //Some code
}
```

这是我的一个面试官提出的，我给出的答案是：一样快！

但面试官说， `while(1)` 的更快！

真的是这样么？

## 回答

对于现代编译器来说，肯定是一样快的！

但对于程序员来说可能不是。从写法上来看，`while(1)` 更符合大众的理解，大家都知道你想表达式无限循环，但如果突然出现个 `while(2)`，稍微细心的人都会稍作停留思考作者为什么这么写。
