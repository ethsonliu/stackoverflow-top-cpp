<https://stackoverflow.com/questions/7868936/read-file-line-by-line-using-ifstream-in-c>

## 问题

下面的文本文件，

```
5 3
6 4
7 1
10 5
11 6
12 3
12 4
```

其中每行的数字，比如 `5 3` 是一对坐标，如何使用 C++ 按行读取获取这些坐标？

## 回答

首先，定义一个 `ifstream` 对象，

```c++
#include <fstream>
std::ifstream infile("thefile.txt");
```

接着有两种方法可以实现，

1. 按空格和换行符进行分割

```c++
int a, b;
while (infile >> a >> b)
{
    // process pair (a,b)
}
```

2. 读取每行，然后按空格分割

```c++
#include <sstream>
#include <string>

std::string line;
while (std::getline(infile, line))
{
    std::istringstream iss(line);
    int a, b;
    if (!(iss >> a >> b)) { break; }

    // process pair (a,b)
}
```
