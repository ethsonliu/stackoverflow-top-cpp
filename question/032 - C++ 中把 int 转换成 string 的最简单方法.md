<https://stackoverflow.com/questions/5590381/easiest-way-to-convert-int-to-string-in-c>

## 问题

有什么好办法可以把一个 `int` 转换成它的 `string` 类型，下面是我所知道的两种方法，还有更好的么？

```c++
int a = 10;
char *intStr = itoa(a);
string str = string(intStr);
```

```c++
int a = 10;
stringstream ss;
ss << a;
string str = ss.str();
```

## 回答

C++ 11 提供了 `std::to_string` 可以快速地转换。

```c++
#include <string> 

std::string s = std::to_string(42);
```
