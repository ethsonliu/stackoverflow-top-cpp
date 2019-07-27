<https://stackoverflow.com/questions/21593/what-is-the-difference-between-include-filename-and-include-filename>

## 问题

如题所问，在 C/C++ 中，`#include <filename>`和`#include "filename"`两种写法有什么区别？

## 回答

一般会去三个地方找：

1. 标准系统库路径。比如`stdio.h`，`iostream`所在的路径就是标准系统库路径。
2. 当前文件的所在路径。比如你当前的文件所在路径是`~/project/myproject/src/`，那么`~/project/myproject/src/`就是这个文件的所在路径。
3. 自定义添加的额外路径。比如`-I`指令指定的路径。

`#include <filename>`，预处理器的查找顺序是：1->3->2。

`#include "filename"`，预处理器的查找顺序是：2->3->1。

>译注：对于上面的顺序，对于 1 和 2 是可以保证的，但 3 只是个人猜测，若有不对，请指正。

换句话说，对于标准库文件，用`#include <filename>`，其它的都用`#include "filename"`。
