<https://stackoverflow.com/questions/1282295/what-exactly-is-nullptr>

## 问题

C++11 中的 nullptr 到底是什么？.

## 回答

在 C 语言中，我们使用 `NULL` 表示空指针，它实际上是一个宏，具体被定义为，

```c
/* C 语言程序 */
#define NULL ((void*)0)
#define NULL 0 /* 两者皆可 */
```

注：因为在C 语言中，是允许 `void` 指针隐式转换为其它类型指针的，所以 `#define NULL ((void*)0)` 这样的定义不会有问题。

C++ 语言出现后，为了保持对 C 语言的兼容，保留了 `NULL`，但对 `NULL` 的定义变得更为严格，

```c++
/* C++ 语言程序 */
#ifdef __cplusplus
    #define NULL 0
#else
    #define NULL ((void*)0)
#endif
```

`NULL` 被定义为 0，而不是 `((void*)0)`，因为在 C++ 语言中，`void` 指针是不可以隐式转换为其它类型指针的，必须显示转换，

```c++
/* C++ 语言程序 */
#define NULL ((void*)0) /* 如果在 C++ 语言中这么定义的话 */

int* a = NULL; /* 隐式转换，错误 */
int* a = (int*)NULL; /* 显示转换，正确，但很麻烦，所以 NULL 都会被定义为 0 */
```

在 C++98 之前（包括 C++98），在对 `NULL` 的使用上，都一直存在一个问题，假设有下面的代码，

```c++
/* C++ 语言程序 */
void func(int i);
void func(char* p);

func(NULL); /* 该调用哪个？*/
```

`NULL` 其实就是等于 0，对于上面的两个函数，它都是符合的，如此，就会出现语义二义性的错误。

为了解决上述重载函数所带来的问题，C++11 的 `nullptr` 应运而生。`nullptr` 实质上是一个常量，实现代码大致如下，

```c++
/* C++ 语言程序 */
const /* 常量 */
class
{
public:
	template<class T>
	operator T*() const /* 向任意类型的非成员指针转换 */
	{
		return 0;
	}

	template<class C, class T>
	operator T C::*() const /* 向任意类型的成员指针转换 */
	{
		return 0;
	}

private:
	void operator&() const /* 不可取地址 */
	{
	}
} nullptr = {};
```

`nullptr` 只是一个常量，这就意味着我们可以在程序中随意定义一个与其名称相同的标识符，但因为 `nullptr` 在实际编程中的应用实在太广泛，因此 C++ 编译器一般都会把 `nullptr` 定为关键字，避免程序员的滥用。

当然，C++11 发布后，并没有因为 `nullptr` 的出现，而摒弃 `NULL`，主要是为了兼容旧版程序。

最后，总结一下，

1. **在 C 语言编程中，请使用 `NULL`。** 此时的 `NULL`，要么是 `((void*)0)`，要么是 0，对于 C 语言而言，都无所谓。

2. **在 C++ 语言编程中，请使用 `nullptr`。** 既为了避免以后出现 bug，也为了养成一个良好的编程习惯。

## 参考文献：

- [http://www.stroustrup.com/N1488-nullptr.pdf](http://www.stroustrup.com/N1488-nullptr.pdf)
