<https://stackoverflow.com/questions/571394/how-to-find-out-if-an-item-is-present-in-a-stdvector>

## 问题

如何检测一个元素是否在 std::vector 中？

## 回答

可以使用头文件 `<algorithm>` 里的方法 `std::find`，

```c++
#include <algorithm>
#include <vector>
vector<int> vec; 

if (std::find(vec.begin(), vec.end(), item) != vec.end())
   do_this();
else
   do_that();
```
