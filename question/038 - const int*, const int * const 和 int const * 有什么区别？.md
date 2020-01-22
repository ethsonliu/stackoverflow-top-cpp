<https://stackoverflow.com/questions/1143262/what-is-the-difference-between-const-int-const-int-const-and-int-const>

## 问题

我经常搞混 `const int *`, `const int * const` 和 `int const *` 的区别，怎么区分它们呢？ 

## 回答

先参考整理的右左法则：https://github.com/EthsonLiu/personal-notes/blob/master/cpp/%E6%8C%87%E9%92%88%E5%8F%B3%E5%B7%A6%E6%B3%95%E5%88%99.md

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
- `int ** const p` - p is const pointer to a pointer to an int
- `int * const * p` - p is pointer to a const pointer to an int
- `int const ** p` - p is pointer to a pointer to a const int
- `int * const * const p` - p is const pointer to a const pointer to an int
