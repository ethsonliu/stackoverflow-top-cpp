<https://stackoverflow.com/questions/332030/when-should-static-cast-dynamic-cast-const-cast-and-reinterpret-cast-be-used>

## 问题

下面这些类型转换的正确用法和应用场景是什么？

- `static_cast`
- `dynamic_cast`
- `const_cast`
- `reinterpret_cast`
- C 语言风格类型转化`(type)value`
- 函数式风格类型转换`type(value)`

## 最佳回答

**`static_cast`** 是一个很有用的转换，建议能用就用。它可以用于那些通常的隐式转换（比如`int`转`float`，指针转`void*`）。

在转换类对象为另一个类型的时候，`static_cast`会自动调用它的显示/隐式转换函数（比如`class Base{...} base; int i = static_cast<int>(base);`）
