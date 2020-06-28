<https://stackoverflow.com/questions/213121/use-class-or-typename-for-template-parameters>

## 问题

当定义一个函数模板或者一个模板类的时候，下面的两种写法都是可以的，

```c++
template <class T> ...
template <typename T> ...
```

那两者有什么区别呢？

## 回答

在一些简单使用上两者是可以相互替换的，也就是没区别，比如上面你给出的例子。但在有一些场景下是有区别不可替换的，比如，

**情况一**

C++ 允许在类内定义类型别名，

```c++
template<typename param_t>
class Foo
{
    typedef typename param_t::baz sub_t;
};
```

加这个 typename 是为了告诉编译器 param_t::baz 是一个类型而不是类内成员。

**情况二**

当定义模板的模板时，也必须用 class，例如，

```c++
template < template < typename, typename > class Container, typename Type >
```

但在 C++ 17 中，typename 也被允许使用在模板的模板中了。

**情况三**

当显式实例化模板的时候，必须用 class，

```c++
template class Foo<int>;
```
