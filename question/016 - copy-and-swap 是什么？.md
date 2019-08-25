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

相信大多数人都是这样写的，但这种写法会存在三个问题。

序号（1）处：判断是否等于自身，这种检查有两个目的。一，防止做无用功；二，防止自赋值时出现问题（看上面的代码就知道了）。但是这种检查没什么意义，因为很少出现，加上它反而徒增消耗。（译注：我随后查看了 boost、folly 和 MSVC 的实现，它们都加上了自判断检查。）

序号（2）处：仅提供了基本异常安全保证。如果在`new`的时候抛出异常，此时`*this`的内容已被修改（早已被`delete`），无法还原至开始状态。如果想要强异常安全保证，可以这样写，

```c++
dumb_array& operator=(const dumb_array& other)
{
    if (this != &other) // (1)
    {
        // get the new data ready before we replace the old
        std::size_t newSize = other.mSize;
        int* newArray = newSize ? new int[newSize]() : nullptr; // (3)
        std::copy(other.mArray, other.mArray + newSize, newArray); // (3)

        // replace the old data (all are non-throwing)
        delete [] mArray;
        mSize = newSize;
        mArray = newArray;
    }

    return *this;
}
```

序号（3）处：代码冗余，主要是内存申请（new）和复制（copy）部分。如果管理多个资源，那么这里的代码就会变得膨胀。（译注：这里的冗余应该是指与复制构造函数的代码实现有重复。）

>译注：评论区有人指出“一个类管理多个资源”这种做法是不提倡的，作者也表示同意，上面那句话之所以那么说，我觉得更多是突出“冗余膨胀”四字，读者可以不必在此处过多纠结。至于为何这种做法是不提倡的，作者也给出了回答：[单一功能原则](https://zh.wikipedia.org/wiki/%E5%8D%95%E4%B8%80%E5%8A%9F%E8%83%BD%E5%8E%9F%E5%88%99)。

copy-and-swap 就可以同时解决上面的三个问题，做法是这样的，

```c++
class dumb_array
{
public:
    // ...

    friend void swap(dumb_array& first, dumb_array& second) // nothrow
    {
        // enable ADL (not necessary in our case, but good practice)
        using std::swap;

        // by swapping the members of two objects,
        // the two objects are effectively swapped
        swap(first.mSize, second.mSize);
        swap(first.mArray, second.mArray);
    }

    dumb_array& operator=(dumb_array other) // (1)
    {
        swap(*this, other); // (2)

        return *this;
    }
};
```

（其中，swap 被定义为 public friend，理由可参见 [https://stackoverflow.com/questions/5695548/public-friend-swap-member-function](https://stackoverflow.com/questions/5695548/public-friend-swap-member-function) 和 Effective C++ 条款 25。）

注意到 `dumb_array& operator=(dumb_array other)` 的参数是值传递，不应该是引用传递么？就像下面这样，

```c++
dumb_array& operator=(const dumb_array& other)
{
    dumb_array temp(other);
    swap(*this, temp);

    return *this;
}
```

因为无法让编译器充分发挥它优化的优势，可以参考，

- [引用传递的弊端](https://stackoverflow.com/questions/261567/function-parameters-copy-or-pointer/261598#261598)
- [aliasing 的解释](https://zh.wikipedia.org/wiki/%E5%88%AB%E5%90%8D_(%E8%AE%A1%E7%AE%97))
- [aliasing 的弊端](https://stackoverflow.com/questions/9709261/what-is-aliasing-and-how-does-it-affect-performance)

现在来看看它是怎么解决上面那三个问题的。

值传递可以在进入函数体内部的时候就已经实现对象的复制，内存的申请，避免了代码冗余，而无异常的 swap 可以提供强异常安全保证，至于自赋值，这里就更不存在了，因为函数体内部的对象完全是一个新对象。





