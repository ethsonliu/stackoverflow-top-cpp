<https://stackoverflow.com/questions/4303513/push-back-vs-emplace-back>

## 问题

```c++
void push_back(const T& value);
void push_back(T&& value);

template< class... Args >
void emplace_back(Args&&... args);
```

`push_back` 和 `emplace_back` 的区别在哪里？

## 回答

`emplace_back` 能就地通过参数构造对象，不需要拷贝或者移动内存，相比 `push_back` 能更好地避免内存的拷贝与移动，使容器插入元素的性能得到进一步提升。在大多数情况下应该优先使用 `emplace_back` 来代替 `push_back`。

下面的代码节选自 <https://en.cppreference.com/w/cpp/container/vector/emplace_back>，可以很好的解释它们的区别，

```c++
#include <vector>
#include <string>
#include <iostream>
 
struct President
{
    std::string name;
    std::string country;
    int year;
 
    President(std::string p_name, std::string p_country, int p_year)
        : name(std::move(p_name)), country(std::move(p_country)), year(p_year)
    {
        std::cout << "I am being constructed.\n";
    }
    President(President&& other)
        : name(std::move(other.name)), country(std::move(other.country)), year(other.year)
    {
        std::cout << "I am being moved.\n";
    }
    President& operator=(const President& other) = default;
};
 
int main()
{
    std::vector<President> elections;
    std::cout << "emplace_back:\n";
    elections.emplace_back("Nelson Mandela", "South Africa", 1994);
 
    std::vector<President> reElections;
    std::cout << "\npush_back:\n";
    reElections.push_back(President("Franklin Delano Roosevelt", "the USA", 1936));
 
    std::cout << "\nContents:\n";
    for (President const& president: elections) {
        std::cout << president.name << " was elected president of "
                  << president.country << " in " << president.year << ".\n";
    }
    for (President const& president: reElections) {
        std::cout << president.name << " was re-elected president of "
                  << president.country << " in " << president.year << ".\n";
    }
}
```

输出：

```
emplace_back:
I am being constructed.
 
push_back:
I am being constructed.
I am being moved.
 
Contents:
Nelson Mandela was elected president of South Africa in 1994.
Franklin Delano Roosevelt was re-elected president of the USA in 1936.
```
