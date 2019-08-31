<https://stackoverflow.com/questions/21593/what-is-the-difference-between-include-filename-and-include-filename>

## 问题

如题所问，在 C/C++ 中，`#include <filename>`和`#include "filename"`两种写法有什么区别？

## 回答

`<filename>`一般会去系统路径和编译器预指定的路径找。比如 Windows 系统库的`#include <Windows.h>`，Linux 系统库的`#include <sys/socket.h>`，C/C++ 编译器已预指定的的标准库`#include <stdio.h>`。GCC 命令中`-I`会给编译器另自指定一条搜寻路径，该路径下的文件，也会用`<>`包含。

`"filename"`一般会去工程目录下找，如果你的工程下有一个文件`~/MyProject/src/widget.h`里包含了`#include "simple_dialog.h"`，那么它会去`~/MyProject/src/`下去找，找不到再依照`<>`查找的路径去找。

总的来说，

- 系统库、标准库、编译器指定的路径（比如 GCC 的`-I`命令），都以`#include <>`来包含文件。
- 程序员自己创建的工程文件，都以`#include ""`来包含。

