<https://stackoverflow.com/questions/332030/when-should-static-cast-dynamic-cast-const-cast-and-reinterpret-cast-be-used>

## 问题

下面这些类型转换的正确用法和应用场景是什么？

- `static_cast`
- `dynamic_cast`
- `const_cast`
- `reinterpret_cast`
- C 语言风格类型转化`(type)value`
- 函数式风格类型转换`type(value)`

## 最佳回答

**`static_cast`** 是静态转换的意思，也就是在编译期间转换，转换失败的话会抛出一个编译错误。主要用于，

1. 基本数据类型之间的转换。如把 int 转换成 char，把 int 转换成 enum。这种转换的安全性需要开发人员来保证。
2. void 指针转换成目标类型的指针。这种转换的安全性需要开发人员来保证。
3. 任何类型的表达式转换成 void 类型。
4. 有转换构造函数或类型转换函数的类与其它类型之间的转换。例如 double 转 Complex（调用转换构造函数）、Complex 转 double（调用类型转换函数）。
5. 类层次结构中基类和子类之间指针或引用的转换。进行上行转换（即子类的指针或引用转换成基类表示）是安全的，不过一般在进行这样的转化时会省略 static_cast；进行下行转换（即基类指针或引用转换成子类表示）时，由于没有动态类型检查，所以是不安全的，一般用 dynamic_cast 来替代。

**dynamic_cast** 是动态转换，会在运行期借助 RTTI 进行类型转换（这就要求基类必须包含虚函数），主要用于类层次间的下行转换（即基类指针或引用转换成子类表示）。对于指针，如果转换失败将返回 NULL；对于引用，如果转换失败将抛出 std::bad_cast 异常。

```c++
class Base { };
class Derived : public Base { };
 
Base a, *ptr_a;
Derived b, *ptr_b;
 
ptr_a = dynamic_cast<Base *>(&b);  // 成功
ptr_b = dynamic_cast<Derived *>(&a);  // 失败，因为基类无虚函数
```

```c++
class Base { virtual void dummy() {} }; // polymorphic class
class Derived : public Base { int a; }; // so is this
 
Base *ptr_a = new Derived{};
Base *ptr_b = new Base{};
 
Derived *ptr_c = nullptr;
Derived *ptr_d = nullptr;
 
ptr_c = dynamic_cast<Derived *>(ptr_a);  // Fine
ptr_d = dynamic_cast<Derived *>(ptr_b);  // ptr_d will be NULL
 
// Check if downcasting succeeded
if (ptr_c != nullptr) {
	// ptr_a actually points to a Derived object 
}
 
if (ptr_d != nullptr) {
    // ptr_b actually points to a Derived object 
}
```

**const_cast** 主要用来修改类型的 const 或 volatile 属性。
