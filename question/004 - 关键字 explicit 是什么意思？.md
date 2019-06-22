<https://stackoverflow.com/questions/121162/what-does-the-explicit-keyword-mean>


## 问题

C++ 中的关键字`explicit`是什么意思？

## 最佳回答

我们知道编译器是允许进行隐式转换（implicit conversion）的，就是说如果构造函数只有一个形参，那么是允许从这个参数对象隐式转换为另一个对象类型的，直接看
个例子就明白了，

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
