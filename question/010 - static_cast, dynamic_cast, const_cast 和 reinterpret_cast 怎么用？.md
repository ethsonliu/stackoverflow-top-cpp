<https://stackoverflow.com/questions/332030/when-should-static-cast-dynamic-cast-const-cast-and-reinterpret-cast-be-used>

## 问题

下面这些类型转换的正确用法和应用场景是什么？

- `static_cast`
- `dynamic_cast`
- `const_cast`
- `reinterpret_cast`
- C 语言风格类型转化`(type)value`
- 函数式风格类型转换`type(value)`

## 回答

**`static_cast`** 是静态转换的意思，也就是在编译期间转换，转换失败的话会抛出一个编译错误。主要用于，

1. 基本数据类型之间的转换。如把 int 转换成 char，把 int 转换成 enum。这种转换的安全性需要开发人员来保证。
2. void 指针转换成目标类型的指针。这种转换的安全性需要开发人员来保证。
3. 任何类型的表达式转换成 void 类型。
4. 有转换构造函数或类型转换函数的类与其它类型之间的转换。例如 double 转 Complex（调用转换构造函数）、Complex 转 double（调用类型转换函数）。
5. 类层次结构中基类和子类之间指针或引用的转换。进行上行转换（即子类的指针或引用转换成基类表示）是安全的，不过一般在进行这样的转化时会省略 static_cast；进行下行转换（即基类指针或引用转换成子类表示）时，由于没有动态类型检查，所以是不安全的，一般用 dynamic_cast 来替代。

```c++
class Complex{
public:
    Complex(double real = 0.0, double imag = 0.0): m_real(real), m_imag(imag){ }
public:
    operator double() const { return m_real; }  // 类型转换函数
private:
    double m_real;
    double m_imag;
};

int m = 100;
Complex c(12.5, 23.8);
long n = static_cast<long>(m);  // 宽转换，没有信息丢失
char ch = static_cast<char>(m);  // 窄转换，可能会丢失信息
int *p1 = static_cast<int*>(malloc(10 * sizeof(int)));  // 将 void 指针转换为具体类型指针
void *p2 = static_cast<void*>(p1);  // 将具体类型指针，转换为 void 指针
double real= static_cast<double>(c);  // 调用类型转换函数
```

**`dynamic_cast`** 是动态转换，会在运行期借助 RTTI 进行类型转换（这就要求基类必须包含虚函数），主要用于类层次间的下行转换（即基类指针或引用转换成子类表示）。对于指针，如果转换失败将返回 NULL；对于引用，如果转换失败将抛出 std::bad_cast 异常。

```c++
class Base { };
class Derived : public Base { };
 
Base a, *ptr_a;
Derived b, *ptr_b;
 
ptr_a = dynamic_cast<Base *>(&b);  // 成功
ptr_b = dynamic_cast<Derived *>(&a);  // 失败，因为基类无虚函数
```

```c++
class Base { virtual void dummy() {} };
class Derived : public Base { int a; };
 
Base *ptr_a = new Derived{};
Base *ptr_b = new Base{};
 
Derived *ptr_c = nullptr;
Derived *ptr_d = nullptr;
 
ptr_c = dynamic_cast<Derived *>(ptr_a);  // 成功
ptr_d = dynamic_cast<Derived *>(ptr_b);  // 失败，返回 NULL
 
// 检查下行转换是否成功
if (ptr_c != nullptr) {
	// ptr_a actually points to a Derived object 
}
 
if (ptr_d != nullptr) {
    // ptr_b actually points to a Derived object 
}
```

**`const_cast`** 主要用来修改类型的 const 或 volatile 属性。

```c++
int a = 5;
const int* pA = &a;
*pA = 10; // 编译错误，不允许修改 pA 指向的对象
int* pX = const_cast<int*>(pA); // 去掉 const 属性
*pX = 10 // 成功赋值
```

注意，如果你要修改的对象实际上是一个常量，这个转换就可能不会生效。

```c++
const int a = 5; // 常量
const int* pA = &a;
*pA = 10; // 编译错误，不允许修改 pA 指向的对象
int* pX = const_cast<int*>(pA); // 去掉 const 属性
*pX = 10 // 是否会真正地修改结果未知，因为对于 a 来说，编译器一般在编译的时候会把它放进常量表中
```

**`reinterpret_cast`** 是重新解释的意思，顾名思义，reinterpret_cast 这种转换仅仅是对二进制位的重新解释，不会借助已有的转换规则对数据进行调整，非常简单粗暴，所以风险很高。

reinterpret_cast 可以认为是 static_cast 的一种补充，一些 static_cast 不能完成的转换，就可以用 reinterpret_cast 来完成。例如两个具体类型指针之间的转换、int 和指针之间的转换（有些编译器只允许 int 转指针，不允许反过来）。

```c++
class A{
public:
    A(int a = 0, int b = 0): m_a(a), m_b(b){}
private:
    int m_a;
    int m_b;
};

// 将 char* 转换为 float*
char str[] = "reinterpret_cast";
float *p1 = reinterpret_cast<float*>(str);

// 将 int 转换为 int*
int *p = reinterpret_cast<int*>(100);

// 将 A* 转换为 int*
p = reinterpret_cast<int*>(new A(25, 96));
```

**`(type)value`和`type(value)`** 其实是一个意思，只是写法风格的差异而已。它涵盖了上面四种`*_cast`的所有功能，同时它的使用需要完全由程序员自己把控。

## 参考

- <https://www.quora.com/How-do-you-explain-the-differences-among-static_cast-reinterpret_cast-const_cast-and-dynamic_cast-to-a-new-C++-programmer>
- <https://www.cnblogs.com/chio/archive/2007/07/18/822389.html>
- <http://c.biancheng.net/cpp/biancheng/view/3297.html>
- <https://blog.csdn.net/bboyfeiyu/article/details/9057447>

