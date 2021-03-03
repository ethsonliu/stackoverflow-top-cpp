<https://stackoverflow.com/questions/347949/how-to-convert-a-stdstring-to-const-char-or-char>

## 问题

std::string 如何转化成 const char * 或者 char * 类型？

## 回答

`string::c_str()` 的返回类型就是 `const char *`，末尾带结束符 `\0`

```c++
std::string str;
const char * c = str.c_str();
```
