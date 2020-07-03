<https://stackoverflow.com/questions/2236197/what-is-the-easiest-way-to-initialize-a-stdvector-with-hardcoded-elements>

## 问题

如何优雅的初始化 std:vector？我知道的是下面的写法，

```c++
std::vector<int> ints;

ints.push_back(10);
ints.push_back(20);
ints.push_back(30);
```

还有更好的么？

## 回答

```c++
static const int arr[] = {16,2,77,29};
vector<int> vec (arr, arr + sizeof(arr) / sizeof(arr[0]));
```

如果你的编译器支持 C++ 11 的话，可以直接这样，

```c++
std::vector<int> v = {1, 2, 3, 4};
```
