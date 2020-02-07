<https://stackoverflow.com/questions/204476/what-should-main-return-in-c-and-c>

## 问题

到底是用 `void main()` 还是 `int main()`？`main()` 的返回值是 0 还是 1 有什么区别？

## 回答

在 C 语言中，`void main()` 和 `int main()` 都可以。但在 C++ 中，`void main()` 已被禁止，只能使用 `int main()`。

对于返回值，返回 0 意味着你的程序是正常退出，非 0 是异常退出。
