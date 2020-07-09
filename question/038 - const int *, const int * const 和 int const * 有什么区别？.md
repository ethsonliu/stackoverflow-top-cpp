<https://stackoverflow.com/questions/1143262/what-is-the-difference-between-const-int-const-int-const-and-int-const>

## 问题

我经常搞混 `const int *`, `const int * const` 和 `int const *` 的区别，怎么区分它们呢？ 

## 回答

先参考翻译的文章：[读懂 C 的类型声明（译）](https://ethsonliu.com/2020/04/reading-c-type-declarations.html)

对于 `const` 关键词，直接填入即可，例如，

- `int * p` - p is pointer to int
- `int const * p` - p is pointer to const int
- `int * const p` - p is const pointer to int
- `int const * const p` - p is const pointer to const int

其中，下面两个是等同的，只是顺序的不同而已，

- const int * == int const *
- const int * const == int const * const

当然还有更复杂的，

- `int ** p` - p is pointer to pointer to int
- `int ** const p` - p is const pointer to pointer to int
- `int * const * p` - p is pointer to const pointer to int
- `int const ** p` - p is pointer to pointer to const int
- `int * const * const p` - p is const pointer to const pointer to int
