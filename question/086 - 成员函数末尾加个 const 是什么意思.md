<https://stackoverflow.com/questions/751681/meaning-of-const-last-in-a-function-declaration-of-a-class>

## 问题

比如下面的代码，

```c++
class foobar
{
  public:
     operator int () const;
     const char* foo() const;
};
```

成员函数末尾加个 const 是什么意思？

## 回答

成员函数末尾加个 `const` 表示该函数不允许修改成员变量（除 `mutable` 修饰的变量），且也只能调用 `const` 成员函数。

```c++
class MyClass
{
private:

    mutable int x;
    int y;

public:

    MyClass() { x = y = 0; }

    void Foo()
    {
        x++; // ok
        y++; // ok
        Boo(); // ok
        Boo_c(); // ok
        std::cout << "Foo" << std::endl;
    }

    void Foo() const // 可以重载
    {
        x++; // ok, it's mutable
        // y++; // error
        // Boo(); // error
        Boo_c(); // ok
        std::cout << "Foo const" << std::endl;
    }
    
    void Boo()
    {
    }
    
    void Boo_c() const
    {
    }
};

int main()
{
    MyClass a;
    const MyClass b;
    a.Foo();
    b.Foo();
}
```

输出：

```
Foo
Foo const
```

