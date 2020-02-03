<https://stackoverflow.com/questions/54585/when-should-you-use-a-class-vs-a-struct-in-c>

## 问题

C++ 中什么时候用 class 什么时候用 struct？

## 回答

C++ 中的 `class` 和 `struct` 的区别如下：

1. `class` 内成员默认是 private，`struct` 是 public
2. 继承子类时，`class` 默认时 private 继承，`struct` 默认是 public 继承

建议，仅当成员都是 POD 类型且都是 public 的时候用 `struct`。
