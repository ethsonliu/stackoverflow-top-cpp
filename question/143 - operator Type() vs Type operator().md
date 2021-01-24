<https://stackoverflow.com/questions/44247797/c-operator-parenthesis-operator-type-vs-type-operator>

## 问题

比如 `int operator()` vs `operator int()`，这两者有什么区别？

## 回答

`int operator()` 是函数调用运算符（Function Call Operator），比如，

```c++
struct Foo
{
    int operator()(int a, int b)
    {
        return a + b;
    }
};

...

Foo foo;
int i = foo(1, 2);  // Call the object as a function, and it returns 3 (1+2)
```

`operator int()` 是类型转换运算符（Type Conversion Operator），比如，

```c++
struct Bar
{
    operator int()
    {
        return 123;
    }
};

...

Bar bar;
int i = bar;  // Calls the conversion operator, which returns 123
```
