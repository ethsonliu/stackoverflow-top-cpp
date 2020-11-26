<https://stackoverflow.com/questions/14116003/difference-between-constexpr-and-const>

## 问题

constexpr 和 const 之间有什么区别？

## 回答

**对变量来说，**

const 表示的只是这个变量不可修改，但并未限定这个变量是编译期常量还是运行期常量；而 constexpr 之能是编译期常量。

```c++
const int kSize = 1; // 编译期常量

void func()
{
    const int kRandomNumber = get_a_random_number(); // 运行期常量
    ...
    ...
}
```

对于 kSize，你既可以用 const 修饰，也可以用 constexpr。但对于 kRandomNumber，你只能用 const。

**对函数来说，**

const 修饰的函数一般都是成员函数，用来表示这个函数不会对成员变量产生写操作，这点很好理解。

我们重点来看 constexpr。

