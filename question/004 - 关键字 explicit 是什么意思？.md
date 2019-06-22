<https://stackoverflow.com/questions/121162/what-does-the-explicit-keyword-mean>


## 问题

C++ 中的关键字`explicit`是什么意思？

## 最佳回答

我们知道编译器是允许进行隐式转换（implicit conversion）的，就是说如果构造函数只有一个形参，那么是允许从这个参数对象隐式转换为另一个对象类型的，直接看个例子就明白了，

```c++
class Foo
{
public:
  // single parameter constructor, can be used as an implicit conversion
  Foo (int foo) : m_foo (foo) 
  {
  }

  int GetFoo () { return m_foo; }

private:
  int m_foo;
};
```

下面是一个以`Foo`类型为参数的函数，

```c++
void DoBar (Foo foo)
{
  int i = foo.GetFoo ();
}
```

下面是调用构造函数，进行隐式转换的例子，

```c++
int main ()
{
  DoBar (42);
}
```

实参`42`是一个整型，不是`Foo`类型的，但是它可以正常调用，这就是因为隐式转换。

因为存在`Foo (int foo)`这个构造函数，所以可以从`int`隐式转换为`Foo`。同样的，如果你定义了这样的构造函数`Foo (double foo)`，也是允许从`double`隐式转化为`Foo`的。

但是如果你现在在构造函数的前面加个关键字`explicit`，它的意思就是要告诉编译器，这个隐式转换不会再被允许了，当编译到`DoBar (42)`的时候就会报错，除非你显示调用，像这样`DoBar (Foo(42))`。

写代码时还是建议能用`explicit`就用，别省，它会帮助我们减少一些 bug。比如下面这个作死的例子，

 - 一个类构造函数`MyString(int size)`，它可以创建一个指定长度`size`的字符串，而你现在有一个函数`print(const MyString&)`，当调用`print(3)`的时候（其实你是想调用`printf("3")`，因为粗心少敲了双引号），按道理你期望得到的值是`3`，但是实际上得到的只是一个长度为 3 的字符串而已。
 
>译注：能用`explicit`的地方我都会用上，别问我为什么，男人的第六感！
