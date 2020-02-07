<https://stackoverflow.com/questions/313970/how-to-convert-stdstring-to-lower-case>

## 问题

我想将一个 `std::string` 全部转为小写字母，有什么好办法么？

## 回答

```c++
#include <algorithm>
#include <cctype>
#include <string>

std::string data = "Abc";
std::transform(data.begin(), data.end(), data.begin(),
    [](unsigned char c){ return std::tolower(c); });
```
