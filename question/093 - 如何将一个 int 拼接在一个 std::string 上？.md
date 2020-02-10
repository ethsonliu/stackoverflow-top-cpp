<https://stackoverflow.com/questions/191757/how-to-concatenate-a-stdstring-and-an-int>

## 问题

例如下面的语句，

```c++
std::string name = "John";
int age = 21;
```

如何把它们连接起来变成 `John21`？

## 回答

```c++
std::string name = "John";
int age = 21;
std::string result;

// 1. with C++11
result = name + std::to_string(age);

// 2. with IOStreams
std::stringstream sstm;
sstm << name << age;
result = sstm.str();

// 3. with itoa
char numstr[21]; // enough to hold all numbers up to 64-bits
result = name + itoa(age, numstr, 10);

// 4. with sprintf
char numstr[21]; // enough to hold all numbers up to 64-bits
sprintf(numstr, "%d", age);
result = name + numstr;
```

1. 安全。需要 C++11 的支持，需 `#include <string>`，标准所支持，跨平台。
2. 安全、低效、代码啰嗦。需 `#include <sstream>`，标准所支持，跨平台。
3. 容易出错（你需要分配足够的内存）、快速、代码啰嗦。`itoa` 不是一个标准方法，不确定可以在所有平台适用。
4. 容易出错（你需要分配足够的内存）、快速、代码啰嗦。标准所支持，跨平台。
