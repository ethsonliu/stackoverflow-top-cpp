<https://stackoverflow.com/questions/3279543/what-is-the-copy-and-swap-idiom>

## 问题

我发现 copy-and-swap 这个名词在很多地方都出现，

- [What are your favorite C++ Coding Style idioms: Copy-swap](https://stackoverflow.com/questions/276173/what-are-your-favorite-c-coding-style-idioms/2034447#2034447)
- [Copy constructor and = operator overload in C++: is a common function possible?](https://stackoverflow.com/questions/1734628/copy-constructor-and-operator-overload-in-c-is-a-common-function-possible/1734640#1734640)
- [What is copy elision and how it optimizes copy-and-swap idiom](https://stackoverflow.com/questions/2143787/what-is-copy-elision-and-how-it-optimizes-copy-and-swap-idiom)
- [C++: dynamically allocating an array of objects?](https://stackoverflow.com/questions/255612/c-dynamically-allocating-an-array-of-objects/255744#255744)

它到底是什么意思？怎么用？在 C++ 11 中它又有什么变化？

## 回答

为什么需要 copy-and-swap 呢？ 任何资源管理类（比如智能指针）都需要遵循一个规则：[三法则](https://github.com/Hapoa/stackoverflow-top-cpp/blob/master/question/014%20-%20%E4%BB%80%E4%B9%88%E6%98%AF%E2%80%9C%E4%B8%89%E6%B3%95%E5%88%99%E2%80%9D%EF%BC%9F.md)。其中复制构造函数和析构函数实现起来比较容易，但是赋值运算符（=）要复杂许多，而 copy-and-swap 就是实现赋值运算符（=）的完美解决方案。它既能避免代码冗余，还可以提供 [强异常安全保证](https://en.wikipedia.org/wiki/Exception_safety)。

那 copy-and-swap 是怎么实现的呢？大致思路是：先用复制构造函数创建一个副本，然后利用函数`swap`交换其成员数据，当作用域退出，副本的析构函数会自动调用。这里有三个注意点：一，复制构造函数应该是可用的；二，这里的`swap`并非指`std::swap`，而是需要我们自己写的，而且需要保证`swap`不会抛出异常；三：析构函数也应该是可用的。

我们以一个例子来更深入地理解。

我们先定义一个类，管理一个动态数组，并实现它的复制构造函数和析构函数，

```c++
#include <algorithm> // std::copy
#include <cstddef> // std::size_t

class dumb_array
{
public:
    // (default) constructor
    dumb_array(std::size_t size = 0)
        : mSize(size),
          mArray(mSize ? new int[mSize]() : nullptr)
    {
    }

    // copy-constructor
    dumb_array(const dumb_array& other)
        : mSize(other.mSize),
          mArray(mSize ? new int[mSize] : nullptr),
    {
        // note that this is non-throwing, because of the data
        // types being used; more attention to detail with regards
        // to exceptions must be given in a more general case, however
        std::copy(other.mArray, other.mArray + mSize, mArray);
    }

    // destructor
    ~dumb_array()
    {
        delete [] mArray;
    }

private:
    std::size_t mSize;
    int* mArray;
};
```

但想让上面的类做的更好，还需要一个赋值运算符（=），如下，

```c++
// the hard part
dumb_array& operator=(const dumb_array& other)
{
    if (this != &other) // (1)
    {
        // get rid of the old data...
        delete [] mArray; // (2)
        mArray = nullptr; // (2) *(see footnote for rationale)

        // ...and put in the new
        mSize = other.mSize; // (3)
        mArray = mSize ? new int[mSize] : nullptr; // (3)
        std::copy(other.mArray, other.mArray + mSize, mArray); // (3)
    }

    return *this;
}
```

相信大多数人都是这样写的，但这种写法存在一些问题。

序号（1）处：判断是否等于自身，这种检查有两个目的。一，防止做无用功；二，防止自赋值时出现问题（看上面的代码就知道了）。但是这种检查没什么意义，因为很少出现，加上它反而徒增消耗。（译注：我随后查看了 boost、folly 和 MSVC 的实现，它们都加上了自判断检查。）


