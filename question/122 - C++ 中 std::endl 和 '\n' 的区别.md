<https://stackoverflow.com/questions/213907/c-stdendl-vs-n>

## 问题

C++ 中 `std::endl` 和 `\n` 有什么区别？

## 回答

除了都是输出一个换行，两者唯一的区别是，`std::endl` 可以刷新输出缓冲区，而 `\n` 不会。说白了就是下面的代码，

```c++
std::cout << std::endl;
```
相当于

```c++
std::cout << '\n' << std::flush;
```
