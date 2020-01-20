<https://stackoverflow.com/questions/7627098/what-is-a-lambda-expression-in-c11>

## 问题

C++ 11 的 lambda 表达式是什么？什么时候去用它？主要用它解决什么问题呢？

## 回答

### 起因

C++ 03 时代，头文件 `<algorithm>` 有很多方便使用的泛型函数，例如 `std::for_each` 和 `std::transform`。但有的时候这些函数用起来又很麻烦，尤其是存在 [functor](https://stackoverflow.com/questions/356950/what-are-c-functors-and-their-uses) 的情况下。

```c++
#include <algorithm>
#include <vector>

namespace {
  struct f {
    void operator()(int) {
      // do something
    }
  };
}

void func(std::vector<int>& v) {
  f f;
  std::for_each(v.begin(), v.end(), f);
}
```

事实上你只调用了 f 一次，但是你还是需要像上面那样定义一个 strcut，如果这种类似的情况比较多，那么代码看起来就显得很乱。

你可能会想到 functor 本地化的办法来解决这个问题，就像下面这样，

```c++
void func2(std::vector<int>& v) {
  struct {
    void operator()(int) {
       // do something
    }
  } f;
  std::for_each(v.begin(), v.end(), f);
}
```

但是 C++ 03 （C++ 11 已经支持）是不支持这种用法的，因为 f 不能应用于 [模板函数](https://en.cppreference.com/w/cpp/language/function_template)。

### C++ 11 新的解决方案

C++ 11 的 lambda 提供了一种匿名函数，可以完美解决这个问题。

```c++
void func3(std::vector<int>& v) {
  std::for_each(v.begin(), v.end(), [](int) { /* do something here*/ });
}
```

### lambda 的返回类型

如果你的 lambda 很简单，那么编译器可以自动推断出你的返回类型，

```c++
void func4(std::vector<double>& v) {
  std::transform(v.begin(), v.end(), v.begin(),
                 [](double d) { return d < 0.00001 ? 0 : d; }
                 );
}
```

但如果你的 lambda 比较复杂，比如下面这样，返回类型既可能是整型也可能是浮点型，

```c++
void func4(std::vector<double>& v) {
    std::transform(v.begin(), v.end(), v.begin(),
        [](double d) {
            if (d < 0.0001) {
                return 0;
            } else {
                return d;
            }
        });
}
```

那么你可以手动给它执行你要返回的类型，你需要这么做，

```c++
void func4(std::vector<double>& v) {
    std::transform(v.begin(), v.end(), v.begin(),
        [](double d) -> double {
            if (d < 0.0001) {
                return 0;
            } else {
                return d;
            }
        });
}
```

### 捕获变量（Capturing variables）

那 lambda 如何在函数体内使用外部的变量呢？其实很简单。

```c++
void func5(std::vector<double>& v, const double& epsilon) {
    std::transform(v.begin(), v.end(), v.begin(),
        [epsilon](double d) -> double {
            if (d < epsilon) {
                return 0;
            } else {
                return d;
            }
        });
}
```





